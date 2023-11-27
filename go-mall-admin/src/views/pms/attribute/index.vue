<template>
  <div class="main-box">
    <div class="card filter">
      <h4 class="title sle">商品类型</h4>
      <el-input placeholder="输入关键字进行过滤" clearable />
      <el-scrollbar :style="{ height: `calc(100% - 95px)` }">
        <el-tree
          ref="treeRef"
          node-key="id"
          :highlight-current="true"
          :check-strictly="false"
          :expand-on-click-node="false"
          :data="treeData"
          :props="defaultProps"
          @node-click="handleNodeClick"
        ></el-tree>
      </el-scrollbar>
    </div>
    <div class="table-box">
      <el-tabs class="card" @tab-click="handleTabClick">
        <el-tab-pane v-for="(item, index) in tabPane" :key="index" :label="item.label" :name="item.name">
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
            <!-- 属性是否可选 -->
            <template #selectType="scope">
              {{ selectTypeFilter(scope.row.selectType) }}
            </template>
            <!-- 属性值的录入方式 -->
            <template #inputType="scope">
              {{ inputTypeFilter(scope.row.inputType) }}
            </template>

            <!-- 表格操作 -->
            <template #operation="scope">
              <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
              <el-button type="primary" link :icon="Delete" @click="deleteProductAttribute(scope.row)">删除</el-button>
            </template>
          </ProTable>
        </el-tab-pane>
      </el-tabs>
      <ProductAttributeDrawer ref="drawerRef" />
    </div>
  </div>
</template>

<script setup lang="ts" name="ProductAttributeManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { ProductAttribute, ProductAttributeCategory } from "@/api/interface";
import { ref, reactive, onBeforeMount } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import ProductAttributeDrawer from "@/views/pms/attribute/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import {
  createProductAttributeApi,
  updateProductAttributeApi,
  deleteProductAttributeApi,
  getProductAttributesApi
} from "@/api/modules/attribute";
import { getProductAttributeCategoriesApi } from "@/api/modules/productAttributeCategory";
import { ElTree } from "element-plus";

const tabPane = [
  {
    label: "属性列表",
    name: "0"
  },
  {
    label: "参数列表",
    name: "1"
  }
];

const defaultProps = {
  children: "children",
  label: "name"
};

const currentNode = ref<ProductAttributeCategory.ProductAttributeCategoryModel>();

const treeData = ref<ProductAttributeCategory.ProductAttributeCategoryModel[]>([]);
onBeforeMount(async () => {
  const { data } = await getProductAttributeCategoriesApi({
    pageNum: 1,
    pageSize: 1000,
    name: ""
  });
  treeData.value = data.data;
});

const handleTabClick = (tab: any) => {
  initParam.type = tab.props.name;
};

const handleNodeClick = (data: ProductAttributeCategory.ProductAttributeCategoryModel) => {
  initParam.productAttributeCategoryId = data.id;
  currentNode.value = data;
};

const selectTypeFilter = (value: number) => {
  if (value === 1) {
    return "单选";
  } else if (value === 2) {
    return "多选";
  } else {
    return "唯一";
  }
};

const inputTypeFilter = (value: number) => {
  if (value === 1) {
    return "从列表中选取";
  } else {
    return "手工录入";
  }
};

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<ProductAttribute.ProductAttributeModel>[]>([
  { prop: "id", label: "编号" },
  {
    prop: "name",
    label: "属性名称",
    search: { el: "input" }
  },
  { prop: "selectType", label: "属性是否可选" },
  { prop: "inputType", label: "属性值录入方式" },
  { prop: "inputList", label: "可选值列表" },
  { prop: "sort", label: "排序" },
  { prop: "operation", label: "操作", fixed: "right", width: 170 }
]);

// 如果表格需要初始化请求参数，直接定义传给 ProTable (之后每次请求都会自动带上该参数，此参数更改之后也会一直带上，改变此参数会自动刷新表格数据)
const initParam = reactive({ productAttributeCategoryId: 0, type: 0 });

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
  return getProductAttributesApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除商品属性参数表
const deleteProductAttribute = async (row: ProductAttribute.ProductAttributeModel) => {
  await useHandleData(deleteProductAttributeApi, row.id, `删除【${row.name}】商品属性参数表`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof ProductAttributeDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<ProductAttribute.ProductAttributeModel> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createProductAttributeApi : title === "编辑" ? updateProductAttributeApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>

<style scoped lang="scss">
@import "./index.scss";
</style>
