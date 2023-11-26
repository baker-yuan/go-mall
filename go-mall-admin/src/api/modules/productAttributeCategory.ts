import {ProductAttributeCategory, ReqPage, ResPage} from "@/api/interface/index";
import http from "@/api";
import {getBrands} from "@/api/modules/brand";

// 添加产品属性分类
export const createProductAttributeCategoryApi = (params: ProductAttributeCategory.ProductAttributeCategoryModel) => {
  return http.post<ProductAttributeCategory.ProductAttributeCategoryModel[]>(`/productAttributeCategories`, params);
};

// 修改产品属性分类
export const updateProductAttributeCategoryApi = (brand: ProductAttributeCategory.ProductAttributeCategoryModel) => {
  return http.put<ProductAttributeCategory.ProductAttributeCategoryModel[]>(`/productAttributeCategories/${brand.id}`, brand);
};

// 分页查询产品属性分类
export const getProductAttributeCategoriesApi = (params: ProductAttributeCategory.ReqProductAttributeCategoryListParams) => {
  return http.get<ResPage<ProductAttributeCategory.ProductAttributeCategoryModel>>(`/productAttributeCategories`, params);
};

// 根据id获取产品属性分类
export const getProductAttributeCategoryApi = (id: number) => {
  return http.get<ProductAttributeCategory.ProductAttributeCategoryModel>(`/productAttributeCategories/${id}`);
};

// 根据id删除产品属性分类
export const deleteProductAttributeCategoryApi = (id: number) => {
  return http.delete<ProductAttributeCategory.ProductAttributeCategoryModel>(`/productAttributeCategories/${id}`);
};
