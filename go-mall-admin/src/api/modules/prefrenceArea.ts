import { ResPage, PrefrenceArea, TransferValue } from "@/api/interface/index";
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
  return http.get<ResPage<PrefrenceArea.PrefrenceAreaModel>>(`/prefrenceAreas`, params);
};

// 根据id获取优选专区
export const getPrefrenceAreaApi = (id: number) => {
  return http.get<PrefrenceArea.PrefrenceAreaModel>(`/prefrenceAreas/${id}`);
};

// 根据id删除优选专区
export const deletePrefrenceAreaApi = (id: number) => {
  return http.delete(`/prefrenceAreas/${id}`);
};

// 查询所有优选专区
export const getAllPrefrenceAreaTransferValueSyncApi = async () => {
  let params = {
    pageNum: 1,
    pageSize: 10000
  };
  let data = await getPrefrenceAreasApi(params);
  return transformPrefrenceAreaToTransferValue(data.data.data);
};

// 数据转换函数
function transformPrefrenceAreaToTransferValue(data: PrefrenceArea.PrefrenceAreaModel[]): TransferValue[] {
  if (!data) {
    return [];
  }
  return data.map(item => ({
    key: item.id,
    label: item.name
  }));
}
