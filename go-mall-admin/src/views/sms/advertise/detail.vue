<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}首页轮播广告表`">
    <el-form
      ref="ruleFormRef"
      label-width="100px"
      label-suffix=" :"
      :rules="rules"
      :disabled="drawerProps.isView"
      :model="drawerProps.row"
      :hide-required-asterisk="drawerProps.isView"
    >
      <el-form-item label="广告名称" prop="name">
        <el-input v-model="drawerProps.row!.name" placeholder="请填写广告名称" clearable></el-input>
      </el-form-item>
      <el-form-item label="广告位置">
        <el-select v-model="drawerProps.row!.type">
          <el-option v-for="type in typeOptions" :key="type.value" :label="type.label" :value="type.value"> </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="开始时间" prop="startTime">
        <el-date-picker type="datetime" placeholder="选择日期" v-model="startTimeDate"></el-date-picker>
      </el-form-item>
      <el-form-item label="到期时间" prop="endTime">
        <el-date-picker type="datetime" placeholder="选择日期" v-model="endTimeDate"></el-date-picker>
      </el-form-item>
      <el-form-item label="上线/下线">
        <el-radio-group v-model="drawerProps.row!.status">
          <el-radio :label="0">下线</el-radio>
          <el-radio :label="1">上线</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="广告图片">
        <UploadImg v-model:image-url="drawerProps.row!.pic" width="200px" height="100px" :file-size="3">
          <template #empty>
            <el-icon><Picture /></el-icon>
            <span>请上传广告图片</span>
          </template>
          <template #tip> 只能上传jpg/png文件，且不超过10MB </template>
        </UploadImg>
      </el-form-item>
      <el-form-item label="排序">
        <el-input-number v-model="drawerProps.row!.sort"></el-input-number>
      </el-form-item>
      <el-form-item label="广告链接" prop="url">
        <el-input v-model="drawerProps.row!.url"></el-input>
      </el-form-item>
      <el-form-item label="广告备注">
        <el-input type="textarea" :rows="5" placeholder="请输入内容" v-model="drawerProps.row!.note"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="!drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="HomeAdvertiseDrawer">
import { ref, reactive, computed, Ref } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { HomeAdvertise } from "@/api/interface";
import { typeOptions } from "@/views/sms/advertise/index";
import UploadImg from "@/components/Upload/Img.vue";

const rules = reactive({
  name: [
    { required: true, message: "请输入广告名称", trigger: "blur" },
    { min: 2, max: 140, message: "长度在 2 到 140 个字符", trigger: "blur" }
  ],
  url: [{ required: true, message: "请输入广告链接", trigger: "blur" }],
  startTime: [{ required: true, message: "请选择开始时间", trigger: "blur" }],
  endTime: [{ required: true, message: "请选择到期时间", trigger: "blur" }],
  pic: [{ required: true, message: "请选择广告图片", trigger: "blur" }]
});

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<HomeAdvertise.HomeAdvertiseModel>; // 数据
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
      ElMessage.success({ message: `${drawerProps.value.title}首页轮播广告表成功！` });
      drawerProps.value.getTableList!();
      drawerVisible.value = false;
    } catch (error) {
      console.log(error);
    }
  });
};

// 开始时间
const startTimeDate: Ref<Date> = computed({
  get: () => {
    const timestamp = drawerProps.value.row?.startTime ?? Math.floor(Date.now() / 1000);
    return timestamp === 0 ? new Date() : new Date(timestamp * 1000);
  },
  set: (value: Date) => {
    drawerProps.value.row.startTime = Math.floor(value.getTime() / 1000);
  }
});
// 到期时间
const endTimeDate: Ref<Date> = computed({
  get: () => {
    const timestamp = drawerProps.value.row?.endTime ?? Math.floor(Date.now() / 1000);
    return timestamp === 0 ? new Date() : new Date(timestamp * 1000);
  },
  set: (value: Date) => {
    drawerProps.value.row.endTime = Math.floor(value.getTime() / 1000);
  }
});

defineExpose({
  acceptParams
});
</script>
