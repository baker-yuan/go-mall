package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/baker-yuan/go-mall/application/admin/internal/usecase/assembler"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductUseCase 商品管理Service实现类
type ProductUseCase struct {
	productRepo                      IProductRepo                      // 操作商品
	brandRepo                        IBrandRepo                        // 操作商品品牌
	productCategoryRepo              IProductCategoryRepo              // 操作商品分类
	memberPriceRepo                  IMemberPriceRepo                  // 商品会员价格
	productLadderRepo                IProductLadderRepo                // 商品阶梯价格
	productFullReductionRepo         IProductFullReductionRepo         // 产品满减
	skuStockRepo                     ISkuStockRepo                     // sku库存
	productAttributeValueRepo        IProductAttributeValueRepo        // 产品参数信息
	subjectProductRelationRepo       ISubjectProductRelationRepo       // 专题商品关系
	prefrenceAreaProductRelationRepo IPrefrenceAreaProductRelationRepo // 优选专区和产品关系
}

// NewProduct 创建商品管理Service实现类
func NewProduct(
	productRepo IProductRepo,
	brandRepo IBrandRepo,
	productCategoryRepo IProductCategoryRepo,
	memberPriceRepo IMemberPriceRepo,
	productLadderRepo IProductLadderRepo,
	productFullReductionRepo IProductFullReductionRepo,
	skuStockRepo ISkuStockRepo,
	productAttributeValueRepo IProductAttributeValueRepo,
	subjectProductRelationRepo ISubjectProductRelationRepo,
	prefrenceAreaProductRelationRepo IPrefrenceAreaProductRelationRepo,
) *ProductUseCase {
	return &ProductUseCase{
		productRepo:                      productRepo,
		brandRepo:                        brandRepo,
		productCategoryRepo:              productCategoryRepo,
		memberPriceRepo:                  memberPriceRepo,
		productLadderRepo:                productLadderRepo,
		productFullReductionRepo:         productFullReductionRepo,
		skuStockRepo:                     skuStockRepo,
		productAttributeValueRepo:        productAttributeValueRepo,
		subjectProductRelationRepo:       subjectProductRelationRepo,
		prefrenceAreaProductRelationRepo: prefrenceAreaProductRelationRepo,
	}
}

// CreateProduct 添加商品
func (c ProductUseCase) CreateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error {
	// 数据转换
	product := assembler.AddOrUpdateProductParamToEntity(param)

	// 事务执行
	return db.Transaction(ctx, func(ctx context.Context) error {
		// 创建商品
		if err := c.productRepo.CreateWithTX(ctx, product); err != nil {
			return err
		}
		// 根据促销类型设置价格：会员价格、阶梯价格、满减价格
		productID := product.ID

		// 会员价格
		if len(param.GetMemberPrices()) != 0 {
			if err := c.memberPriceRepo.BatchCreateWithTX(ctx, productID, assembler.MemberPricesToEntity(param.GetMemberPrices())); err != nil {
				return err
			}
		}

		// 阶梯价格
		if len(param.GetProductLadders()) != 0 {
			if err := c.productLadderRepo.BatchCreateWithTX(ctx, productID, assembler.ProductLaddersToEntity(param.GetProductLadders())); err != nil {
				return err
			}
		}

		// 满减价格
		if len(param.GetProductFullReductions()) != 0 {
			if err := c.productFullReductionRepo.BatchCreateWithTX(ctx, productID, assembler.ProductFullReductionsToEntity(param.GetProductFullReductions())); err != nil {
				return err
			}
		}

		// 添加sku库存信息
		if len(param.GetSkuStocks()) != 0 {
			// 处理sku的编码
			c.handleSkuStockCode(param.GetSkuStocks(), productID)
			if err := c.skuStockRepo.BatchCreateWithTX(ctx, productID, assembler.SkuStocksToEntity(param.GetSkuStocks())); err != nil {
				return err
			}
		}

		// 添加商品参数，添加自定义商品规格
		if len(param.GetAttributeValues()) != 0 {
			if err := c.productAttributeValueRepo.BatchCreateWithTX(ctx, productID, assembler.ProductAttributeValuesToEntity(param.GetAttributeValues())); err != nil {
				return err
			}
		}

		// 关联专题
		if len(param.GetSubjectProductRelations()) != 0 {
			if err := c.subjectProductRelationRepo.BatchCreateWithTX(ctx, productID, assembler.SubjectProductRelationsToEntity(param.GetSubjectProductRelations())); err != nil {
				return err
			}
		}

		// 关联优选
		if len(param.GetPrefrenceAreaProductRelations()) != 0 {
			if err := c.prefrenceAreaProductRelationRepo.BatchCreateWithTX(ctx, productID, assembler.PrefrenceAreaProductRelationsToEntity(param.GetPrefrenceAreaProductRelations())); err != nil {
				return err
			}
		}

		return nil
	})
}

func (c ProductUseCase) handleSkuStockCode(skuStockList []*pb.SkuStock, productId uint64) {
	if len(skuStockList) == 0 {
		return
	}
	for i := 0; i < len(skuStockList); i++ {
		skuStock := skuStockList[i]
		if skuStock.SkuCode == "" {
			// 日期
			date := time.Now().Format("20060102")
			// 四位商品id
			productIdStr := fmt.Sprintf("%04d", productId)
			// 三位索引id
			indexIdStr := fmt.Sprintf("%03d", i+1)
			// 设置sku编码
			skuStock.SkuCode = strings.Join([]string{date, productIdStr, indexIdStr}, "")
			skuStockList[i] = skuStock
		}
	}
}

// UpdateProduct 修改商品
func (c ProductUseCase) UpdateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error {
	oldProduct, err := c.productRepo.GetByID(ctx, param.GetId())
	if err != nil {
		return err
	}

	// 数据转换
	product := assembler.AddOrUpdateProductParamToEntity(param)
	product.ID = param.Id
	product.CreatedAt = oldProduct.CreatedAt

	// 更新商品
	return c.productRepo.Update(ctx, product)
}

// GetProducts 分页查询商品
func (c ProductUseCase) GetProducts(ctx context.Context, param *pb.GetProductsParam) ([]*pb.Product, uint32, error) {
	opts := make([]db.DBOption, 0)
	if param.GetId() != nil {
		opts = append(opts, c.productRepo.WithByID(param.GetId().GetValue()))
	}
	if len(param.GetName()) != 0 {
		opts = append(opts, c.productRepo.WithByName(param.GetName()))
	}
	if len(param.GetProductSn()) != 0 {
		opts = append(opts, c.productRepo.WithByProductSN(param.GetProductSn()))
	}
	if param.GetBrandId() != nil {
		opts = append(opts, c.productRepo.WithByBrandID(param.GetBrandId().GetValue()))
	}
	if param.GetProductCategoryId() != nil {
		opts = append(opts, c.productRepo.WithByProductCategoryID(param.GetProductCategoryId().GetValue()))
	}
	if param.GetPublishStatus() != nil {
		opts = append(opts, c.productRepo.WithByPublishStatus(param.GetPublishStatus().GetValue()))
	}
	if param.GetVerifyStatus() != nil {
		opts = append(opts, c.productRepo.WithByVerifyStatus(param.GetVerifyStatus().GetValue()))
	}
	products, pageTotal, err := c.productRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	categories, err := c.productCategoryRepo.GetByIDs(ctx, products.CategoryIDs())
	if err != nil {
		return nil, 0, err
	}

	brands, err := c.brandRepo.GetByIDs(ctx, products.BrandIDs())
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.Product, 0)
	for _, product := range products {
		results = append(results, assembler.ProductEntityToModel(product, categories.NameMap(), brands.NameMap()))

	}
	return results, pageTotal, nil
}

// GetProduct 根据id获取商品
func (c ProductUseCase) GetProduct(ctx context.Context, id uint64) (*pb.Product, error) {
	product, err := c.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	categories, err := c.productCategoryRepo.GetByIDs(ctx, []uint64{product.ProductCategoryID})
	if err != nil {
		return nil, err
	}

	brands, err := c.brandRepo.GetByIDs(ctx, []uint64{product.BrandID})
	if err != nil {
		return nil, err
	}

	return assembler.ProductEntityToModel(product, categories.NameMap(), brands.NameMap()), nil
}

// DeleteProduct 删除商品
func (c ProductUseCase) DeleteProduct(ctx context.Context, id uint64) error {
	return c.productRepo.DeleteByID(ctx, id)
}
