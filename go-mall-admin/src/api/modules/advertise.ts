import { ResPage, HomeAdvertise } from "@/api/interface/index";
import http from "@/api";

// 添加首页轮播广告表
export const createHomeAdvertiseApi = (params: HomeAdvertise.HomeAdvertiseModel) => {
  return http.post(`/homeAdvertises`, params);
};

// 修改首页轮播广告表
export const updateHomeAdvertiseApi = (homeAdvertise: HomeAdvertise.HomeAdvertiseModel) => {
  return http.put(`/homeAdvertises/${homeAdvertise.id}`, homeAdvertise);
};

// 分页查询首页轮播广告表
export const getHomeAdvertisesApi = (params: HomeAdvertise.ReqHomeAdvertiseListParams) => {
  return http.get<ResPage<HomeAdvertise.HomeAdvertiseModel>>(`/homeAdvertises`, params);
};

// 根据id获取首页轮播广告表
export const getHomeAdvertiseApi = (id: number) => {
  return http.get<HomeAdvertise.HomeAdvertiseModel>(`/homeAdvertises/${id}`);
};

// 根据id删除首页轮播广告表
export const deleteHomeAdvertiseApi = (id: number) => {
  return http.delete(`/homeAdvertises/${id}`);
};
