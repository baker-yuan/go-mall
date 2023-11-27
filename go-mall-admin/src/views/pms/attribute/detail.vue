<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="700px" :title="`${drawerProps.title}商品属性参数`">
    <el-form
      ref="ruleFormRef"
      label-width="140px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="属性名称" prop="name">
        <el-input v-model="drawerProps.row!.name"></el-input>
      </el-form-item>

      <el-form-item label="商品类型">
        <el-select v-model="drawerProps.row!.productAttributeCategoryId" placeholder="请选择">
          <el-option v-for="(item, index) in productAttributeCategoryData" :key="index" :label="item.name" :value="item.id">
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="分类筛选样式">
        <el-radio-group v-model="drawerProps.row!.filterType">
          <el-radio :label="0">普通</el-radio>
          <el-radio :label="1">颜色</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="能否进行检索">
        <el-radio-group v-model="drawerProps.row!.searchType">
          <el-radio :label="0">不需要检索</el-radio>
          <el-radio :label="1">关键字检索</el-radio>
          <el-radio :label="2">范围检索</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="商品属性关联">
        <el-radio-group v-model="drawerProps.row!.relatedStatus">
          <el-radio :label="1">是</el-radio>
          <el-radio :label="0">否</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="属性是否可选">
        <el-radio-group v-model="drawerProps.row!.selectType">
          <el-radio :label="0">唯一</el-radio>
          <el-radio :label="1">单选</el-radio>
          <el-radio :label="2">复选</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="属性值的录入方式">
        <el-radio-group v-model="drawerProps.row!.inputType">
          <el-radio :label="0">手工录入</el-radio>
          <el-radio :label="1">从下面列表中选择</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="属性值可选值列表">
        <comma-input v-model="drawerProps.row!.inputList"></comma-input>
      </el-form-item>
      <!--
      <el-form-item label="属性值可选值列表">
        <el-input type="textarea" v-model="drawerProps.row!.inputList"></el-input>
      </el-form-item>
      -->
      <el-form-item label="是否支持手动新增">
        <el-radio-group v-model="drawerProps.row!.handAddStatus">
          <el-radio :label="1">是</el-radio>
          <el-radio :label="0">否</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="排序属性">
        <el-input-number v-model="drawerProps.row!.sort" placeholder="请填写排序" clearable></el-input-number>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="ProductAttributeDrawer">
import { ref, reactive, onBeforeMount } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { ProductAttribute, ProductAttributeCategory } from "@/api/interface";
import { getProductAttributeCategoriesApi } from "@/api/modules/productAttributeCategory";
import CommaInput from "@/components/CommaInput/index.vue";

const rules = reactive({
  name: [{ required: true, message: "请填写属性名称" }]
});

const productAttributeCategoryData = ref<ProductAttributeCategory.ProductAttributeCategoryModel[]>([]);
onBeforeMount(async () => {
  const { data } = await getProductAttributeCategoriesApi({
    pageNum: 1,
    pageSize: 1000,
    name: ""
  });
  productAttributeCategoryData.value = data.data;
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<ProductAttribute.ProductAttributeModel>; // 数据
  api?: (params: any) => Promise<any>; // 提交api
  getTableList?: () => void; // 刷新列表
}

// 是否显示弹框
const drawerVisible = ref(false);
// 参数
const drawerProps = ref<DrawerProps>({
  isView: false,
  title: "",
  row: {}
});

// 接收父组件传过来的参数
const acceptParams = (params: DrawerProps) => {
  drawerProps.value = params;
  drawerVisible.value = true;
};

// 提交数据（新增/编辑）
const ruleFormRef = ref<FormInstance>();
const handleSubmit = () => {
  ruleFormRef.value!.validate(async valid => {
    if (!valid) return;
    try {
      await drawerProps.value.api!(drawerProps.value.row);
      ElMessage.success({ message: `${drawerProps.value.title}商品属性参数成功！` });
      drawerProps.value.getTableList!();
      drawerVisible.value = false;
    } catch (error) {
      console.log(error);
    }
  });
};

defineExpose({
  acceptParams
});
</script>
