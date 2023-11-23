<template>
  <div class="table-box">
    <ProTable
      ref="proTable"
      row-key="id"
      :columns="columns"
      :request-api="getTableList"
      :init-param="initParam"
      :data-callback="dataCallback"
      @darg-sort="sortTable"
      :border="false"
      :is-lazy-load="true"
      :lazy-load-api="lazyLoadTableList"
    >
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" :icon="CirclePlus" @click="openDrawer('新增')">新增分类</el-button>
      </template>

      <!-- 级别 -->
      <template #level="scope">
        <span v-if="scope.row.level === 0">一级</span>
        <span v-else>二级</span>
      </template>

      <!-- 是否进行显示 -->
      <template #showStatus="scope">
        <el-switch
          @change="handleShowStatusChange(scope.row)"
          :active-value="1"
          :inactive-value="0"
          v-model="scope.row.showStatus"
        >
        </el-switch>
      </template>
      <!-- 是否在导航栏显示 -->
      <template #navStatus="scope">
        <el-switch @change="handleNavStatusChange(scope.row)" :active-value="1" :inactive-value="0" v-model="scope.row.navStatus">
        </el-switch>
      </template>

      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
        <el-button type="primary" link :icon="Delete" @click="deleteCategory(scope.row)">删除</el-button>
      </template>
    </ProTable>
    <CategoryDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="CategoryManage">
import ProTable from "@/components/ProTable/index.vue";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import { Category } from "@/api/interface";
import { ref, reactive } from "vue";
import { getCategories, createCategory, updateCategory, deleteCategoryApi } from "@/api/modules/category";
import { useHandleData } from "@/hooks/useHandleData";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import CategoryDrawer from "@/views/pms/category/detail.vue";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<Category.CategoryModel>[]>([
  {
    prop: "id",
    label: "编号",
    search: { el: "input" }
  },
  {
    prop: "name",
    label: "分类名称",
    search: { el: "input" }
  },
  { prop: "level", label: "级别" },
  { prop: "productCount", label: "商品数量" },
  { prop: "productUnit", label: "商品单位" },
  { prop: "navStatus", label: "导航栏" },
  { prop: "showStatus", label: "是否显示" },
  { prop: "sort", label: "排序" },
  {
    prop: "setup",
    label: "设置",
    width: 180
  },
  { prop: "operation", label: "操作", fixed: "right", width: 170 }
]);

// 如果表格需要初始化请求参数，直接定义传给 ProTable (之后每次请求都会自动带上该参数，此参数更改之后也会一直带上，改变此参数会自动刷新表格数据)
const initParam = reactive({ parentId: 0 });

// dataCallback 是对于返回的表格数据做处理，如果你后台返回的数据不是 list && total && pageNum && pageSize 这些字段，可以在这里进行处理成这些字段
// 或者直接去 hooks/useTable.ts 文件中把字段改为你后端对应的就行
const dataCallback = (data: any) => {
  // https://element-plus.org/zh-CN/component/table.html#%E6%A0%91%E5%BD%A2%E6%95%B0%E6%8D%AE%E4%B8%8E%E6%87%92%E5%8A%A0%E8%BD%BD
  // 设置默认都有字节点
  for (let i = 0; i < data.data.length; i++) {
    data.data[i].hasChildren = true;
  }
  return {
    list: data.data,
    total: data.pageTotal,
    pageNum: data.pageNum,
    pageSize: data.pageSize
  };
};

// 如果你想在请求之前对当前请求参数做一些操作，可以自定义如下函数：params 为当前所有的请求参数（包括分页），最后返回请求列表接口
// 默认不做操作就直接在 ProTable 组件上绑定	:requestApi="getUserList"
const getTableList = (params: any) => {
  let newParams = JSON.parse(JSON.stringify(params));
  console.log(params, newParams);
  return getCategories(newParams);
};

const lazyLoadTableList = (
  row: Category.CategoryModel,
  treeNode: unknown,
  resolve: (categories: Category.CategoryModel[]) => void
) => {
  getCategories({
    pageNum: 1,
    pageSize: 10000,
    parentId: row.id
  })
    .then(data => {
      resolve(data.data.data);
    })
    .catch(error => {
      console.error(error);
    });
};

// 是否显示改变
const handleShowStatusChange = (row: Category.CategoryModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateCategory(newParams);
};

// 是否在导航栏显示改变
const handleNavStatusChange = (row: Category.CategoryModel) => {
  if (row.id == 0 || row.id == undefined) {
    return;
  }
  let newParams = JSON.parse(JSON.stringify(row));
  updateCategory(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除分类
const deleteCategory = async (row: Category.CategoryModel) => {
  await useHandleData(deleteCategoryApi, row.id, `删除【${row.name}】分类`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof CategoryDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<Category.CategoryModel> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createCategory : title === "编辑" ? updateCategory : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
