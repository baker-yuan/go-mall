// 请求响应参数（不包含data）

export interface Result {
  code: string;
  msg: string;
}

// 请求响应参数（包含data）
export interface ResultData<T = any> extends Result {
  data: T;
}

// 分页响应参数
export interface ResPage<T> {
  data: T[];
  pageNum: number;
  pageSize: number;
  total: number;
}

// 分页请求参数
export interface ReqPage {
  pageNum: number;
  pageSize: number;
}

// 文件上传模块
export namespace Upload {
  export interface ResFileUrl {
    fileUrl: string;
  }
}

// 登录模块
export namespace Login {
  export interface ReqLoginForm {
    username: string;
    password: string;
  }
  export interface ResLogin {
    access_token: string;
  }
  export interface ResAuthButtons {
    [key: string]: string[];
  }
}

// https://element-plus.org/zh-CN/component/cascader.html
export interface CascaderValue {
  value: number;
  label: string;
  children?: CascaderValue[];
}

export interface OptionValue {
  value: number;
  label: string;
}

export interface TransferValue {
  key: number;
  label: string;
}

// 分类管理模块
export namespace Category {
  // 分类模型
  export interface CategoryModel {
    id: number; // 主键
    parentId: number; // 上级分类的编号：0表示一级分类
    name: string; // 名称
    level: number; // 分类级别：0->1级；1->2级
    productCount: number; // 商品数量
    productUnit: number; // 商品单位
    navStatus: number; // 是否显示在导航栏：0->不显示；1->显示
    showStatus: number; // 显示状态：0->不显示；1->显示
    sort: number; // 排序
    icon: string; // 图标
    keywords: string; // 关键字
    description: string; // 描述
  }
  // 分页查询分类
  export interface ReqCategoryListParams extends ReqPage {
    parentId: number | null; // 上级分类的编号：0表示一级分类
  }
  // 分类管理
  export interface CategoryTree {
    category: CategoryModel;
    children: CategoryModel[];
  }
}

// 品牌管理模块
export namespace Brand {
  // 品牌模型
  export interface BrandModel {
    id: number; // 主键
    name: string; // 名称
    firstLetter: string; // 首字母
    sort: number; // 排序
    factoryStatus: number; // 是否为品牌制造商：0->不是；1->是
    showStatus: number; // 是否显示
    productCount: number; // 产品数量
    productCommentCount: number; // 产品评论数量
    logo: string; // 品牌logo
    bigPic: string; // 专区大图
    brandStory: string; // 品牌故事
  }

  // 分页查询商品品牌
  export interface ReqBrandListParams extends ReqPage {
    name?: string; // 名称
  }
}

// 品牌管理模块
export namespace ProductAttributeCategory {
  // 品牌模型
  export interface ProductAttributeCategoryModel {
    id: number; // 主键
    name: string; // 类型名称
    attributeCount: number; // 属性数量
    paramCount: number; // 参数数量
  }

  // 分页查询商品品牌
  export interface ReqProductAttributeCategoryListParams extends ReqPage {
    name?: string; // 名称
  }
}

// 商品属性参数表管理模块
export namespace ProductAttribute {
  // 商品属性参数表模型
  export interface ProductAttributeModel {
    id: number; // 编号
    productAttributeCategoryId: number; // 产品属性分类表ID
    name: string; // 属性名称
    selectType: number; // 属性选择类型：0->唯一；1->单选；2->多选
    inputType: number; // 属性录入方式：0->手工录入；1->从列表中选取
    inputList: string; // 可选值列表，以逗号隔开
    sort: number; // 排序字段
    filterType: number; // 分类筛选样式：1->普通；1->颜色
    searchType: number; // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
    relatedStatus: number; // 相同属性产品是否关联；0->不关联；1->关联
    handAddStatus: number; // 是否支持手动新增；0->不支持；1->支持
    type: number; // 属性的类型；0->规格；1->参数
  }

  // 分页查询商品属性参数表
  export interface ReqProductAttributeListParams extends ReqPage {
    type?: number; // 属性的类型；0->规格；1->参数
    productAttributeCategoryId?: number; // 产品属性分类表ID
  }
}

// 商品管理模块
export namespace Product {
  // 商品模型
  export interface ProductModel {
    // 基本信息
    id: number; // 主键
    productCategoryId: number; // 商品分类id
    name: string; // 商品名称
    subTitle: string; // 副标题
    brandId: number; // 品牌id
    description: string; // 商品描述
    productSn: string; // 货号
    price: number; // 价格
    originalPrice: number; // 市场价
    stock: number; // 库存
    unit: string; // 单位
    weight: number; // 商品重量，默认为克
    sort: number; // 排序

    // 促销信息
    giftPoint: number; // 赠送的积分
    giftGrowth: number; // 赠送的成长值
    usePointLimit: number; // 限制使用的积分数
    previewStatus: number; // 是否为预告商品：0->不是；1->是
    publishStatus: number; // 上架状态：0->下架；1->上架
    newStatus: number; // 新品状态:0->不是新品；1->新品
    recommandStatus: number; // 推荐状态；0->不推荐；1->推荐
    serviceIds: string; // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
    detailTitle: string; // 详情标题
    detailDesc: string; // 详情描述
    keywords: string; // 关键字
    note: string; // 备注
    promotionType: number; // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
    promotionPrice: number; // 促销价格
    promotionStartTime: number; // 促销开始时间
    promotionEndTime: number; // 促销结束时间

    // 属性信息
    productAttributeCategoryId: number; // 品牌属性分类id
    pic: string; // 图片
    albumPics: string; // 画册图片，连产品图片限制为5张，以逗号分割
    detailHtml: string; // 电脑端详情
    detailMobileHtml: string; // 移动端详情

    // 状态
    verifyStatus: number; // 审核状态：0->未审核；1->审核通过
    deleteStatus: number; // 删除状态：0->未删除；1->已删除

    // 其他
    feightTemplateId: number; // 运费模版id
    sale: number; // 销量
    lowStock: number; // 库存预警值
    promotionPerLimit: number; // 活动限购数量

    // 冗余字段
    brandName: string; // 品牌名称
    productCategoryName: string; // 商品分类名称

    // 设置
    productLadders: ProductLadder[]; // 商品阶梯价格设置
    productFullReductions: ProductFullReduction[]; // 商品满减价格设置
    memberPrices: MemberPrice[]; // 商品会员价格设置 {"memberLevelId":0,"memberPrice":0,"memberLevelName":""}
    skuStocks: SkuStock.SkuStockModel[]; // 商品sku库存信息 {"lowStock":0,"pic":"","price":0,"sale":0,"skuCode":"","spData":"","stock":0}
    productAttributeValues: ProductAttributeValue[]; // 商品属性设置 {"productAttributeId":0,"value":""}
    subjectProductRelations: SubjectProductRelation[]; // 商品专题设置 {"subjectId":0}
    prefrenceAreaProductRelations: PrefrenceAreaProductRelation[]; // 商品优选设置 {"prefrenceAreaId":0}
  }

  // 商品阶梯价格设置
  export interface ProductLadder {
    id?: number; // 主键
    productId?: number; // 商品id
    count: number; // 满足的商品数量
    discount: number; // 折扣
    price: number; // 折后价格
  }

  // 商品满减价格设置
  export interface ProductFullReduction {
    id?: number; // 主键
    productId?: number; // 商品id
    fullPrice: number; // 商品满足金额
    reducePrice: number; // 商品减少金额
  }

  // 商品会员价格设置
  export interface MemberPrice {
    id: number; // 主键
    productId: number; // 商品id
    memberLevelId: number; // 会员等级id
    memberPrice: number; // 会员价格
    memberLevelName: string; // 会员等级名称
  }

  // 商品属性设置
  export interface ProductAttributeValue {
    id?: number; // 主键ID
    productId?: number; // 商品ID
    productAttributeId: number; // 商品属性ID
    value: string; // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
  }

  // 商品专题设置
  export interface SubjectProductRelation {
    id?: number; // 主键
    subjectId: number; // 专题ID
    productId?: number; // 产品ID
  }

  // 商品优选设置
  export interface PrefrenceAreaProductRelation {
    id?: number; // 主键
    prefrenceAreaId: number; // 优选专区ID
    productId?: number; // 产品ID
  }

  // 添加修改商品
  export interface AddOrUpdateProductModel extends ProductModel {
    productLadders: ProductLadder[]; // 商品阶梯价格设置
    productFullReductions: ProductFullReduction[]; // 商品满减价格设置
    memberPrices: MemberPrice[]; // 商品会员价格设置 {"memberLevelId":0,"memberPrice":0,"memberLevelName":""}
    skuStocks: SkuStock.SkuStockModel[]; // 商品sku库存信息 {"lowStock":0,"pic":"","price":0,"sale":0,"skuCode":"","spData":"","stock":0}
    productAttributeValues: ProductAttributeValue[]; // 商品属性设置 {"productAttributeId":0,"value":""}
    subjectProductRelations: SubjectProductRelation[]; // 商品专题设置 {"subjectId":0}
    prefrenceAreaProductRelations: PrefrenceAreaProductRelation[]; // 商品优选设置 {"prefrenceAreaId":0}
  }

  // 分页查询商品
  export interface ReqProductListParams extends ReqPage {}
}

// sku的库存管理模块
export namespace SkuStock {
  // sku的库存模型
  export interface SkuStockModel {
    id: number; // 主键
    productId: number; // 产品ID
    skuCode: string; // sku编码
    price: number; // 价格
    stock: number; // 库存
    lowStock: number; // 预警库存
    pic: string; // 展示图片
    sale: number; // 销量
    promotionPrice: number; // 单品促销价格
    lockStock: number; // 锁定库存
    spData: string; // 商品销售属性，json格式
  }

  // 分页查询sku的库存
  export interface ReqSkuStockListParams extends ReqPage {}
  // 根据商品id分页查询sku的库存
  export interface GetSkuStocksByProductIdParams extends ReqPage {
    productId: number; // 产品ID
    skuCode?: string; // sku编码
  }
  // 批量更新sku的库存
  export interface BatchUpdateSkuStockParams {
    productId: number; // 产品ID
    skuStocks: SkuStockModel[]; // 商品SKU
  }
}

// 专题表管理模块
export namespace Subject {
  // 专题表模型
  export interface SubjectModel {
    id: number; // 主键
    categoryId: number; // 分类id
    title: string; // 标题
    pic: string; // 专题主图
    productCount: number; // 关联产品数量
    recommendStatus: number; // 推荐状态
    createTime: number; // 创建时间
    collectCount: number; // 收藏数量
    readCount: number; // 阅读数量
    commentCount: number; // 评论数量
    albumPics: string; // 画册图片，用逗号分割
    description: string; // 描述
    showStatus: number; // 显示状态：0->不显示；1->显示
    content: string; // 内容
    forwardCount: number; // 转发数
    categoryName: string; // 专题分类名称
  }

  // 分页查询专题表
  export interface ReqSubjectListParams extends ReqPage {}
}

// 优选专区管理模块
export namespace PrefrenceArea {
  // 优选专区模型
  export interface PrefrenceAreaModel {
    id: number; // 主键
    name: string; // 名称
    subTitle: string; // 子标题
    pic: string; // 展示图片
    sort: number; // 排序
    showStatus: number; // 显示状态
  }

  // 分页查询优选专区
  export interface ReqPrefrenceAreaListParams extends ReqPage {}
}
