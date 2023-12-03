<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      :columns="columns"
      :request-api="getTableList"
      :init-param="initParam"
      :data-callback="dataCallback"
      @darg-sort="sortTable"
      :border="false"
    >
      <!-- 提交时间 -->
      <template #createTime="scope">
        <div>{{ formatTimestamp(scope.row.createTime) }}</div>
      </template>
      <!-- 订单金额 -->
      <template #totalAmount="scope">
        <div>￥{{ scope.row.totalAmount }}</div>
      </template>
      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="Pointer" @click="handleViewOrder('查看', scope.row)">查看订单</el-button>
        <el-button
          type="primary"
          v-show="scope.row.status === 0"
          link
          :icon="Delete"
          @click="handleCloseOrder(scope.$index, scope.row)"
        >
          关闭订单
        </el-button>
        <el-button
          type="primary"
          link
          :icon="Shop"
          @click="handleDeliveryOrder(scope.$index, scope.row)"
          v-show="scope.row.status === 1"
        >
          订单发货
        </el-button>
        <el-button
          v-show="scope.row.status === 2 || scope.row.status === 3"
          type="primary"
          link
          :icon="View"
          @click="handleViewLogistics(scope.$index, scope.row)"
        >
          订单跟踪
        </el-button>
        <el-button type="danger" link :icon="Delete" v-show="scope.row.status === 4" @click="handleDeleteOrder(scope.row)">
          删除订单
        </el-button>
      </template>
    </ProTable>
    <OrderDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="OrderManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { Order } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import OrderDrawer from "@/views/oms/order/detail.vue";
import { Delete, Pointer, View, Shop } from "@element-plus/icons-vue";
import { deleteOrderApi, getOrdersApi } from "@/api/modules/order";
import { formatTimestamp } from "@/utils/time";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<Order.OrderModel>[]>([
  { type: "selection", fixed: "left", width: 70 },
  {
    prop: "id",
    label: "编号",
    search: { el: "input" }
  },
  {
    prop: "orderSn",
    label: "订单编号",
    search: { el: "input" }
  },
  {
    prop: "createTime",
    label: "提交时间"
  },
  {
    prop: "memberUsername",
    label: "用户账号"
  },
  {
    prop: "totalAmount",
    label: "订单金额"
  },
  {
    prop: "payType",
    label: "支付方式",
    enum: [
      {
        value: 0,
        label: "未支付"
      },
      {
        value: 1,
        label: "支付宝"
      },
      {
        value: 2,
        label: "微信"
      }
    ]
  },
  {
    prop: "sourceType",
    label: "订单来源",
    enum: [
      {
        value: 0,
        label: "PC订单"
      },
      {
        value: 1,
        label: "APP订单"
      }
    ]
  },
  {
    prop: "status",
    label: "订单状态",
    enum: [
      {
        value: 0,
        label: "待付款"
      },
      {
        value: 1,
        label: "待发货"
      },
      {
        value: 2,
        label: "已发货"
      },
      {
        value: 3,
        label: "已完成"
      },
      {
        value: 4,
        label: "已关闭"
      },
      {
        value: 5,
        label: "无效订单"
      }
    ]
  },
  { prop: "operation", label: "操作", fixed: "right", width: 250 }
]);

// 如果表格需要初始化请求参数，直接定义传给 ProTable (之后每次请求都会自动带上该参数，此参数更改之后也会一直带上，改变此参数会自动刷新表格数据)
const initParam = reactive({});

// dataCallback 是对于返回的表格数据做处理，如果你后台返回的数据不是 list && total && pageNum && pageSize 这些字段，可以在这里进行处理成这些字段
// 或者直接去 hooks/useTable.ts 文件中把字段改为你后端对应的就行
const dataCallback = (data: any) => {
  console.log("dataCallback", data);
  return {
    list: data.data,
    total: data.pageTotal,
    pageNum: data.pageNum,
    pageSize: data.pageSize
  };
};

// 如果你想在请求之前对当前请求参数做一些操作，可以自定义如下函数：params 为当前所有的请求参数（包括分页），最后返回请求列表接口
// 默认不做操作就直接在 ProTable 组件上绑定	:requestApi="getTableList"
const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  console.log("getTableList", newParams);
  return getOrdersApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除订单
const handleDeleteOrder = async (row: Order.OrderModel) => {
  await useHandleData(deleteOrderApi, row.id, `删除【${row.name}】订单`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof OrderDrawer> | null>(null);
const handleViewOrder = (title: string, row: Partial<Order.OrderModel> = {}) => {
  const params = {
    title,
    row: { ...row }
  };
  drawerRef.value?.acceptParams(params);
};
</script>
