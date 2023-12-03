<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="70%" :title="`${drawerProps.title}订单`">
    <div v-if="orderDetail != null" class="detail-container">
      <div>
        <el-steps :active="formatStepStatus(orderDetail.status)" finish-status="success" align-center>
          <el-step title="提交订单" :description="formatTimestamp(orderDetail.createTime)"></el-step>
          <el-step title="支付订单" :description="formatTimestamp(orderDetail.paymentTime)"></el-step>
          <el-step title="平台发货" :description="formatTimestamp(orderDetail.deliveryTime)"></el-step>
          <el-step title="确认收货" :description="formatTimestamp(orderDetail.receiveTime)"></el-step>
          <el-step title="完成评价" :description="formatTimestamp(orderDetail.commentTime)"></el-step>
        </el-steps>
      </div>
      <el-card shadow="never" style="margin-top: 15px">
        <div class="operate-container color-danger">
          <el-row>
            <el-col :span="12">
              <div style="display: flex; align-items: center">
                <el-icon class="el-icon-warning" style="margin-left: 20px">
                  <WarningFilled />
                </el-icon>
                <span class="color-danger">当前订单状态：{{ formatStatus(orderDetail.status) }}</span>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="operate-button-container" v-show="orderDetail.status === 0">
                <el-button size="small" @click="showUpdateReceiverDialog">修改收货人信息</el-button>
                <el-button size="small" @click="showUpdateMoneyDialog">修改费用信息</el-button>
                <el-button size="small" @click="showMessageDialog">发送站内信</el-button>
                <el-button size="small" @click="showCloseOrderDialog">关闭订单</el-button>
                <el-button size="small" @click="showMarkOrderDialog">备注订单</el-button>
              </div>
              <div class="operate-button-container" v-show="orderDetail.status === 1">
                <el-button size="small" @click="showUpdateReceiverDialog">修改收货人信息</el-button>
                <el-button size="small" @click="showMessageDialog">发送站内信</el-button>
                <el-button size="small">取消订单</el-button>
                <el-button size="small" @click="showMarkOrderDialog">备注订单</el-button>
              </div>
              <div class="operate-button-container" v-show="orderDetail.status === 2 || orderDetail.status === 3">
                <el-button size="small" @click="showLogisticsDialog">订单跟踪</el-button>
                <el-button size="small" @click="showMessageDialog">发送站内信</el-button>
                <el-button size="small" @click="showMarkOrderDialog">备注订单</el-button>
              </div>
              <div class="operate-button-container" v-show="orderDetail.status === 4">
                <el-button size="small" @click="handleDeleteOrder">删除订单</el-button>
                <el-button size="small" @click="showMarkOrderDialog">备注订单</el-button>
              </div>
            </el-col>
          </el-row>
        </div>
        <!-- 基本信息 -->
        <div style="margin-top: 20px; display: flex; align-items: center">
          <SvgIcon name="marker" :icon-style="{ width: '20px', height: '20px', color: '#606266' }" />
          <span class="font-small">基本信息</span>
        </div>
        <div class="table-layout">
          <el-row>
            <el-col :span="4" class="table-cell-title">订单编号</el-col>
            <el-col :span="4" class="table-cell-title">发货单流水号</el-col>
            <el-col :span="4" class="table-cell-title">用户账号</el-col>
            <el-col :span="4" class="table-cell-title">支付方式</el-col>
            <el-col :span="4" class="table-cell-title">订单来源</el-col>
            <el-col :span="4" class="table-cell-title">订单类型</el-col>
          </el-row>
          <el-row>
            <el-col :span="4" class="table-cell">{{ orderDetail.orderSn }}</el-col>
            <el-col :span="4" class="table-cell">暂无</el-col>
            <el-col :span="4" class="table-cell">{{ orderDetail.memberUsername }}</el-col>
            <el-col :span="4" class="table-cell">{{ formatPayType(orderDetail.payType) }}</el-col>
            <el-col :span="4" class="table-cell">{{ formatSourceType(orderDetail.sourceType) }}</el-col>
            <el-col :span="4" class="table-cell">{{ formatOrderType(orderDetail.orderType) }}</el-col>
          </el-row>
          <el-row>
            <el-col :span="4" class="table-cell-title">配送方式</el-col>
            <el-col :span="4" class="table-cell-title">物流单号</el-col>
            <el-col :span="4" class="table-cell-title">自动确认收货时间</el-col>
            <el-col :span="4" class="table-cell-title">订单可得优币</el-col>
            <el-col :span="4" class="table-cell-title">订单可得成长值</el-col>
            <el-col :span="4" class="table-cell-title">活动信息</el-col>
          </el-row>
          <el-row>
            <el-col :span="4" class="table-cell">{{ formatNull(orderDetail.deliveryCompany) }}</el-col>
            <el-col :span="4" class="table-cell">{{ formatNull(orderDetail.deliverySn) }}</el-col>
            <el-col :span="4" class="table-cell">{{ orderDetail.autoConfirmDay }}天</el-col>
            <el-col :span="4" class="table-cell">{{ orderDetail.integration }}</el-col>
            <el-col :span="4" class="table-cell">{{ orderDetail.growth }}</el-col>
            <el-col :span="4" class="table-cell">
              <el-popover placement="top-start" title="活动信息" width="200" trigger="hover" :content="orderDetail.promotionInfo">
                <template #reference>
                  <span>{{ formatLongText(orderDetail.promotionInfo) }}</span>
                </template>
              </el-popover>
            </el-col>
          </el-row>
        </div>
        <!-- 收货人信息 -->
        <div style="margin-top: 20px; display: flex; align-items: center">
          <SvgIcon name="marker" :icon-style="{ width: '20px', height: '20px', color: '#606266' }" />
          <span class="font-small">收货人信息</span>
        </div>
        <div class="table-layout">
          <el-row>
            <el-col :span="6" class="table-cell-title">收货人</el-col>
            <el-col :span="6" class="table-cell-title">手机号码</el-col>
            <el-col :span="6" class="table-cell-title">邮政编码</el-col>
            <el-col :span="6" class="table-cell-title">收货地址</el-col>
          </el-row>
          <el-row>
            <el-col :span="6" class="table-cell">{{ orderDetail.receiverName }}</el-col>
            <el-col :span="6" class="table-cell">{{ orderDetail.receiverPhone }}</el-col>
            <el-col :span="6" class="table-cell">{{ orderDetail.receiverPostCode }}</el-col>
            <el-col :span="6" class="table-cell">{{ formatAddress(orderDetail) }}</el-col>
          </el-row>
        </div>
        <!-- 商品信息 -->
        <div style="margin-top: 20px; display: flex; align-items: center">
          <SvgIcon name="marker" :icon-style="{ width: '20px', height: '20px', color: '#606266' }" />
          <span class="font-small">商品信息</span>
        </div>
        <el-table ref="orderItemTable" :data="orderDetail.orderItems" style="width: 100%; margin-top: 20px" border>
          <el-table-column label="商品图片" width="120" align="center">
            <template #default="scope">
              <img :src="scope.row.productPic" style="height: 80px" />
            </template>
          </el-table-column>
          <el-table-column label="商品名称" align="center">
            <template #default="scope">
              <p>{{ scope.row.productName }}</p>
              <p>品牌：{{ scope.row.productBrand }}</p>
            </template>
          </el-table-column>
          <el-table-column label="价格/货号" width="120" align="center">
            <template #default="scope">
              <p>价格：￥{{ scope.row.productPrice }}</p>
              <p>货号：{{ scope.row.productSn }}</p>
            </template>
          </el-table-column>
          <el-table-column label="属性" width="120" align="center">
            <template #default="scope">
              {{ formatProductAttr(scope.row.productAttr) }}
            </template>
          </el-table-column>
          <el-table-column label="数量" width="120" align="center">
            <template #default="scope">
              {{ scope.row.productQuantity }}
            </template>
          </el-table-column>
          <el-table-column label="小计" width="120" align="center">
            <template #default="scope"> ￥{{ scope.row.productPrice * scope.row.productQuantity }} </template>
          </el-table-column>
        </el-table>
        <div style="float: right; margin: 20px">
          合计：<span class="color-danger">￥{{ orderDetail.totalAmount }}</span>
        </div>
        <!-- 费用信息 -->
        <div style="margin-top: 20px; display: flex; align-items: center">
          <SvgIcon name="marker" :icon-style="{ width: '20px', height: '20px', color: '#606266' }" />
          <span class="font-small">费用信息</span>
        </div>
        <div class="table-layout">
          <el-row>
            <el-col :span="6" class="table-cell-title">商品合计</el-col>
            <el-col :span="6" class="table-cell-title">运费</el-col>
            <el-col :span="6" class="table-cell-title">优惠券</el-col>
            <el-col :span="6" class="table-cell-title">积分抵扣</el-col>
          </el-row>
          <el-row>
            <el-col :span="6" class="table-cell">￥{{ orderDetail.totalAmount }}</el-col>
            <el-col :span="6" class="table-cell">￥{{ orderDetail.freightAmount }}</el-col>
            <el-col :span="6" class="table-cell">-￥{{ orderDetail.couponAmount }}</el-col>
            <el-col :span="6" class="table-cell">-￥{{ orderDetail.integrationAmount }}</el-col>
          </el-row>
          <el-row>
            <el-col :span="6" class="table-cell-title">活动优惠</el-col>
            <el-col :span="6" class="table-cell-title">折扣金额</el-col>
            <el-col :span="6" class="table-cell-title">订单总金额</el-col>
            <el-col :span="6" class="table-cell-title">应付款金额</el-col>
          </el-row>
          <el-row>
            <el-col :span="6" class="table-cell">-￥{{ orderDetail.promotionAmount }}</el-col>
            <el-col :span="6" class="table-cell">-￥{{ orderDetail.discountAmount }}</el-col>
            <el-col :span="6" class="table-cell">
              <span class="color-danger">￥{{ orderDetail.totalAmount + orderDetail.freightAmount }}</span>
            </el-col>
            <el-col :span="6" class="table-cell">
              <span class="color-danger">
                ￥{{ orderDetail.payAmount + orderDetail.freightAmount - orderDetail.discountAmount }}
              </span>
            </el-col>
          </el-row>
        </div>
        <!-- 操作信息 -->
        <div style="margin-top: 20px; display: flex; align-items: center">
          <SvgIcon name="marker" :icon-style="{ width: '20px', height: '20px', color: '#606266' }" />
          <span class="font-small">操作信息</span>
        </div>
        <el-table style="margin-top: 20px; width: 100%" ref="orderHistoryTable" :data="orderDetail.orderOperateHistories" border>
          <el-table-column label="操作者" width="120" align="center">
            <template #default="scope">
              {{ scope.row.operateMan }}
            </template>
          </el-table-column>
          <el-table-column label="操作时间" width="160" align="center">
            <template #default="scope">
              {{ formatTimestamp(scope.row.createTime) }}
            </template>
          </el-table-column>
          <el-table-column label="订单状态" width="120" align="center">
            <template #default="scope">
              {{ formatStatus(scope.row.orderStatus) }}
            </template>
          </el-table-column>
          <el-table-column label="付款状态" width="120" align="center">
            <template #default="scope">
              {{ formatPayStatus(scope.row.orderStatus) }}
            </template>
          </el-table-column>
          <el-table-column label="发货状态" width="120" align="center">
            <template #default="scope">
              {{ formatDeliverStatus(scope.row.orderStatus) }}
            </template>
          </el-table-column>
          <el-table-column label="备注" align="center">
            <template #default="scope">
              {{ scope.row.note }}
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </div>
  </el-drawer>
</template>

<script setup lang="ts" name="OrderDrawer">
import SvgIcon from "@/components/SvgIcon/index.vue";
import { ElIcon } from "element-plus";
import { WarningFilled } from "@element-plus/icons-vue";
import { ref } from "vue";
import { Order } from "@/api/interface";
import { getOrderSyncApi } from "@/api/modules/order";
import { formatTimestamp } from "@/utils/time";
import {
  formatAddress,
  formatDeliverStatus,
  formatLongText,
  formatNull,
  formatOrderType,
  formatPayStatus,
  formatPayType,
  formatProductAttr,
  formatSourceType,
  formatStatus
} from "./order";

// 格式化状态
const formatStepStatus = (value: number) => {
  if (value === 1) {
    // 待发货
    return 2;
  } else if (value === 2) {
    // 已发货
    return 3;
  } else if (value === 3) {
    // 已完成
    return 4;
  } else {
    // 待付款、已关闭、无限订单
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
const acceptParams = async (params: DrawerProps) => {
  drawerProps.value = params;
  drawerVisible.value = true;
  if (params.row.id && params.row.id !== 0) {
    orderDetail.value = await getOrderSyncApi(params.row.id);
  }
};

defineExpose({
  acceptParams
});
</script>

<style scoped lang="scss">
.detail-container {
  width: 100%;
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
