import { ResPage, OrderReturnApply } from "@/api/interface/index";
import http from "@/api";

// 分页查询订单退货申请
export const getOrderReturnAppliesApi = (params: OrderReturnApply.ReqOrderReturnApplyListParams) => {
  return http.get<ResPage<OrderReturnApply.OrderReturnApplyModel>>(`/orderReturnApplies`, params);
};

// 根据id获取订单退货申请
export const getOrderReturnApplyApi = (id: number) => {
  return http.get<OrderReturnApply.OrderReturnApplyModel>(`/orderReturnApplies/${id}`);
};
