<template>
  <el-drawer
    v-model="drawerVisible"
    :destroy-on-close="true"
    :close-on-click-modal="false"
    :show-close="false"
    size="900px"
    :title="`${drawerProps.title}商品表`"
  >
    <el-steps :active="activeStep" finish-status="success" align-center>
      <el-step title="填写商品信息"></el-step>
      <el-step title="填写商品促销"></el-step>
      <el-step title="填写商品属性"></el-step>
      <el-step title="选择商品关联"></el-step>
    </el-steps>
    <div style="margin-top: 50px">
      <el-form
        ref="ruleFormRef"
        :model="productDetail"
        :rules="rules"
        label-suffix=" :"
        label-width="120px"
        class="form-inner-container"
      >
        <!-- 填写商品信息 -->
        <div v-show="activeStep === 0">
          <el-form-item label="商品分类" prop="productCategoryId">
            <el-cascader v-model="productDetail!.productCategoryId" :options="categoryOptions"> </el-cascader>
          </el-form-item>
          <el-form-item label="商品名称" prop="name">
            <el-input v-model="productDetail!.name"></el-input>
          </el-form-item>
          <el-form-item label="副标题" prop="subTitle">
            <el-input v-model="productDetail!.subTitle"></el-input>
          </el-form-item>
          <el-form-item label="商品品牌" prop="brandId">
            <el-select v-model="productDetail!.brandId" placeholder="请选择品牌">
              <el-option v-for="item in brandOptions" :key="item.value" :label="item.label" :value="item.value"> </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="商品介绍">
            <el-input :auto-size="true" v-model="productDetail!.description" type="textarea" placeholder="请输入内容"></el-input>
          </el-form-item>
          <el-form-item label="商品货号">
            <el-input v-model="productDetail!.productSn"></el-input>
          </el-form-item>
          <el-form-item label="商品售价">
            <el-input v-model="productDetail!.price"></el-input>
          </el-form-item>
          <el-form-item label="市场价">
            <el-input v-model="productDetail!.originalPrice"></el-input>
          </el-form-item>
          <el-form-item label="商品库存">
            <el-input v-model="productDetail!.stock"></el-input>
          </el-form-item>
          <el-form-item label="计量单位">
            <el-input v-model="productDetail!.unit"></el-input>
          </el-form-item>
          <el-form-item label="商品重量">
            <el-input v-model="productDetail!.weight" style="width: 300px"></el-input>
            <span style="margin-left: 20px">克</span>
          </el-form-item>
          <el-form-item label="排序">
            <el-input v-model="productDetail!.sort"></el-input>
          </el-form-item>
        </div>

        <!-- 填写商品促销 -->
        <div v-show="activeStep === 1">
          <el-form-item label="赠送积分">
            <el-input v-model="productDetail!.giftPoint"></el-input>
          </el-form-item>
          <el-form-item label="赠送成长值">
            <el-input v-model="productDetail!.giftGrowth"></el-input>
          </el-form-item>
          <el-form-item label="积分购买限制">
            <el-input v-model="productDetail!.usePointLimit"></el-input>
          </el-form-item>
          <el-form-item label="预告商品">
            <el-switch v-model="productDetail!.previewStatus" :active-value="1" :inactive-value="0"></el-switch>
          </el-form-item>
          <el-form-item label="商品上架">
            <el-switch v-model="productDetail!.publishStatus" :active-value="1" :inactive-value="0"></el-switch>
          </el-form-item>
          <el-form-item label="商品推荐">
            <span style="margin-right: 10px">新品</span>
            <el-switch v-model="productDetail!.newStatus" :active-value="1" :inactive-value="0"></el-switch>
            <span style="margin-left: 10px; margin-right: 10px">推荐</span>
            <el-switch v-model="productDetail!.recommandStatus" :active-value="1" :inactive-value="0"></el-switch>
          </el-form-item>
          <el-form-item label="服务保证">
            <el-checkbox-group v-model="selectServiceList">
              <el-checkbox :label="1">无忧退货</el-checkbox>
              <el-checkbox :label="2">快速退款</el-checkbox>
              <el-checkbox :label="3">免费包邮</el-checkbox>
            </el-checkbox-group>
          </el-form-item>
          <el-form-item label="详细页标题">
            <el-input v-model="productDetail!.detailTitle"></el-input>
          </el-form-item>
          <el-form-item label="详细页描述">
            <el-input v-model="productDetail!.detailDesc"></el-input>
          </el-form-item>
          <el-form-item label="商品关键字">
            <el-input v-model="productDetail!.keywords"></el-input>
          </el-form-item>
          <el-form-item label="商品备注">
            <el-input v-model="productDetail!.note" type="textarea" :auto-size="true"></el-input>
          </el-form-item>
          <el-form-item label="选择优惠方式">
            <el-radio-group v-model="productDetail!.promotionType" size="small">
              <el-radio-button :label="0">无优惠</el-radio-button>
              <el-radio-button :label="1">特惠促销</el-radio-button>
              <el-radio-button :label="2">会员价格</el-radio-button>
              <el-radio-button :label="3">阶梯价格</el-radio-button>
              <el-radio-button :label="4">满减价格</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <!-- 特惠促销 -->
          <el-form-item v-show="productDetail!.promotionType === 1">
            <div>
              开始时间：
              <el-date-picker
                v-model="promotionStartTime"
                type="datetime"
                :picker-options="pickerOptions"
                placeholder="选择开始时间"
              >
              </el-date-picker>
            </div>
            <div class="littleMargin">
              结束时间：
              <el-date-picker
                v-model="promotionEndTime"
                :picker-options="pickerOptions"
                type="datetime"
                placeholder="选择结束时间"
              >
              </el-date-picker>
            </div>
            <div class="littleMargin">
              促销价格：
              <el-input style="width: 220px" v-model="productDetail!.promotionPrice" placeholder="输入促销价格"></el-input>
            </div>
          </el-form-item>
          <!-- 会员价格 -->
          <el-form-item v-show="productDetail!.promotionType === 2">
            <div v-for="(item, index) in productDetail!.memberPrices" :key="index" :class="{ littleMargin: index !== 0 }">
              {{ item.memberLevelName }}：
              <el-input v-model="item.memberPrice" style="width: 200px"></el-input>
            </div>
          </el-form-item>
          <!-- 使用阶梯价格 -->
          <el-form-item v-show="productDetail!.promotionType === 3">
            <el-table :data="productDetail!.productLadders" style="width: 80%" border>
              <el-table-column label="数量" align="center" width="120">
                <template #default="scope">
                  <el-input v-model="scope.row.count"></el-input>
                </template>
              </el-table-column>
              <el-table-column label="折扣" align="center" width="120">
                <template #default="scope">
                  <el-input v-model="scope.row.discount"></el-input>
                </template>
              </el-table-column>
              <el-table-column align="center" label="操作">
                <template #default="scope">
                  <el-button type="text" @click="handleAddProductLadder()">添加</el-button>
                  <el-button type="text" @click="handleRemoveProductLadder(scope.$index)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-form-item>
          <!-- 使用满减价格 -->
          <el-form-item v-show="productDetail!.promotionType === 4">
            <el-table :data="productDetail!.productFullReductions" style="width: 80%" border>
              <el-table-column label="满" align="center" width="120">
                <template #default="scope">
                  <el-input v-model="scope.row.fullPrice"></el-input>
                </template>
              </el-table-column>
              <el-table-column label="立减" align="center" width="120">
                <template #default="scope">
                  <el-input v-model="scope.row.reducePrice"></el-input>
                </template>
              </el-table-column>
              <el-table-column align="center" label="操作">
                <template #default="scope">
                  <el-button type="text" @click="handleAddFullReduction()">添加</el-button>
                  <el-button type="text" @click="handleRemoveFullReduction(scope.$index)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-form-item>
        </div>

        <!-- 填写商品属性 -->
        <div v-show="activeStep === 2">
          <el-form-item label="属性类型">
            <el-select
              v-model="productDetail!.productAttributeCategoryId"
              placeholder="请选择属性类型"
              @change="handleProductAttrChange"
            >
              <el-option
                v-for="item in productAttributeCategoryOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="商品规格">
            <el-card shadow="never" class="cardBg">
              <div v-for="(productAttr, idx) in selectProductAttr" :key="idx">
                {{ productAttr.name }}：
                <el-checkbox-group v-if="productAttr.handAddStatus === 0" v-model="selectProductAttr[idx].values">
                  <el-checkbox
                    v-for="item in getInputListArr(productAttr.inputList)"
                    :label="item"
                    :key="item"
                    class="littleMarginLeft"
                  ></el-checkbox>
                </el-checkbox-group>
                <div v-else>
                  <el-checkbox-group v-model="selectProductAttr[idx].values">
                    <div
                      v-for="(item, index) in selectProductAttr[idx].options"
                      :key="index"
                      style="display: inline-block"
                      class="littleMarginLeft"
                    >
                      <el-checkbox :label="item" :key="item"></el-checkbox>
                      <el-button type="text" class="littleMarginLeft" @click="handleRemoveProductAttrValue(idx, index)">
                        删除
                      </el-button>
                    </div>
                  </el-checkbox-group>
                  <el-input v-model="addProductAttrValue" style="width: 160px; margin-left: 10px" clearable></el-input>
                  <el-button class="littleMarginLeft" @click="handleAddProductAttrValue(idx)">增加</el-button>
                </div>
              </div>
            </el-card>
            <el-table style="width: 100%; margin-top: 20px" :data="productDetail!.skuStocks" border>
              <el-table-column v-for="(item, index) in selectProductAttr" :label="item.name" :key="item.id" align="center">
                <template #default="scope">
                  {{ getProductSkuSp(scope.row, index) }}
                </template>
              </el-table-column>
              <el-table-column label="销售价格" width="100" align="center">
                <template #default="scope">
                  <el-input v-model="scope.row.price"></el-input>
                </template>
              </el-table-column>
              <el-table-column label="促销价格" width="100" align="center">
                <template #default="scope">
                  <el-input v-model="scope.row.promotionPrice"></el-input>
                </template>
              </el-table-column>
              <el-table-column label="商品库存" width="80" align="center">
                <template #default="scope">
                  <el-input v-model="scope.row.stock"></el-input>
                </template>
              </el-table-column>
              <el-table-column label="库存预警值" width="80" align="center">
                <template #default="scope">
                  <el-input v-model="scope.row.lowStock"></el-input>
                </template>
              </el-table-column>
              <el-table-column label="SKU编号" width="160" align="center">
                <template #default="scope">
                  <el-input v-model="scope.row.skuCode"></el-input>
                </template>
              </el-table-column>
              <el-table-column label="操作" width="80" align="center">
                <template #default="scope">
                  <el-button type="text" @click="handleRemoveProductSku(scope.$index, scope.row)">删除 </el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-button type="primary" style="margin-top: 20px" @click="handleRefreshProductSkuList">刷新列表 </el-button>
            <el-button type="primary" style="margin-top: 20px" @click="handleSyncProductSkuPrice">同步价格 </el-button>
            <el-button type="primary" style="margin-top: 20px" @click="handleSyncProductSkuStock">同步库存 </el-button>
          </el-form-item>
          <el-form-item label="属性图片：" v-if="hasAttrPic">
            <el-card shadow="never" class="cardBg">
              <div v-for="(item, index) in selectProductAttrPics" :key="index">
                <span>{{ item.name }}:</span>
                <single-upload v-model="item.pic" style="width: 300px; display: inline-block; margin-left: 10px"></single-upload>
              </div>
            </el-card>
          </el-form-item>

          <el-form-item label="商品参数">
            <el-card shadow="never" class="cardBg">
              <div v-for="(item, index) in selectProductParam" :key="index" :class="{ littleMarginTop: index !== 0 }">
                <div class="paramInputLabel">{{ item.name }}:</div>
                <el-select v-if="item.inputType === 1" class="paramInput" v-model="selectProductParam[index].value">
                  <el-option v-for="item in getParamInputList(item.inputList)" :key="item" :label="item" :value="item">
                  </el-option>
                </el-select>
                <el-input v-else class="paramInput" v-model="selectProductParam[index].value"></el-input>
              </div>
            </el-card>
          </el-form-item>

          <el-form-item label="商品相册">
            <multi-upload v-model="selectProductPics"></multi-upload>
          </el-form-item>
          <el-form-item label="商品详情">
            <el-tabs v-model="activeHtmlName" type="card">
              <el-tab-pane label="电脑端详情" name="pc">
                <tinymce :width="595" :height="300" v-model="productDetail!.detailHtml"></tinymce>
              </el-tab-pane>
              <el-tab-pane label="移动端详情" name="mobile">
                <tinymce :width="595" :height="300" v-model="productDetail!.detailMobileHtml"></tinymce>
              </el-tab-pane>
            </el-tabs>
          </el-form-item>
        </div>

        <!-- 选择商品关联 -->
        <div v-show="activeStep === 3">
          <el-form-item label="关联专题">
            <el-transfer
              style="display: inline-block"
              :filterable="true"
              :filter-method="filterMethod"
              filter-placeholder="请输入专题名称"
              v-model="selectSubject"
              :titles="subjectTitles"
              :data="subjectList"
            >
            </el-transfer>
          </el-form-item>
          <el-form-item label="关联优选">
            <el-transfer
              style="display: inline-block"
              filterable
              :filter-method="filterMethod"
              filter-placeholder="请输入优选名称"
              v-model="selectPrefrenceArea"
              :titles="prefrenceAreaTitles"
              :data="prefrenceAreaList"
            >
            </el-transfer>
          </el-form-item>
        </div>
      </el-form>
    </div>
    <template #footer>
      <span style="float: left">
        <el-button :disabled="activeStep == 0" type="primary" @click="handlePrev('productInfoForm')">上一步</el-button>
        <el-button :disabled="activeStep == 3" type="primary" @click="handleNext('productInfoForm')">下一步</el-button>
      </span>
      <span>
        <el-button @click="drawerVisible = false">取消</el-button>
        <el-button :disabled="activeStep != 3 || drawerProps.isView" type="primary" @click="handleSubmit">确定</el-button>
      </span>
    </template>
  </el-drawer>
</template>

<script setup lang="ts" name="ProductDrawer">
import { computed, onBeforeMount, reactive, ref } from "vue";
import { ElMessage, FormInstance } from "element-plus";
import { CascaderValue, OptionValue, Product, TransferValue } from "@/api/interface";
import { getCategoryOptionsApi } from "@/api/modules/category";
import { getBrandOptionsApi } from "@/api/modules/brand";
import { getProductSyncApi } from "@/api/modules/product";
import { getProductAttributeCategoryOptionsApi } from "@/api/modules/productAttributeCategory";
import { getAllSubjectTransferValueSyncApi } from "@/api/modules/subject";
import { getAllPrefrenceAreaTransferValueSyncApi } from "@/api/modules/prefrenceArea";

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

const filterMethod = (query: string, item: any) => {
  return item.label.indexOf(query) > -1;
};

// 专题左右标题
const subjectTitles = ref<[string, string]>(["待选择", "已选择"]);
// 所有专题列表
const subjectList = ref<TransferValue[]>([]);
// 商品专题设置
const selectSubject = computed({
  get: () => {
    let value = productDetail.value;
    let subjects: number[] = [];
    if (value === undefined) {
      return subjects;
    }
    if (value.subjectProductRelations.length <= 0) {
      return subjects;
    }
    for (let i = 0; i < value.subjectProductRelations.length; i++) {
      subjects.push(value.subjectProductRelations[i].subjectId);
    }
    return subjects;
  },
  set: newValue => {
    let value = productDetail.value;
    if (value === undefined) {
      return;
    }
    value.subjectProductRelations = [];
    for (let i = 0; i < newValue.length; i++) {
      value.subjectProductRelations.push({ subjectId: newValue[i] });
    }
  }
});

// 优选左右标题
const prefrenceAreaTitles = ref<[string, string]>(["待选择", "已选择"]);
// 所有优选列表
const prefrenceAreaList = ref<TransferValue[]>([]);
// 优选设置
const selectPrefrenceArea = computed({
  get: () => {
    let value = productDetail.value;
    let prefrenceAreas: number[] = [];
    if (value === undefined) {
      return prefrenceAreas;
    }
    if (value.prefrenceAreaProductRelations.length <= 0) {
      return prefrenceAreas;
    }
    for (let i = 0; i < value.prefrenceAreaProductRelations.length; i++) {
      prefrenceAreas.push(value.prefrenceAreaProductRelations[i].prefrenceAreaId);
    }
    return prefrenceAreas;
  },
  set: newValue => {
    let value = productDetail.value;
    if (value === undefined) {
      return;
    }
    value.prefrenceAreaProductRelations = [];
    for (let i = 0; i < newValue.length; i++) {
      value.prefrenceAreaProductRelations.push({ prefrenceAreaId: newValue[i] });
    }
  }
});

// 表单引用
const ruleFormRef = ref<FormInstance>();

// 当前tab
const activeStep = ref<number>(0);
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

// 所有分类（树形）
const categoryOptions = ref<CascaderValue[]>([]);
// 所有的品牌
const brandOptions = ref<OptionValue[]>([]);
// 所有商品属性分类
const productAttributeCategoryOptions = ref<OptionValue[]>([]);
// 商品详情
const productDetail = ref<Product.ProductModel | undefined>(undefined);

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
const acceptParams = async (params: DrawerProps) => {
  // 获取详情
  if (params.row.id && params.row.id !== 0) {
    productDetail.value = await getProductSyncApi(params.row.id);
  }
  drawerProps.value = params;
  drawerVisible.value = true;
  activeStep.value = 0;
};

// 服务保证，存储是逗号分隔，展示是多选框
const selectServiceList = computed({
  get() {
    let value = drawerProps.value.row;
    let list: number[] = [];
    if (value.serviceIds === undefined || value.serviceIds == null || value.serviceIds === "") return list;
    let ids = value.serviceIds.split(",");
    for (let i = 0; i < ids.length; i++) {
      list.push(Number(ids[i]));
    }
    return list;
  },
  set(newValue) {
    let value = drawerProps.value.row;
    let serviceIds = "";
    if (newValue != null && newValue.length > 0) {
      for (let i = 0; i < newValue.length; i++) {
        serviceIds += newValue[i] + ",";
      }
      // 移除末尾逗号
      if (serviceIds.endsWith(",")) {
        serviceIds = serviceIds.slice(0, -1);
      }
      value.serviceIds = serviceIds;
    } else {
      value.serviceIds = "";
    }
  }
});

// 时间戳和Date相互转换
const promotionStartTime = computed({
  get() {
    return drawerProps.value.row?.promotionStartTime ? new Date(drawerProps.value.row.promotionStartTime * 1000) : new Date();
  },
  set(newValue: Date) {
    if (drawerProps.value.row) {
      drawerProps.value.row.promotionStartTime = newValue ? Math.floor(newValue.getTime() / 1000) : 0;
    }
  }
});

const promotionEndTime = computed({
  get() {
    return drawerProps.value.row?.promotionEndTime ? new Date(drawerProps.value.row.promotionEndTime * 1000) : new Date();
  },
  set(newValue: Date) {
    if (drawerProps.value.row) {
      drawerProps.value.row.promotionEndTime = newValue ? Math.floor(newValue.getTime() / 1000) : 0;
    }
  }
});

// 时间选择选项
const pickerOptions = reactive({
  disabledDate(time: Date) {
    return time.getTime() < Date.now();
  }
});

// 删除阶梯价格
const handleRemoveProductLadder = (index: number) => {
  let productLadders = productDetail.value!.productLadders;
  if (productLadders.length === 1) {
    productLadders.pop();
    productLadders.push({
      count: 0,
      discount: 0,
      price: 0
    });
  } else {
    productLadders.splice(index, 1);
  }
};
// 添加阶梯价格
const handleAddProductLadder = () => {
  let productLadders = productDetail.value!.productLadders;
  if (productLadders.length < 3) {
    productLadders.push({
      count: 0,
      discount: 0,
      price: 0
    });
  } else {
    ElMessage.warning("最多只能添加三条");
  }
};

// 添加满减价格
const handleAddFullReduction = () => {
  let productFullReductions = productDetail.value!.productFullReductions;
  if (productFullReductions.length < 3) {
    productFullReductions.push({
      fullPrice: 0,
      reducePrice: 0
    });
  } else {
    ElMessage.warning("最多只能添加三条");
  }
};

// 删除满减价格
const handleRemoveFullReduction = (index: number) => {
  let productFullReductions = productDetail.value!.productFullReductions;
  if (productFullReductions.length === 1) {
    productFullReductions.pop();
    productFullReductions.push({
      fullPrice: 0,
      reducePrice: 0
    });
  } else {
    productFullReductions.splice(index, 1);
  }
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
  categoryOptions.value = await getCategoryOptionsApi();
  brandOptions.value = await getBrandOptionsApi();
  productAttributeCategoryOptions.value = await getProductAttributeCategoryOptionsApi();
  subjectList.value = await getAllSubjectTransferValueSyncApi();
  prefrenceAreaList.value = await getAllPrefrenceAreaTransferValueSyncApi();
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

.littleMargin {
  margin-top: 10px;
}

.littleMarginLeft {
  margin-left: 10px;
}

.littleMarginTop {
  margin-top: 10px;
}

.paramInput {
  width: 250px;
}

.paramInputLabel {
  display: inline-block;
  width: 100px;
  text-align: right;
  padding-right: 10px;
}

.cardBg {
  background: #f8f9fc;
}
</style>
