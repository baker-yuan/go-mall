<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="60%" :title="`订单退货申请`">
    <div v-if="returnApplyDetail != null" class="detail-container">
      <!-- 退货商品 -->
      <el-card shadow="never">
        <span class="font-title-medium">退货商品</span>
        <el-table border class="standard-margin" ref="productTable" :data="returnApplies">
          <el-table-column label="商品图片" width="160" align="center">
            <template #default="scope">
              <img style="height: 80px" :src="scope.row.productPic" />
            </template>
          </el-table-column>
          <el-table-column label="商品名称" align="center">
            <template #default="scope">
              <span class="font-small">{{ scope.row.productName }}</span>
              <br />
              <span class="font-small">品牌：{{ scope.row.productBrand }}</span>
            </template>
          </el-table-column>
          <el-table-column label="价格/货号" width="180" align="center">
            <template #default="scope">
              <span class="font-small">价格：￥{{ scope.row.productRealPrice }}</span>
              <br />
              <span class="font-small">货号：NO.{{ scope.row.productId }}</span>
            </template>
          </el-table-column>
          <el-table-column label="属性" width="180" align="center">
            <template #default="scope">{{ scope.row.productAttr }}</template>
          </el-table-column>
          <el-table-column label="数量" width="100" align="center">
            <template #default="scope">{{ scope.row.productCount }}</template>
          </el-table-column>
          <el-table-column label="小计" width="100" align="center">
            <template #default>￥{{ totalAmount }}</template>
          </el-table-column>
        </el-table>
        <div style="float: right; margin-top: 15px; margin-bottom: 15px">
          <span class="font-title-medium">合计：</span>
          <span class="font-title-medium color-danger">￥{{ totalAmount }}</span>
        </div>
      </el-card>
      <!-- 服务单信息 -->
      <el-card shadow="never" class="standard-margin">
        <span class="font-title-medium">服务单信息</span>
        <div class="form-container-border">
          <el-row>
            <el-col :span="6" class="form-border form-left-bg font-small">服务单号</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.id }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">申请状态</el-col>
            <el-col class="form-border font-small" :span="18">{{ formatStatus(returnApplyDetail.status) }}</el-col>
          </el-row>
          <el-row>
            <el-col :span="6" class="form-border form-left-bg font-small" style="height: 50px; line-height: 30px">
              订单编号
            </el-col>
            <el-col class="form-border font-small" :span="18" style="height: 50px">
              {{ returnApplyDetail.orderSn }}
              <el-button type="text" size="small" @click="handleViewOrder">查看</el-button>
            </el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">申请时间</el-col>
            <el-col class="form-border font-small" :span="18">{{ formatTimestamp(returnApplyDetail.createTime) }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">用户账号</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.memberUsername }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">联系人</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.returnName }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">联系电话</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.returnPhone }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">退货原因</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.reason }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">问题描述</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.description }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6" style="height: 100px; line-height: 80px">
              凭证图片
            </el-col>
            <el-col class="form-border font-small" :span="18" style="height: 100px">
              <img v-for="item in returnApplyDetail.proofPics" :key="item" style="width: 80px; height: 80px" :src="item" alt="" />
            </el-col>
          </el-row>
        </div>
        <!-- -->
        <div class="form-container-border">
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">订单金额</el-col>
            <el-col class="form-border font-small" :span="18">￥{{ totalAmount }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6" style="height: 52px; line-height: 32px">
              确认退款金额
            </el-col>
            <el-col class="form-border font-small" style="height: 52px" :span="18">
              ￥
              <el-input
                size="small"
                v-model="updateStatusParam.returnAmount"
                :disabled="returnApplyDetail.status !== 0"
                style="width: 200px; margin-left: 10px"
              ></el-input>
            </el-col>
          </el-row>
          <div v-show="returnApplyDetail.status !== 3">
            <el-row>
              <el-col class="form-border form-left-bg font-small" :span="6" style="height: 52px; line-height: 32px">
                选择收货点
              </el-col>
              <el-col class="form-border font-small" style="height: 52px" :span="18">
                <el-select
                  size="small"
                  style="width: 200px"
                  :disabled="returnApplyDetail.status !== 0"
                  v-model="updateStatusParam.companyAddressId"
                >
                  <el-option v-for="item in companyAddressOptions" :key="item.value" :label="item.label" :value="item.value">
                  </el-option>
                </el-select>
              </el-col>
            </el-row>
            <el-row>
              <el-col class="form-border form-left-bg font-small" :span="6">收货人姓名</el-col>
              <el-col class="form-border font-small" :span="18">{{ companyAddress.name }}</el-col>
            </el-row>
            <el-row>
              <el-col class="form-border form-left-bg font-small" :span="6">所在区域</el-col>
              <el-col class="form-border font-small" :span="18">{{ formatRegion(companyAddress) }}</el-col>
            </el-row>
            <el-row>
              <el-col class="form-border form-left-bg font-small" :span="6">详细地址</el-col>
              <el-col class="form-border font-small" :span="18">{{ companyAddress.detailAddress }}</el-col>
            </el-row>
            <el-row>
              <el-col class="form-border form-left-bg font-small" :span="6">联系电话</el-col>
              <el-col class="form-border font-small" :span="18">{{ companyAddress.phone }}</el-col>
            </el-row>
          </div>
        </div>
        <!-- 处理 -->
        <div class="form-container-border" v-show="returnApplyDetail.status !== 0">
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">处理人员</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.handleMan }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">处理时间</el-col>
            <el-col class="form-border font-small" :span="18">{{ formatTimestamp(returnApplyDetail.handleTime) }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">处理备注</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.handleNote }}</el-col>
          </el-row>
        </div>
        <!-- 收货(退货中) -->
        <div class="form-container-border" v-show="returnApplyDetail.status === 2">
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">收货人员</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.receiveMan }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">收货时间</el-col>
            <el-col class="form-border font-small" :span="18">{{ formatTimestamp(returnApplyDetail.receiveTime) }}</el-col>
          </el-row>
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6">收货备注</el-col>
            <el-col class="form-border font-small" :span="18">{{ returnApplyDetail.receiveNote }}</el-col>
          </el-row>
        </div>
        <!-- 处理备注(待处理) -->
        <div class="form-container-border" v-show="returnApplyDetail.status === 0">
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6" style="height: 52px; line-height: 32px">
              处理备注
            </el-col>
            <el-col class="form-border font-small" :span="18">
              <el-input size="small" v-model="updateStatusParam.handleNote" style="width: 200px; margin-left: 10px"></el-input>
            </el-col>
          </el-row>
        </div>
        <!-- 收货备注(退货中) -->
        <div class="form-container-border" v-show="returnApplyDetail.status === 1">
          <el-row>
            <el-col class="form-border form-left-bg font-small" :span="6" style="height: 52px; line-height: 32px">
              收货备注
            </el-col>
            <el-col class="form-border font-small" :span="18">
              <el-input size="small" v-model="updateStatusParam.receiveNote" style="width: 200px; margin-left: 10px"></el-input>
            </el-col>
          </el-row>
        </div>
        <!-- 待处理 -->
        <div style="margin-top: 15px; text-align: center" v-show="returnApplyDetail.status === 0">
          <el-button type="primary" size="small" @click="handleUpdateStatus(1)">确认退货</el-button>
          <el-button type="danger" size="small" @click="handleUpdateStatus(3)">拒绝退货</el-button>
        </div>
        <!-- 退货中 -->
        <div style="margin-top: 15px; text-align: center" v-show="returnApplyDetail.status === 1">
          <el-button type="primary" size="small" @click="handleUpdateStatus(2)">确认收货</el-button>
        </div>
      </el-card>
    </div>
  </el-drawer>
</template>

<script setup lang="ts" name="OrderReturnApplyDrawer">
import { computed, ref, watch } from "vue";
import { CompanyAddress, OptionValue, OrderReturnApply } from "@/api/interface";
import { getOrderReturnAppliesSyncApi } from "@/api/modules/orderReturnApply";
import { formatRegion, formatStatus } from "@/views/oms/orderReturnApply/orderReturnApply";
import { formatTimestamp } from "../../../utils/time";
import { getCompanyAddressesSyncApi, getCompanyAddressOptionsApi } from "@/api/modules/companyAddress";

// 订单退货申请详情
const returnApplyDetail = ref<OrderReturnApply.OrderReturnApplyModel | null>(null);
// 订单退货申请(为了渲染tab搞了一个数组) productList = [returnApplyDetail]
const returnApplies = ref<OrderReturnApply.OrderReturnApplyModel[]>([]);
// 公司收发货地址
const companyAddress = ref<CompanyAddress.CompanyAddressModel>({
  id: 0,
  addressName: "",
  sendStatus: 0,
  receiveStatus: 0,
  name: "",
  phone: "",
  province: "",
  city: "",
  region: "",
  detailAddress: ""
});
// 公司
const companyAddressOptions = ref<OptionValue[]>([]);

interface UpdateStatusParamReq {
  receiveNote: string; // 收货备注
  companyAddressId: number; // 收货地址表id
  handleNote: string; // 处理备注
  returnAmount: number; // 退款金额
}
// updateStatusParam
const updateStatusParam = ref<UpdateStatusParamReq>({
  receiveNote: "",
  companyAddressId: 0,
  handleNote: "",
  returnAmount: 0
});

let companyAddressId = computed(() => updateStatusParam.value.companyAddressId);
watch(companyAddressId, newValue => {
  if (newValue === 0) {
    return;
  }
  const fetchCompanyAddress = async () => {
    companyAddress.value = await getCompanyAddressesSyncApi(newValue);
  };
  fetchCompanyAddress();
});

// 总价
const totalAmount = computed(() => {
  if (returnApplyDetail.value != null) {
    return returnApplyDetail.value.productRealPrice * returnApplyDetail.value.productCount;
  } else {
    return 0;
  }
});

interface DrawerProps {
  row: Partial<OrderReturnApply.OrderReturnApplyModel>; // 数据
}

// 是否显示弹框
const drawerVisible = ref(false);
// 参数
const drawerProps = ref<DrawerProps>({
  row: {}
});

// 接收父组件传过来的参数
const acceptParams = async (params: DrawerProps) => {
  drawerProps.value = params;
  drawerVisible.value = true;

  if (params.row.id && params.row.id !== 0) {
    returnApplyDetail.value = await getOrderReturnAppliesSyncApi(params.row.id);
    returnApplies.value.push(returnApplyDetail.value);
    if (returnApplyDetail.value.companyAddress) {
      companyAddress.value = returnApplyDetail.value.companyAddress;
    }
    companyAddressOptions.value = await getCompanyAddressOptionsApi();
  }
};

defineExpose({
  acceptParams
});
</script>
<style scoped lang="scss">
.detail-container {
  //position: absolute;
  //left: 0;
  //right: 0;
  //width: 1080px;
  //padding: 35px 35px 15px 35px;
  //margin: 20px auto;
}

.standard-margin {
  margin-top: 15px;
}
.form-border {
  border-right: 1px solid #dcdfe6;
  border-bottom: 1px solid #dcdfe6;
  padding: 10px;
}

.form-container-border {
  border-left: 1px solid #dcdfe6;
  border-top: 1px solid #dcdfe6;
  margin-top: 15px;
}

.form-left-bg {
  background: #f2f6fc;
}
</style>
