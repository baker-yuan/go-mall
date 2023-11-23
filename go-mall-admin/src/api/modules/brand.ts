import { Brand, ResPage } from "@/api/interface/index";
import http from "@/api";

// 添加商品品牌
export const createBrand = (params: Brand.BrandModel) => {
  return http.post<Brand.BrandModel[]>(`/brands`, params);
};

// 修改商品品牌
export const updateBrand = (brand: Brand.BrandModel) => {
  return http.put<Brand.BrandModel[]>(`/brands/${brand.id}`, brand);
};

// 分页查询商品品牌
export const getBrands = (params: Brand.ReqBrandListParams) => {
  return http.get<ResPage<Brand.BrandModel>>(`/brands`, params);
};

// 根据id获取商品品牌
export const getBrand = (id: number) => {
  return http.get<Brand.BrandModel>(`/brands/${id}`);
};

// 根据id删除商品品牌
export const deleteBrandApi = (id: number) => {
  return http.delete<Brand.BrandModel>(`/brands/${id}`);
};
