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
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="CirclePlus" @click="openDrawer('新增')">新增退货原因</el-button>
      </template>
      <!-- 是否可用 -->
      <template #status="scope">
        <el-switch @change="handleStatusChange(scope.row)" :active-value="1" :inactive-value="0" v-model="scope.row.status">
        </el-switch>
      </template>
      <!-- 添加时间 -->
      <template #createTime="scope">
        <div>{{ formatTimestamp(scope.row.createTime) }}</div>
      </template>

      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
        <el-button type="primary" link :icon="Delete" @click="deleteOrderReturnReason(scope.row)">删除</el-button>
      </template>
    </ProTable>
    <OrderReturnReasonDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="OrderReturnReasonManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { OrderReturnReason } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import OrderReturnReasonDrawer from "@/views/oms/returnReason/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import {
  createOrderReturnReasonApi,
  updateOrderReturnReasonApi,
  deleteOrderReturnReasonApi,
  getOrderReturnReasonsApi
} from "@/api/modules/returnReason.ts";
import { formatTimestamp } from "@/utils/time.ts";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<OrderReturnReason.OrderReturnReasonModel>[]>([
  {
    prop: "id",
    label: "编号",
    search: { el: "input" }
  },
  {
    prop: "name",
    label: "原因类型"
  },
  {
    prop: "sort",
    label: "排序"
  },
  {
    prop: "status",
    label: "是否可用"
  },
  {
    prop: "createTime",
    label: "添加时间"
  },
  { prop: "operation", label: "操作", fixed: "right", width: 170 }
]);

// 是否显示改变
const handleStatusChange = (row: OrderReturnReason.OrderReturnReasonModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateOrderReturnReasonApi(newParams);
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
  return getOrderReturnReasonsApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除退货原因
const deleteOrderReturnReason = async (row: OrderReturnReason.OrderReturnReasonModel) => {
  await useHandleData(deleteOrderReturnReasonApi, row.id, `删除【${row.name}】退货原因`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof OrderReturnReasonDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<OrderReturnReason.OrderReturnReasonModel> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createOrderReturnReasonApi : title === "编辑" ? updateOrderReturnReasonApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
