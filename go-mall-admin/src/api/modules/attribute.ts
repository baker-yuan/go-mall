import { ResPage, ProductAttribute } from "@/api/interface/index";
import http from "@/api";

// 添加商品属性参数表
export const createProductAttributeApi = (params: ProductAttribute.ProductAttributeModel) => {
  return http.post(`/productAttributes`, params);
};

// 修改商品属性参数表
export const updateProductAttributeApi = (productAttribute: ProductAttribute.ProductAttributeModel) => {
  return http.put(`/productAttributes/${productAttribute.id}`, productAttribute);
};

// 分页查询商品属性参数表
export const getProductAttributesApi = (params: ProductAttribute.ReqProductAttributeListParams) => {
  return http.get<ResPage<ProductAttribute.ProductAttributeModel>>(`/productAttributes`, params);
};

// 根据id获取商品属性参数表
export const getProductAttributeApi = (id: number) => {
  return http.get<ProductAttribute.ProductAttributeModel>(`/productAttributes/${id}`);
};

// 根据id删除商品属性参数表
export const deleteProductAttributeApi = (id: number) => {
  return http.delete(`/productAttributes/${id}`);
};

// 分页查询商品属性参数表
export const getProductAttributesSyncApi = async (params: ProductAttribute.ReqProductAttributeListParams) => {
  return getProductAttributesApi(params);
};
