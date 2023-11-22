<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="550px" :title="`${drawerProps.title}分类`">
    <el-form
      ref="ruleFormRef"
      label-width="130px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="分类名称" prop="name">
        <el-input v-model="drawerProps.row!.name" placeholder="请填写商品分类名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="上级分类" prop="parentId">
        <el-select v-model="drawerProps.row!.parentId" placeholder="请选择分类">
          <el-option v-for="item in selectProductCateList" :key="item.id" :label="item.name" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="数量单位" prop="productUnit">
        <el-input v-model="drawerProps.row!.productUnit" placeholder="请填写分类单位" clearable></el-input>
      </el-form-item>
      <el-form-item label="排序" prop="sort">
        <el-input-number v-model="drawerProps.row!.sort" placeholder="请填写排序" clearable></el-input-number>
      </el-form-item>
      <el-form-item label="是否显示" prop="showStatus">
        <el-radio-group v-model="drawerProps.row!.showStatus">
          <el-radio :label="1">是</el-radio>
          <el-radio :label="0">否</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="是否显示在导航栏" prop="navStatus">
        <el-radio-group v-model="drawerProps.row!.navStatus">
          <el-radio :label="1">是</el-radio>
          <el-radio :label="0">否</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="分类图标" prop="icon">
        <UploadImg v-model:image-url="drawerProps.row!.icon" width="135px" height="135px" :file-size="3">
          <template #empty>
            <el-icon><Picture /></el-icon>
            <span>请上传分类图标</span>
          </template>
          <template #tip> 只能上传jpg/png文件，且不超过10MB </template>
        </UploadImg>
      </el-form-item>
      <el-form-item label="关键字" prop="keywords">
        <el-input v-model="drawerProps.row!.keywords" placeholder="请填写关键字" clearable></el-input>
      </el-form-item>
      <el-form-item label="分类描述" prop="description">
        <el-input v-model="drawerProps.row!.description" type="textarea" placeholder="请填写商品分类描述" clearable></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="CategoryDrawer">
import { ref, reactive, onMounted } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { Category } from "@/api/interface";
import UploadImg from "@/components/Upload/Img.vue";
import { getCategories } from "@/api/modules/category";

onMounted(async () => {
  initSelectProductCateList();
});

const rules = reactive({
  name: [{ required: true, message: "请填写商品分类名称" }]
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Category.CategoryModel>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}分类成功！` });
      drawerProps.value.getTableList!();
      drawerVisible.value = false;
    } catch (error) {
      console.log(error);
    }
  });
};

const selectProductCateList = ref<Category.CategoryModel[]>([]);

const initSelectProductCateList = () => {
  getCategories({
    pageNum: 1,
    pageSize: 10000,
    parentId: 0
  })
    .then(data => {
      selectProductCateList.value = data.data.data;
    })
    .catch(error => {
      console.error(error);
    });
};

defineExpose({
  acceptParams
});
</script>
