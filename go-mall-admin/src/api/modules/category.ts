import { CascaderValue, Category, ResPage } from "@/api/interface/index";
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

// 商品分类
export const categoryTreeApi = async () => {
  const response = await http.get<Category.CategoryTree[]>(`/categories/search/categoryTree`);
  // 检查响应是否有效
  if (response && response.data) {
    // 转换数据
    let res = transformData(response.data);
    console.log("transformData", JSON.stringify(res));
    return res;
  } else {
    // 如果响应无效，返回一个空数组
    return [];
  }
};

// 数据转换函数
function transformData(data: Category.CategoryTree[]): CascaderValue[] {
  if (!data) {
    return [];
  }
  return data.map(item => ({
    value: item.category.id,
    label: item.category.name,
    children: item.children ? transformChildren(item.children) : undefined
  }));
}

function transformChildren(children: Category.CategoryModel[]): CascaderValue[] {
  if (!children) {
    return [];
  }
  return children.map(child => ({
    value: child.id,
    label: child.name
  }));
}
