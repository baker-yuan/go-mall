<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}品牌`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="品牌名称" prop="name">
        <el-input v-model="drawerProps.row!.name" placeholder="请填写品牌名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="品牌首字母" prop="firstLetter">
        <el-input v-model="drawerProps.row!.firstLetter" placeholder="请填写品牌首字母" clearable></el-input>
      </el-form-item>
      <el-form-item label="品牌logo" prop="logo">
        <UploadImg v-model:image-url="drawerProps.row!.logo" width="135px" height="135px" :file-size="3">
          <template #empty>
            <el-icon><Picture /></el-icon>
            <span>请上传logo</span>
          </template>
          <template #tip> 只能上传jpg/png文件，且不超过10MB </template>
        </UploadImg>
      </el-form-item>
      <el-form-item label="品牌专区大图" prop="bigPic">
        <UploadImg v-model:image-url="drawerProps.row!.bigPic" width="135px" height="135px" :file-size="3">
          <template #empty>
            <el-icon><Picture /></el-icon>
            <span>请上传logo</span>
          </template>
          <template #tip> 只能上传jpg/png文件，且不超过10MB </template>
        </UploadImg>
      </el-form-item>
      <el-form-item label="品牌故事" prop="brandStory">
        <el-input type="textarea" v-model="drawerProps.row!.brandStory" placeholder="请填写品牌故事" clearable></el-input>
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
      <el-form-item label="品牌制造商" prop="factoryStatus">
        <el-radio-group v-model="drawerProps.row!.factoryStatus">
          <el-radio :label="1">是</el-radio>
          <el-radio :label="0">否</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="BrandDrawer">
import { ref, reactive } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import UploadImg from "@/components/Upload/Img.vue";
import { Brand } from "@/api/interface";

const rules = reactive({
  name: [{ required: true, message: "请填写品牌名称" }],
  logo: [{ required: true, message: "请上传品logo" }],
  firstLetter: [{ required: true, message: "请填写品牌首字母" }]
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Brand.BrandModel>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}品牌成功！` });
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
