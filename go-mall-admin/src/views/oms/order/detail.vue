<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="50%" :title="`${drawerProps.title}订单`">
    <div class="detail-container">
      <div>
        <el-steps :active="formatStepStatus(orderDetail!.status)" finish-status="success" align-center>
          <el-step title="提交订单" :description="formatTimestamp(orderDetail!.createTime)"></el-step>
          <el-step title="支付订单" :description="formatTimestamp(orderDetail!.paymentTime)"></el-step>
          <el-step title="平台发货" :description="formatTimestamp(orderDetail!.deliveryTime)"></el-step>
          <el-step title="确认收货" :description="formatTimestamp(orderDetail!.receiveTime)"></el-step>
          <el-step title="完成评价" :description="formatTimestamp(orderDetail!.commentTime)"></el-step>
        </el-steps>
      </div>
    </div>
    <template #footer>
      <el-button @click="drawerVisible = false">关闭</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="OrderDrawer">
import { ref } from "vue";
import { Order } from "@/api/interface";
import { getOrderSyncApi } from "@/api/modules/order";
import { formatTimestamp } from "@/utils/time";

// 格式化状态
const formatStepStatus = (value: number) => {
  if (value === 1) {
    //待发货
    return 2;
  } else if (value === 2) {
    //已发货
    return 3;
  } else if (value === 3) {
    //已完成
    return 4;
  } else {
    //待付款、已关闭、无限订单
    return 1;
  }
};

// 订单详情
const orderDetail = ref<Order.OrderModel | null>(null);

interface DrawerProps {
  title: string; // 标题
  row: Partial<Order.OrderModel>; // 数据
}

// 是否显示弹框
const drawerVisible = ref(false);
// 参数
const drawerProps = ref<DrawerProps>({
  title: "",
  row: {}
});

// 接收父组件传过来的参数
const acceptParams = async (params: OrderProps) => {
  drawerProps.value = params;
  drawerVisible.value = true;
  if (params.row.id && params.row.id !== 0) {
    orderDetail.value = await getOrderSyncApi(drawerProps.value.row.id);
  }
};

defineExpose({
  acceptParams
});
</script>

<style scoped lang="scss">
.detail-container {
  width: 80%;
  padding: 20px 20px 20px 20px;
  margin: 20px auto;
}

.operate-container {
  background: #f2f6fc;
  height: 80px;
  margin: -20px -20px 0;
  line-height: 80px;
}

.operate-button-container {
  float: right;
  margin-right: 20px;
}

.table-layout {
  margin-top: 20px;
  border-left: 1px solid #dcdfe6;
  border-top: 1px solid #dcdfe6;
}

.table-cell {
  height: 60px;
  line-height: 40px;
  border-right: 1px solid #dcdfe6;
  border-bottom: 1px solid #dcdfe6;
  padding: 10px;
  font-size: 14px;
  color: #606266;
  text-align: center;
  overflow: hidden;
}

.table-cell-title {
  border-right: 1px solid #dcdfe6;
  border-bottom: 1px solid #dcdfe6;
  padding: 10px;
  background: #f2f6fc;
  text-align: center;
  font-size: 14px;
  color: #303133;
}
</style>
