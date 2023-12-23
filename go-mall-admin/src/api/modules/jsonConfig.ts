import { ResPage, JsonDynamicConfig } from "@/api/interface/index";
import http from "@/api";

// 添加JSON动态配置
export const createJsonDynamicConfigApi = (params: JsonDynamicConfig.JsonDynamicConfigModel) => {
  return http.post(`/jsonDynamicConfigs`, params);
};

// 修改JSON动态配置
export const updateJsonDynamicConfigApi = (jsonDynamicConfig: JsonDynamicConfig.JsonDynamicConfigModel) => {
  return http.put(`/jsonDynamicConfigs/${jsonDynamicConfig.id}`, jsonDynamicConfig);
};

// 分页查询JSON动态配置
export const getJsonDynamicConfigsApi = (params: JsonDynamicConfig.ReqJsonDynamicConfigListParams) => {
  return http.get<ResPage<JsonDynamicConfig.JsonDynamicConfigModel>>(`/jsonDynamicConfigs`, params);
};

// 根据id获取JSON动态配置
export const getJsonDynamicConfigApi = (id: number) => {
  return http.get<JsonDynamicConfig.JsonDynamicConfigModel>(`/jsonDynamicConfigs/${id}`);
};

// 根据id删除JSON动态配置
export const deleteJsonDynamicConfigApi = (id: number) => {
  return http.delete(`/jsonDynamicConfigs/${id}`);
};
