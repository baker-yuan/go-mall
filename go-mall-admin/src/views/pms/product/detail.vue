<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="900px" :title="`${drawerProps.title}商品表`">
    <el-steps :active="activeStep" finish-status="success" align-center>
      <el-step title="填写商品信息"></el-step>
      <el-step title="填写商品促销"></el-step>
      <el-step title="填写商品属性"></el-step>
      <el-step title="选择商品关联"></el-step>
    </el-steps>
    <div style="margin-top: 50px">
      <el-form
        :model="drawerProps.row"
        :rules="productInfoDetailRules"
        ref="productInfoForm"
        label-width="120px"
        class="form-inner-container"
        size="small"
      >
        <!--
        <el-form-item label="商品分类：" prop="productCategoryId">
          <el-cascader
            v-model="selectProductCateValue"
            :options="productCateOptions">
          </el-cascader>
        </el-form-item>
        -->
        <el-form-item label="商品名称：" prop="name">
          <el-input v-model="drawerProps.row!.name"></el-input>
        </el-form-item>
        <el-form-item label="副标题：" prop="subTitle">
          <el-input v-model="drawerProps.row!.subTitle"></el-input>
        </el-form-item>

        <!--
          <el-form-item label="商品品牌：" prop="brandId">
            <el-select
              v-model="drawerProps.row?.brandId"
              @change="handleBrandChange"
              placeholder="请选择品牌">
              <el-option
                v-for="item in brandOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
        -->
        <el-form-item label="商品介绍：">
          <el-input :auto-size="true" v-model="drawerProps.row!.description" type="textarea" placeholder="请输入内容"></el-input>
        </el-form-item>
        <el-form-item label="商品货号：">
          <el-input v-model="drawerProps.row!.productSn"></el-input>
        </el-form-item>
        <el-form-item label="商品售价：">
          <el-input v-model="drawerProps.row!.price"></el-input>
        </el-form-item>
        <el-form-item label="市场价：">
          <el-input v-model="drawerProps.row!.originalPrice"></el-input>
        </el-form-item>
        <el-form-item label="商品库存：">
          <el-input v-model="drawerProps.row!.stock"></el-input>
        </el-form-item>
        <el-form-item label="计量单位：">
          <el-input v-model="drawerProps.row!.unit"></el-input>
        </el-form-item>
        <el-form-item label="商品重量：">
          <el-input v-model="drawerProps.row!.weight" style="width: 300px"></el-input>
          <span style="margin-left: 20px">克</span>
        </el-form-item>
        <el-form-item label="排序">
          <el-input v-model="drawerProps.row!.sort"></el-input>
        </el-form-item>
        <el-form-item style="text-align: center">
          <el-button type="primary" @click="handleNext('productInfoForm')">下一步，填写商品促销</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-drawer>
</template>

<script setup lang="ts" name="ProductDrawer">
import { reactive, ref } from "vue";
// import { ElMessage, FormInstance } from "element-plus";
import { Product } from "@/api/interface";

const productInfoDetailRules = reactive({
  // name: [
  //   {required: true, message: '请输入商品名称', trigger: 'blur'},
  //   {min: 2, max: 140, message: '长度在 2 到 140 个字符', trigger: 'blur'}
  // ],
  // subTitle: [{required: true, message: '请输入商品副标题', trigger: 'blur'}],
  // productCategoryId: [{required: true, message: '请选择商品分类', trigger: 'blur'}],
  // brandId: [{required: true, message: '请选择商品品牌', trigger: 'blur'}],
  // description: [{required: true, message: '请输入商品介绍', trigger: 'blur'}],
  // requiredProp: [{required: true, message: '该项为必填项', trigger: 'blur'}]
});

// const rules = reactive({});

const activeStep = ref<number>(0);

const handleNext = (formName: string) => {
  console.log(formName);
};

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Product.ProductModel>; // 数据
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
// const ruleFormRef = ref<FormInstance>();
// const handleSubmit = () => {
//   ruleFormRef.value!.validate(async valid => {
//     if (!valid) return;
//     try {
//       await drawerProps.value.api!(drawerProps.value.row);
//       ElMessage.success({ message: `${drawerProps.value.title}商品表成功！` });
//       drawerProps.value.getTableList!();
//       drawerVisible.value = false;
//     } catch (error) {
//       console.log(error);
//     }
//   });
// };

defineExpose({
  acceptParams
});
</script>

<style scoped lang="scss">
.form-container {
  width: 960px;
}
.form-inner-container {
  width: 800px;
}
</style>
