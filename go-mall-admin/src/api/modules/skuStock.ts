import { ResPage, SkuStock } from "@/api/interface/index";
import http from "@/api";

// 批量更新sku的库存
export const batchUpdateSkuStockApi = (params: SkuStock.BatchUpdateSkuStockParams) => {
  return http.put(`/skuStocks/update/batchUpdateByProductId`, params);
};

// 根据商品id分页查询sku的库存
export const getSkuStocksByProductIdApi = (params: SkuStock.GetSkuStocksByProductIdParams) => {
  return http.get<ResPage<SkuStock.SkuStockModel>>(`/skuStocks/search/searchByProductId`, params);
};
