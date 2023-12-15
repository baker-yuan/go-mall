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
        <el-button type="primary" :icon="CirclePlus" @click="openDrawer('新增')">新增</el-button>
      </template>

      <!-- 广告位置 -->
      <template #type="scope">
        {{ formatType(scope.row.type) }}
      </template>

      <!-- 广告图片 -->
      <template #pic="scope">
        <img style="height: 80px; width: 100px" :src="scope.row.pic" />
      </template>
      <!-- 时间 -->
      <template #timeRange="scope">
        <p>开始时间：{{ formatTimestamp(scope.row.startTime) }}</p>
        <p>到期时间：{{ formatTimestamp(scope.row.endTime) }}</p>
      </template>

      <!-- 上线/下线 -->
      <template #status="scope">
        <el-switch @change="handleUpdateStatus(scope.row)" :active-value="1" :inactive-value="0" v-model="scope.row.status">
        </el-switch>
      </template>

      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
        <el-button type="primary" link :icon="Delete" @click="deleteHomeAdvertise(scope.row)">删除</el-button>
      </template>
    </ProTable>
    <HomeAdvertiseDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="HomeAdvertiseManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { HomeAdvertise } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import HomeAdvertiseDrawer from "@/views/sms/advertise/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import {
  createHomeAdvertiseApi,
  updateHomeAdvertiseApi,
  deleteHomeAdvertiseApi,
  getHomeAdvertisesApi
} from "@/api/modules/advertise.ts";
import { formatType, typeOptions } from "@/views/sms/advertise/index";
import { formatTimestamp } from "@/utils/time";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<HomeAdvertise.HomeAdvertiseModel>[]>([
  { type: "selection", fixed: "left", width: 70 },
  {
    prop: "id",
    label: "编号"
  },
  {
    prop: "name",
    label: "广告名称",
    search: { el: "input" }
  },
  {
    prop: "type",
    label: "广告位置",
    enum: typeOptions,
    search: { el: "select", props: { filterable: true } }
  },
  {
    prop: "pic",
    label: "广告图片",
    showOverflowTooltip: false
  },
  {
    prop: "timeRange",
    label: "时间",
    showOverflowTooltip: false,
    width: 300
  },
  {
    prop: "status",
    label: "上线/下线"
  },
  {
    prop: "clickCount",
    label: "点击次数"
  },
  {
    prop: "orderCount",
    label: "下单数"
  },
  {
    prop: "endTime",
    label: "到期时间",
    isShow: false,
    search: { el: "date-picker", props: { filterable: true } }
  },

  { prop: "operation", label: "操作", fixed: "right", width: 170 }
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
  return getHomeAdvertisesApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除首页轮播广告表
const deleteHomeAdvertise = async (row: HomeAdvertise.HomeAdvertiseModel) => {
  await useHandleData(deleteHomeAdvertiseApi, row.id, `删除【${row.name}】首页轮播广告表`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof HomeAdvertiseDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<HomeAdvertise.HomeAdvertiseModel> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createHomeAdvertiseApi : title === "编辑" ? updateHomeAdvertiseApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};

// 上线/下线
const handleUpdateStatus = (row: HomeAdvertise.HomeAdvertiseModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateHomeAdvertiseApi(newParams);
};
</script>
