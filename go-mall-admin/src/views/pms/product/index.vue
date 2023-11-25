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
        <el-button type="primary" :icon="CirclePlus" @click="openDrawer('新增')">新增商品</el-button>
      </template>

      <!-- 商品图片 -->
      <template #pic="scope">
        <!-- eslint-disable-next-line prettier/prettier -->
        <img style="height: 80px" :src="scope.row.pic">
      </template>

      <!-- 商品名称 -->
      <template #nameAndBrand="scope">
        <p>{{ scope.row.name }}</p>
        <p>品牌：{{ scope.row.brandName }}</p>
      </template>

      <!-- 价格/货号 -->
      <template #priceAndSn="scope">
        <p>价格：￥{{ scope.row.price }}</p>
        <p>货号：{{ scope.row.productSN }}</p>
      </template>

      <!-- 标签 -->
      <template #setupState="scope">
        <p>
          上架：
          <el-switch
            @change="handlePublishStatusChange(scope.row)"
            :active-value="1"
            :inactive-value="0"
            v-model="scope.row.publishStatus"
          >
          </el-switch>
        </p>
        <p>
          新品：
          <el-switch
            @change="handleNewStatusChange(scope.row)"
            :active-value="1"
            :inactive-value="0"
            v-model="scope.row.newStatus"
          >
          </el-switch>
        </p>
        <p>
          推荐：
          <el-switch
            @change="handleRecommendStatusChange(scope.row)"
            :active-value="1"
            :inactive-value="0"
            v-model="scope.row.recommandStatus"
          >
          </el-switch>
        </p>
      </template>

      <!-- SKU库存 -->
      <template #setupSku="scope">
        <el-button type="primary" :icon="EditPen" @click="handleShowSkuEditDialog(scope.row)" circle></el-button>
      </template>

      <!-- 审核状态 -->
      <template #verifyStatus="scope">
        <p>{{ verifyStatusFilter(scope.row.verifyStatus) }}</p>
        <p>
          <el-button type="text" @click="handleShowVerifyDetail(scope.row)">审核详情</el-button>
        </p>
      </template>

      <!-- 表格操作 -->
      <template #operation="scope">
        <el-button type="primary" link :icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
        <el-button type="primary" link :icon="Delete" @click="deleteProduct(scope.row)">删除</el-button>
      </template>
    </ProTable>
    <ProductDrawer ref="drawerRef" />
  </div>
</template>

<script setup lang="ts" name="ProductManage">
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { Product } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import ProductDrawer from "@/views/pms/brand/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import { createProductApi, updateProductApi, deleteProductApi, getProductsApi } from "@/api/modules/product";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<Product.ProductModel>[]>([
  { type: "selection", width: 60, fixed: "left" },
  {
    prop: "id",
    width: 100,
    label: "编号",
    search: { el: "input" }
  },
  {
    prop: "pic",
    width: 120,
    label: "商品图片"
  },
  {
    prop: "nameAndBrand",
    width: 200,
    label: "商品名称"
  },
  {
    prop: "priceAndSn",
    width: 120,
    label: "价格/货号"
  },
  {
    prop: "setupState",
    width: 140,
    label: "标签"
  },
  {
    prop: "sort",
    width: 100,
    label: "排序"
  },
  {
    prop: "setupSku",
    width: 100,
    label: "SKU库存"
  },
  {
    prop: "sale",
    width: 100,
    label: "销量"
  },
  {
    prop: "verifyStatus",
    width: 100,
    label: "审核状态"
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
  return getProductsApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

const verifyStatusFilter = (value: number) => {
  if (value === 1) {
    return "审核通过";
  } else {
    return "未审核";
  }
};

const handleShowVerifyDetail = (row: Product.ProductModel) => {
  console.log(row);
};

// 删除商品
const deleteProduct = async (row: Product.ProductModel) => {
  await useHandleData(deleteProductApi, row.id, `删除【${row.name}】商品`);
  proTable.value?.getTableList();
};

// 上架切换
const handlePublishStatusChange = (row: Product.ProductModel) => {
  console.log(row);
};

// 新品切换
const handleNewStatusChange = (row: Product.ProductModel) => {
  console.log(row);
};

// 推荐切换
const handleRecommendStatusChange = (row: Product.ProductModel) => {
  console.log(row);
};

// sku库存
const handleShowSkuEditDialog = (row: Product.ProductModel) => {
  console.log(row);
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof ProductDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<Product.ProductModel> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createProductApi : title === "编辑" ? updateProductApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
