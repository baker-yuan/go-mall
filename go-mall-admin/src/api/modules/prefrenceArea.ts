import { ResPage, PrefrenceArea } from "@/api/interface/index";
import http from "@/api";

// 添加优选专区
export const createPrefrenceAreaApi = (params: PrefrenceArea.PrefrenceAreaModel) => {
  return http.post(`/prefrenceAreas`, params);
};

// 修改优选专区
export const updatePrefrenceAreaApi = (prefrenceArea: PrefrenceArea.PrefrenceAreaModel) => {
  return http.put(`/brands/${prefrenceArea.id}`, prefrenceArea);
};

// 分页查询优选专区
export const getPrefrenceAreasApi = (params: PrefrenceArea.ReqPrefrenceAreaListParams) => {
  return http.get<ResPage<PrefrenceArea.PrefrenceAreaModel[]>>(`/prefrenceAreas`, params);
};

// 根据id获取优选专区
export const getPrefrenceAreaApi = (id: number) => {
  return http.get<PrefrenceArea.PrefrenceAreaModel>(`/prefrenceAreas/${id}`);
};

// 根据id删除优选专区
export const deletePrefrenceAreaApi = (id: number) => {
  return http.delete(`/prefrenceAreas/${id}`);
};
