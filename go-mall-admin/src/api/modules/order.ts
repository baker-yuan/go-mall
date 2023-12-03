import { ResPage, Order } from "@/api/interface/index";
import http from "@/api";

// 添加订单
export const createOrderApi = (params: Order.OrderModel) => {
  return http.post(`/orders`, params);
};

// 修改订单
export const updateOrderApi = (order: Order.OrderModel) => {
  return http.put(`/orders/${order.id}`, order);
};

// 分页查询订单
export const getOrdersApi = (params: Order.ReqOrderListParams) => {
  return http.get<ResPage<Order.OrderModel>>(`/orders`, params);
};

// 根据id获取订单
export const getOrderApi = (id: number) => {
  return http.get<Order.OrderModel>(`/orders/${id}`);
};

// 根据id删除订单
export const deleteOrderApi = (id: number) => {
  return http.delete(`/orders/${id}`);
};

// 根据id获取订单
export const getOrderSyncApi = async (id: number) => {
  let response = await getOrderApi(id);
  return response.data;
};
