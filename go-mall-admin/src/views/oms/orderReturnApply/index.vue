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
      <!-- 申请时间 -->
      <template #createTime="scope">
        <div>{{ formatTimestamp(scope.row.createTime) }}</div>
      </template>
      <!-- 退款金额 -->
      <template #returnAmount="scope">
        <div>￥{{ formatReturnAmount(scope.row) }}</div>
      </template>
      <!-- 申请状态 -->
      <template #status="scope">
        <div>{{ formatStatus(scope.row.status) }}</div>
      </template>
      <!-- 处理时间 -->
      <template #handleTime="scope">
        <div>{{ formatTimestamp(scope.row.handleTime) }}</div>
      </template>
      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="openDrawer(scope.row)">查看</el-button>
      </template>
    </ProTable>
    <OrderReturnApplyDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="OrderReturnApplyManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { OrderReturnApply } from "@/api/interface";
import { ref, reactive } from "vue";
import OrderReturnApplyDrawer from "@/views/oms/orderReturnApply/detail.vue";
import { EditPen } from "@element-plus/icons-vue";
import { getOrderReturnAppliesApi } from "@/api/modules/orderReturnApply";
import { formatTimestamp } from "@/utils/time";
import { defaultStatusOptions, formatStatus } from "@/views/oms/orderReturnApply/orderReturnApply";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<OrderReturnApply.OrderReturnApplyModel>[]>([
  { type: "selection", fixed: "left", width: 70 },
  {
    prop: "id",
    label: "服务单号",
    search: { el: "input" }
  },
  {
    prop: "createTime",
    label: "申请时间"
  },
  {
    prop: "memberUsername",
    label: "用户账号"
  },
  {
    prop: "returnAmount",
    label: "退款金额"
  },
  {
    prop: "status",
    label: "申请状态",
    enum: defaultStatusOptions,
    search: { el: "select", props: { filterable: true } }
  },
  {
    prop: "handleTime",
    label: "处理时间"
  },
  { prop: "operation", label: "操作", fixed: "right", width: 170 }
]);

// 退款金额
const formatReturnAmount = (row: OrderReturnApply.OrderReturnApplyModel) => {
  return row.productRealPrice * row.productCount;
};

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
  return getOrderReturnAppliesApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof OrderReturnApplyDrawer> | null>(null);
const openDrawer = (row: Partial<OrderReturnApply.OrderReturnApplyModel> = {}) => {
  const params = {
    row: { ...row }
  };
  drawerRef.value?.acceptParams(params);
};
</script>
