// 格式化状态
import { Order } from "@/api/interface";

export const formatStatus = (value: number) => {
  if (value === 1) {
    return "待发货";
  } else if (value === 2) {
    return "已发货";
  } else if (value === 3) {
    return "已完成";
  } else if (value === 4) {
    return "已关闭";
  } else if (value === 5) {
    return "无效订单";
  } else {
    return "待付款";
  }
};

export const formatPayType = (value: number) => {
  if (value === 1) {
    return "支付宝";
  } else if (value === 2) {
    return "微信";
  } else {
    return "未支付";
  }
};

export const formatSourceType = (value: number) => {
  if (value === 1) {
    return "APP订单";
  } else {
    return "PC订单";
  }
};

export const formatOrderType = (value: number) => {
  if (value === 1) {
    return "秒杀订单";
  } else {
    return "正常订单";
  }
};

export const formatPayStatus = (value: number) => {
  if (value === 0) {
    return "未支付";
  } else if (value === 4) {
    return "已退款";
  } else {
    return "已支付";
  }
};
export const formatDeliverStatus = (value: number) => {
  if (value === 0 || value === 1) {
    return "未发货";
  } else {
    return "已发货";
  }
};

export const formatAddress = (order: Order.OrderModel) => {
  let str = order.receiverProvince;
  if (order.receiverCity != null) {
    str += "  " + order.receiverCity;
  }
  str += "  " + order.receiverRegion;
  str += "  " + order.receiverDetailAddress;
  return str;
};

export const formatNull = (value: any) => {
  if (value === undefined || value === null || value === "") {
    return "暂无";
  } else {
    return value;
  }
};

export const formatLongText = (value: string) => {
  if (value === undefined || value === null || value === "") {
    return "暂无";
  } else if (value.length > 8) {
    return value.substr(0, 8) + "...";
  } else {
    return value;
  }
};

export const formatProductAttr = (value: string) => {
  if (value == null) {
    return "";
  } else {
    let attr = JSON.parse(value);
    let result = "";
    for (let i = 0; i < attr.length; i++) {
      result += attr[i].key;
      result += ":";
      result += attr[i].value;
      result += ";";
    }
    return result;
  }
};
