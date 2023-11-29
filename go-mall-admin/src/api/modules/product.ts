import { ResPage, Product } from "@/api/interface/index";
import http from "@/api";

// 添加商品
export const createProductApi = (params: Product.ProductModel) => {
  return http.post(`/products`, params);
};

// 修改商品
export const updateProductApi = (product: Product.ProductModel) => {
  return http.put(`/products/${product.id}`, product);
};

// 分页查询商品
export const getProductsApi = (params: Product.ReqProductListParams) => {
  return http.get<ResPage<Product.ProductModel[]>>(`/products`, params);
};

// 根据id获取商品
export const getProductApi = (id: number) => {
  return http.get<Product.AddOrUpdateProductModel>(`/products/${id}`);
};

// 根据id删除商品
export const deleteProductApi = (id: number) => {
  return http.delete(`/products/${id}`);
};
