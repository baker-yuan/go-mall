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
      <template #tableHeader="scope">
        <el-button type="primary" :icon="CirclePlus" @click="openDrawer('新增')">新增品牌</el-button>
        <el-button type="warning" :disabled="!scope.isSelected" @click="batchHideBrand(scope.selectedListIds)">
          批量隐藏品牌
        </el-button>
        <el-button type="warning" :disabled="!scope.isSelected" @click="batchShowBrand(scope.selectedListIds)">
          批量显示品牌
        </el-button>
      </template>

      <!-- 品牌制造商 -->
      <template #factoryStatus="scope">
        <el-switch
          @change="handleFactoryStatusChange(scope.row)"
          :active-value="1"
          :inactive-value="0"
          v-model="scope.row.factoryStatus"
        >
        </el-switch>
      </template>
      <!-- 是否显示 -->
      <template #showStatus="scope">
        <el-switch
          @change="handleShowStatusChange(scope.row)"
          :active-value="1"
          :inactive-value="0"
          v-model="scope.row.showStatus"
        >
        </el-switch>
      </template>

      <!-- 相关 -->
      <template #related="scope">
        <div style="text-align: left">商品：{{ scope.row.productCommentCount }} 评价：{{ scope.row.productCount }}</div>
      </template>

      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
        <el-button type="primary" link :icon="Delete" @click="deleteBrand(scope.row)">删除</el-button>
      </template>
    </ProTable>
    <BrandDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="BrandManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { Brand } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import BrandDrawer from "@/views/pms/brand/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import { getBrands, createBrand, updateBrand, deleteBrandApi } from "@/api/modules/brand";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<Brand.BrandModel>[]>([
  { type: "selection", fixed: "left", width: 70 },
  { prop: "id", label: "编号" },
  {
    prop: "name",
    label: "品牌名称",
    search: { el: "input" }
  },
  { prop: "firstLetter", label: "品牌首字母" },
  { prop: "sort", label: "排序" },
  { prop: "factoryStatus", label: "品牌制造商" },
  {
    prop: "showStatus",
    label: "是否显示",
    enum: [
      {
        label: "否",
        value: 0
      },
      {
        label: "是",
        value: 1
      }
    ],
    search: { el: "select", props: { filterable: true } }
  },
  {
    prop: "related",
    label: "相关",
    width: 180
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
  return getBrands(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 批量隐藏品牌
const batchHideBrand = async (id: string[]) => {
  console.log(id);
};

// 品牌制造商改变
const handleFactoryStatusChange = (row: Brand.BrandModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateBrand(newParams);
};

// 是否显示改变
const handleShowStatusChange = (row: Brand.BrandModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateBrand(newParams);
};

// 批量显示品牌
const batchShowBrand = async (id: string[]) => {
  console.log(id);
};

// 删除品牌
const deleteBrand = async (row: Brand.BrandModel) => {
  await useHandleData(deleteBrandApi, row.id, `删除【${row.name}】品牌`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof BrandDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<Brand.BrandModel> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createBrand : title === "编辑" ? updateBrand : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
