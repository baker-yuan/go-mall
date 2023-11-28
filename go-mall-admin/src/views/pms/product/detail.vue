<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="900px" :title="`${drawerProps.title}商品表`">
    <el-steps :active="activeStep" finish-status="success" align-center>
      <el-step title="填写商品信息"></el-step>
      <el-step title="填写商品促销"></el-step>
      <el-step title="填写商品属性"></el-step>
      <el-step title="选择商品关联"></el-step>
    </el-steps>
    <div style="margin-top: 50px">
      <el-form :model="drawerProps.row" :rules="rules" ref="ruleFormRef" label-width="120px" class="form-inner-container">
        <div v-show="activeStep === 0">
          <el-form-item label="商品分类" prop="productCategoryId">
            <el-cascader v-model="drawerProps.row!.productCategoryId" :options="categoryTreeData"> </el-cascader>
          </el-form-item>
          <el-form-item label="商品名称" prop="name">
            <el-input v-model="drawerProps.row!.name"></el-input>
          </el-form-item>
          <el-form-item label="副标题" prop="subTitle">
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
            <el-input
              :auto-size="true"
              v-model="drawerProps.row!.description"
              type="textarea"
              placeholder="请输入内容"
            ></el-input>
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
          <!--
          <el-form-item style="text-align: center">
            <el-button type="primary" @click="handleNext('productInfoForm')">下一步，填写商品促销</el-button>
          </el-form-item>
          -->
        </div>
        <div v-show="activeStep === 1">
          <el-form-item label="赠送积分：">
            <el-input v-model="drawerProps.row!.giftPoint"></el-input>
          </el-form-item>
          <el-form-item label="赠送成长值：">
            <el-input v-model="drawerProps.row!.giftGrowth"></el-input>
          </el-form-item>
          <el-form-item label="积分购买限制：">
            <el-input v-model="drawerProps.row!.usePointLimit"></el-input>
          </el-form-item>
          <el-form-item label="预告商品：">
            <el-switch v-model="drawerProps.row!.previewStatus" :active-value="1" :inactive-value="0"></el-switch>
          </el-form-item>
          <el-form-item label="商品上架：">
            <el-switch v-model="drawerProps.row!.publishStatus" :active-value="1" :inactive-value="0"></el-switch>
          </el-form-item>
          <el-form-item label="商品推荐：">
            <span style="margin-right: 10px">新品</span>
            <el-switch v-model="drawerProps.row!.newStatus" :active-value="1" :inactive-value="0"></el-switch>
            <span style="margin-left: 10px; margin-right: 10px">推荐</span>
            <el-switch v-model="drawerProps.row!.recommandStatus" :active-value="1" :inactive-value="0"></el-switch>
          </el-form-item>
          <!--
          <el-form-item label="服务保证：">
            <el-checkbox-group v-model="selectServiceList">
              <el-checkbox :label="1">无忧退货</el-checkbox>
              <el-checkbox :label="2">快速退款</el-checkbox>
              <el-checkbox :label="3">免费包邮</el-checkbox>
            </el-checkbox-group>
          </el-form-item>
          -->
          <el-form-item label="详细页标题：">
            <el-input v-model="drawerProps.row!.detailTitle"></el-input>
          </el-form-item>
          <el-form-item label="详细页描述：">
            <el-input v-model="drawerProps.row!.detailDesc"></el-input>
          </el-form-item>
          <el-form-item label="商品关键字：">
            <el-input v-model="drawerProps.row!.keywords"></el-input>
          </el-form-item>
          <el-form-item label="商品备注：">
            <el-input v-model="drawerProps.row!.note" type="textarea" :auto-size="true"></el-input>
          </el-form-item>
          <el-form-item label="选择优惠方式：">
            <el-radio-group v-model="drawerProps.row!.promotionType" size="small">
              <el-radio-button :label="0">无优惠</el-radio-button>
              <el-radio-button :label="1">特惠促销</el-radio-button>
              <el-radio-button :label="2">会员价格</el-radio-button>
              <el-radio-button :label="3">阶梯价格</el-radio-button>
              <el-radio-button :label="4">满减价格</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <!--
          <el-form-item v-show="drawerProps.row!.promotionType===1">
            <div>
              开始时间：
              <el-date-picker
                v-model="drawerProps.row!.promotionStartTime"
                type="datetime"
                :picker-options="pickerOptions1"
                placeholder="选择开始时间">
              </el-date-picker>
            </div>
            <div class="littleMargin">
              结束时间：
              <el-date-picker
                v-model="drawerProps.row!.promotionEndTime"
                type="datetime"
                :picker-options="pickerOptions1"
                placeholder="选择结束时间">
              </el-date-picker>
            </div>
            <div class="littleMargin">
              促销价格：
              <el-input style="width: 220px" v-model="drawerProps.row!.promotionPrice" placeholder="输入促销价格"></el-input>
            </div>
          </el-form-item>

          <el-form-item v-show="drawerProps.row!.promotionType===2">
            <div v-for="(item, index) in drawerProps.row!.memberPriceList" :class="{littleMargin:index!==0}">
              {{item.memberLevelName}}：
              <el-input v-model="item.memberPrice" style="width: 200px"></el-input>
            </div>
          </el-form-item>
          <el-form-item v-show="drawerProps.row!.promotionType===3">
            <el-table :data="drawerProps.row!.productLadderList"
                      style="width: 80%" border>
              <el-table-column
                label="数量"
                align="center"
                width="120">
                <template slot-scope="scope">
                  <el-input v-model="scope.row.count"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                label="折扣"
                align="center"
                width="120">
                <template slot-scope="scope">
                  <el-input v-model="scope.row.discount"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                align="center"
                label="操作">
                <template slot-scope="scope">
                  <el-button type="text" @click="handleRemoveProductLadder(scope.$index, scope.row)">删除</el-button>
                  <el-button type="text" @click="handleAddProductLadder(scope.$index, scope.row)">添加</el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-form-item>
          <el-form-item v-show="value.promotionType===4">
            <el-table :data="value.productFullReductionList"
                      style="width: 80%" border>
              <el-table-column
                label="满"
                align="center"
                width="120">
                <template slot-scope="scope">
                  <el-input v-model="scope.row.fullPrice"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                label="立减"
                align="center"
                width="120">
                <template slot-scope="scope">
                  <el-input v-model="scope.row.reducePrice"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                align="center"
                label="操作">
                <template slot-scope="scope">
                  <el-button type="text" @click="handleRemoveFullReduction(scope.$index, scope.row)">删除</el-button>
                  <el-button type="text" @click="handleAddFullReduction(scope.$index, scope.row)">添加</el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-form-item>
          -->
        </div>
        <div v-show="activeStep === 2">
          333
          <!--
          <el-form-item label="属性类型：">
            <el-select v-model="drawerProps.row!.productAttributeCategoryId"
                       placeholder="请选择属性类型"
                       @change="handleProductAttrChange">
              <el-option
                v-for="item in productAttributeCategoryOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="商品规格：">
            <el-card shadow="never" class="cardBg">
              <div v-for="(productAttr,idx) in selectProductAttr">
                {{productAttr.name}}：
                <el-checkbox-group v-if="productAttr.handAddStatus===0" v-model="selectProductAttr[idx].values">
                  <el-checkbox v-for="item in getInputListArr(productAttr.inputList)" :label="item" :key="item"
                               class="littleMarginLeft"></el-checkbox>
                </el-checkbox-group>
                <div v-else>
                  <el-checkbox-group v-model="selectProductAttr[idx].values">
                    <div v-for="(item,index) in selectProductAttr[idx].options" style="display: inline-block"
                         class="littleMarginLeft">
                      <el-checkbox :label="item" :key="item"></el-checkbox>
                      <el-button type="text" class="littleMarginLeft" @click="handleRemoveProductAttrValue(idx,index)">删除
                      </el-button>
                    </div>
                  </el-checkbox-group>
                  <el-input v-model="addProductAttrValue" style="width: 160px;margin-left: 10px" clearable></el-input>
                  <el-button class="littleMarginLeft" @click="handleAddProductAttrValue(idx)">增加</el-button>
                </div>
              </div>
            </el-card>
            <el-table style="width: 100%;margin-top: 20px"
                      :data="drawerProps.row!.skuStockList"
                      border>
              <el-table-column
                v-for="(item,index) in selectProductAttr"
                :label="item.name"
                :key="item.id"
                align="center">
                <template slot-scope="scope">
                  {{getProductSkuSp(scope.row,index)}}
                </template>
              </el-table-column>
              <el-table-column
                label="销售价格"
                width="100"
                align="center">
                <template slot-scope="scope">
                  <el-input v-model="drawerProps.row!.row.price"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                label="促销价格"
                width="100"
                align="center">
                <template slot-scope="scope">
                  <el-input v-model="drawerProps.row!.row.promotionPrice"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                label="商品库存"
                width="80"
                align="center">
                <template slot-scope="scope">
                  <el-input v-model="drawerProps.row!.row.stock"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                label="库存预警值"
                width="80"
                align="center">
                <template slot-scope="scope">
                  <el-input v-model="drawerProps.row!.row.lowStock"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                label="SKU编号"
                width="160"
                align="center">
                <template slot-scope="scope">
                  <el-input v-model="drawerProps.row!.row.skuCode"></el-input>
                </template>
              </el-table-column>
              <el-table-column
                label="操作"
                width="80"
                align="center">
                <template slot-scope="scope">
                  <el-button
                    type="text"
                    @click="handleRemoveProductSku(scope.$index, scope.row)">删除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-button
              type="primary"
              style="margin-top: 20px"
              @click="handleRefreshProductSkuList">刷新列表
            </el-button>
            <el-button
              type="primary"
              style="margin-top: 20px"
              @click="handleSyncProductSkuPrice">同步价格
            </el-button>
            <el-button
              type="primary"
              style="margin-top: 20px"
              @click="handleSyncProductSkuStock">同步库存
            </el-button>
          </el-form-item>
          <el-form-item label="属性图片：" v-if="hasAttrPic">
            <el-card shadow="never" class="cardBg">
              <div v-for="(item,index) in selectProductAttrPics">
                <span>{{item.name}}:</span>
                <single-upload v-model="item.pic"
                               style="width: 300px;display: inline-block;margin-left: 10px"></single-upload>
              </div>
            </el-card>
          </el-form-item>
          <el-form-item label="商品参数：">
            <el-card shadow="never" class="cardBg">
              <div v-for="(item,index) in selectProductParam" :class="{littleMarginTop:index!==0}">
                <div class="paramInputLabel">{{item.name}}:</div>
                <el-select v-if="item.inputType===1" class="paramInput" v-model="selectProductParam[index].value">
                  <el-option
                    v-for="item in getParamInputList(item.inputList)"
                    :key="item"
                    :label="item"
                    :value="item">
                  </el-option>
                </el-select>
                <el-input v-else class="paramInput" v-model="selectProductParam[index].value"></el-input>
              </div>
            </el-card>
          </el-form-item>
          <el-form-item label="商品相册：">
            <multi-upload v-model="selectProductPics"></multi-upload>
          </el-form-item>
          <el-form-item label="商品详情：">
            <el-tabs v-model="activeHtmlName" type="card">
              <el-tab-pane label="电脑端详情" name="pc">
                <tinymce :width="595" :height="300" v-model="value.detailHtml"></tinymce>
              </el-tab-pane>
              <el-tab-pane label="移动端详情" name="mobile">
                <tinymce :width="595" :height="300" v-model="value.detailMobileHtml"></tinymce>
              </el-tab-pane>
            </el-tabs>
          </el-form-item>
          -->
        </div>
        <div v-show="activeStep === 3">
          444
          <!--
          <el-form-item label="关联专题：">
            <el-transfer
              style="display: inline-block"
              filterable
              :filter-method="filterMethod"
              filter-placeholder="请输入专题名称"
              v-model="selectSubject"
              :titles="subjectTitles"
              :data="subjectList">
            </el-transfer>
          </el-form-item>
          <el-form-item label="关联优选：">
            <el-transfer
              style="display: inline-block"
              filterable
              :filter-method="filterMethod"
              filter-placeholder="请输入优选名称"
              v-model="selectPrefrenceArea"
              :titles="prefrenceAreaTitles"
              :data="prefrenceAreaList">
            </el-transfer>
          </el-form-item>
          -->
        </div>
      </el-form>
    </div>
    <template #footer>
      <el-button v-show="activeStep != 0" type="primary" @click="handlePrev('productInfoForm')">上一步</el-button>
      <el-button v-show="activeStep != 3" type="primary" @click="handleNext('productInfoForm')">下一步</el-button>
      <el-button @click="drawerVisible = false">取消</el-button>
      <el-button v-show="activeStep == 3 && !drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="ProductDrawer">
import { onBeforeMount, reactive, ref } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { CascaderValue, Product } from "@/api/interface";
import { categoryTreeApi } from "@/api/modules/category";

const rules = reactive({
  name: [
    { required: true, message: "请输入商品名称", trigger: "blur" },
    { min: 2, max: 140, message: "长度在 2 到 140 个字符", trigger: "blur" }
  ],
  subTitle: [{ required: true, message: "请输入商品副标题", trigger: "blur" }],
  productCategoryId: [{ required: true, message: "请选择商品分类", trigger: "blur" }],
  brandId: [{ required: true, message: "请选择商品品牌", trigger: "blur" }],
  description: [{ required: true, message: "请输入商品介绍", trigger: "blur" }],
  requiredProp: [{ required: true, message: "该项为必填项", trigger: "blur" }]
});

// 当前tab
const activeStep = ref<number>(0);
// 表单引用
const ruleFormRef = ref<FormInstance>();
// 切换到下一个tab
const handleNext = (formName: string) => {
  activeStep.value = activeStep.value + 1;
  if (formName === "productInfoForm") {
    ruleFormRef.value!.validate(async valid => {
      if (!valid) return;
    });
  }
};
// 切换到上一个tab
const handlePrev = (formName: string) => {
  console.log(formName);
  activeStep.value = activeStep.value - 1;
};

const categoryTreeData = ref<CascaderValue[]>([]);

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
  activeStep.value = 0;
};

// 提交数据（新增/编辑）
const handleSubmit = async () => {
  try {
    await drawerProps.value.api!(drawerProps.value.row);
    ElMessage.success({ message: `${drawerProps.value.title}商品表成功！` });
    drawerProps.value.getTableList!();
    drawerVisible.value = false;
  } catch (error) {
    console.log(error);
  }
};

onBeforeMount(async () => {
  categoryTreeData.value = await categoryTreeApi();
});

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
