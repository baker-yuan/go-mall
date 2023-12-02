import { ResPage, OrderReturnReason } from "@/api/interface/index";
import http from "@/api";

// 添加退货原因
export const createOrderReturnReasonApi = (params: OrderReturnReason.OrderReturnReasonModel) => {
  return http.post(`/orderReturnReasons`, params);
};

// 修改退货原因
export const updateOrderReturnReasonApi = (orderReturnReason: OrderReturnReason.OrderReturnReasonModel) => {
  return http.put(`/orderReturnReasons/${orderReturnReason.id}`, orderReturnReason);
};

// 分页查询退货原因
export const getOrderReturnReasonsApi = (params: OrderReturnReason.ReqOrderReturnReasonListParams) => {
  return http.get<ResPage<OrderReturnReason.OrderReturnReasonModel>>(`/orderReturnReasons`, params);
};

// 根据id获取退货原因
export const getOrderReturnReasonApi = (id: number) => {
  return http.get<OrderReturnReason.OrderReturnReasonModel>(`/orderReturnReasons/${id}`);
};

// 根据id删除退货原因
export const deleteOrderReturnReasonApi = (id: number) => {
  return http.delete(`/orderReturnReasons/${id}`);
};
