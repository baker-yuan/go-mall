<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="450px" :title="`${drawerProps.title}订单退货申请`">
    <template #footer>
      <el-button @click="drawerVisible = false">取消</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="OrderReturnApplyDrawer">
import { ref } from "vue";
import { OrderReturnApply } from "@/api/interface";

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<OrderReturnApply.OrderReturnApplyModel>; // 数据
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

defineExpose({
  acceptParams
});
</script>
