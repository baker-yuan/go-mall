import { ResPage, Subject, TransferValue } from "@/api/interface/index";
import http from "@/api";

// 添加专题
export const createSubjectApi = (params: Subject.SubjectModel) => {
  return http.post(`/subjects`, params);
};

// 修改专题
export const updateSubjectApi = (subject: Subject.SubjectModel) => {
  return http.put(`/brands/${subject.id}`, subject);
};

// 分页查询专题
export const getSubjectsApi = (params: Subject.ReqSubjectListParams) => {
  return http.get<ResPage<Subject.SubjectModel>>(`/subjects`, params);
};

// 根据id获取专题
export const getSubjectApi = (id: number) => {
  return http.get<Subject.SubjectModel>(`/subjects/${id}`);
};

// 根据id删除专题
export const deleteSubjectApi = (id: number) => {
  return http.delete(`/subjects/${id}`);
};

// 查询所有专题
export const getAllSubjectTransferValueSyncApi = async () => {
  let params = {
    pageNum: 1,
    pageSize: 10000
  };
  let data = await getSubjectsApi(params);
  return transformSubjectToTransferValue(data.data.data);
};

// 数据转换函数
function transformSubjectToTransferValue(data: Subject.SubjectModel[]): TransferValue[] {
  if (!data) {
    return [];
  }
  return data.map(item => ({
    key: item.id,
    label: item.title
  }));
}
