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
        <el-button type="primary" :icon="CirclePlus" @click="openDrawer('新增')">新增配置</el-button>
      </template>

      <!-- 表格操作 -->
      <template #operation="scope">
        <el-row :gutter="10">
          <el-col :span="24">
            <el-button
              style="display: block; margin-bottom: 10px"
              type="primary"
              link
              :icon="EditPen"
              @click="openEditJsonDialog(scope.row)"
            >
              编辑内容
            </el-button>
          </el-col>
          <el-col :span="24">
            <el-button
              style="display: block; margin-bottom: 10px"
              type="primary"
              link
              :icon="EditPen"
              @click="openDrawer('编辑', scope.row)"
              >编辑约束
            </el-button>
          </el-col>
          <el-col :span="24">
            <el-button
              style="display: block; margin-bottom: 10px"
              type="primary"
              link
              :icon="EditPen"
              @click="openDrawer('编辑', scope.row)"
              >编辑
            </el-button>
          </el-col>
          <el-col :span="24">
            <el-button
              style="display: block; margin-bottom: 10px"
              type="primary"
              link
              :icon="Delete"
              @click="deleteJsonDynamicConfig(scope.row)"
              >删除
            </el-button>
          </el-col>
        </el-row>
      </template>
    </ProTable>
    <JsonDynamicConfigDrawer ref="drawerRef" />

    <!-- 编辑json -->
    <el-dialog v-model="showEditJsonContent">
      <VueForm
        @submit="handlerContentFormSubmit"
        @cancel="handlerContentFormCancel"
        v-model="contentFormData"
        :schema="contentSchema"
      ></VueForm>
    </el-dialog>
  </div>
</template>

<script setup lang="ts" name="JsonDynamicConfigManage">
import VueForm from "@lljj/vue3-form-element";
import { ProTableInstance, ColumnProps } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import { JsonDynamicConfig } from "@/api/interface";
import { ref, reactive } from "vue";
import { useHandleData } from "@/hooks/useHandleData";
import JsonDynamicConfigDrawer from "@/views/cms/jsonConfig/detail.vue";
import { CirclePlus, Delete, EditPen } from "@element-plus/icons-vue";
import {
  createJsonDynamicConfigApi,
  updateJsonDynamicConfigApi,
  deleteJsonDynamicConfigApi,
  getJsonDynamicConfigsApi
} from "@/api/modules/jsonConfig";

// ProTable 实例
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns = reactive<ColumnProps<JsonDynamicConfig.JsonDynamicConfigModel>[]>([
  {
    prop: "bizType",
    label: "业务类型",
    showOverflowTooltip: false,
    search: { el: "input" }
  },
  {
    prop: "bizDesc",
    label: "业务描述",
    search: { el: "input" }
  },
  {
    prop: "content",
    label: "内容",
    search: { el: "input" }
  },
  {
    prop: "jsonSchema",
    label: "json内容约束",
    search: { el: "input" }
  },
  { prop: "operation", label: "操作", fixed: "right", width: 170 }
]);

// 展示编辑json
const showEditJsonContent = ref(false);
const contentFormData = ref<any>({});
const contentSchema = ref<any>({});

const openEditJsonDialog = (row: Partial<JsonDynamicConfig.JsonDynamicConfigModel> = {}) => {
  if (!row) {
    return;
  }
  console.log(row.content, row.jsonSchema);
  showEditJsonContent.value = true;
  if (row.jsonSchema) {
    contentSchema.value = JSON.parse(row.jsonSchema);
    console.log(contentSchema.value);
  }
  if (row.content) {
    contentFormData.value = JSON.parse(row.content);
    console.log(contentFormData.value);
  }
};

const handlerContentFormSubmit = () => {
  showEditJsonContent.value = false;
};

const handlerContentFormCancel = () => {
  showEditJsonContent.value = false;
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
  return getJsonDynamicConfigsApi(newParams);
};

// 表格拖拽排序
const sortTable = ({ newIndex, oldIndex }: { newIndex?: number; oldIndex?: number }) => {
  console.log(newIndex, oldIndex);
};

// 删除JSON动态配置
const deleteJsonDynamicConfig = async (row: JsonDynamicConfig.JsonDynamicConfigModel) => {
  await useHandleData(deleteJsonDynamicConfigApi, row.id, `删除【${row.bizType}】JSON动态配置`);
  proTable.value?.getTableList();
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof JsonDynamicConfigDrawer> | null>(null);
const openDrawer = (title: string, row: Partial<JsonDynamicConfig.JsonDynamicConfigModel> = {}) => {
  const params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title === "新增" ? createJsonDynamicConfigApi : title === "编辑" ? updateJsonDynamicConfigApi : undefined,
    getTableList: proTable.value?.getTableList
  };
  drawerRef.value?.acceptParams(params);
};
</script>
