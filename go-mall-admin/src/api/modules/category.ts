import {Category, ResPage} from "@/api/interface/index";
import http from "@/api";

// 添加商品分类
export const createCategory = (params: Category.CategoryModel) => {
  return http.post<Category.CategoryModel[]>(`/categories`, params);
};

// 修改商品分类
export const updateCategory = (category: Category.CategoryModel) => {
  return http.put<Category.CategoryModel[]>(`/categories/${category.id}`, category);
};

// 分页查询商品分类
export const getCategories = (params: Category.ReqCategoryListParams) => {
  return http.get<ResPage<Category.CategoryModel>>(`/categories`, params);
};

// 根据id获取商品分类
export const getCategory = (id: number) => {
  return http.get<Category.CategoryModel>(`/categories/${id}`);
};

// 删除商品分类
export const deleteCategoryApi = (id: number) => {
  return http.delete<Category.CategoryModel>(`/categories/${id}`);
};
