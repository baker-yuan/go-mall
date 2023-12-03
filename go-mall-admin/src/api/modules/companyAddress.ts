import { ResPage, CompanyAddress, OptionValue } from "@/api/interface/index";
import http from "@/api";

// 添加公司收发货地址
export const createCompanyAddressApi = (params: CompanyAddress.CompanyAddressModel) => {
  return http.post(`/companyAddresses`, params);
};

// 修改公司收发货地址
export const updateCompanyAddressApi = (companyAddress: CompanyAddress.CompanyAddressModel) => {
  let id = companyAddress.id;
  return http.put(`/companyAddresses/${id}`, companyAddress);
};

// 分页查询公司收发货地址
export const getCompanyAddressesApi = (params: CompanyAddress.ReqCompanyAddressListParams) => {
  return http.get<ResPage<CompanyAddress.CompanyAddressModel>>(`/companyAddresses`, params);
};

// 根据id获取公司收发货地址
export const getCompanyAddressApi = (id: number) => {
  return http.get<CompanyAddress.CompanyAddressModel>(`/companyAddresses/${id}`);
};

// 根据id删除公司收发货地址
export const deleteCompanyAddressApi = (id: number) => {
  return http.delete(`/companyAddresses/${id}`);
};

// 分页查询公司收发货地址
export const getCompanyAddressOptionsApi = async () => {
  let params = {
    pageNum: 1,
    pageSize: 10000
  };
  const res = await getCompanyAddressesApi(params);
  return transformCompanyAddressesData(res.data.data);
};

// 数据转换函数
function transformCompanyAddressesData(data: CompanyAddress.CompanyAddressModel[]): OptionValue[] {
  if (!data) {
    return [];
  }
  return data.map(item => ({
    value: item.id,
    label: item.addressName
  }));
}
