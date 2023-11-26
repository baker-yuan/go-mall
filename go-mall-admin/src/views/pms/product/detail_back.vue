<template>
  <el-drawer v-model="drawerVisible" :destroy-on-close="true" size="1100px" :title="`${drawerProps.title}商品`">
    <el-card class="form-container" shadow="never">
      <el-steps :active="activeStep" finish-status="success" align-center>
        <el-step title="填写商品信息"></el-step>
        <el-step title="填写商品促销"></el-step>
        <el-step title="填写商品属性"></el-step>
        <el-step title="选择商品关联"></el-step>
      </el-steps>

      <div style="margin-top: 50px">
        <el-form :model="drawerProps.row" :rules="productInfoDetailRules" ref="productInfoForm" label-width="120px" class="form-inner-container" size="small">




        </el-form>


      </div>
    </el-card>
  </el-drawer>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { Product } from "@/api/interface";

// const defaultProductParam: Product.AddOrUpdateProductModel = {
//   // 基本信息
//   id: 0, // 主键
//   productCategoryId: 0, // 商品分类id
//   name: "", // 商品名称
//   subTitle: "", // 副标题
//   brandId: 0, // 品牌id
//   description: "", // 商品描述
//   productSn: "", // 货号
//   price: 0, // 价格
//   originalPrice: 0, // 市场价
//   stock: 0, // 库存
//   unit: "", // 单位
//   weight: 0, // 商品重量，默认为克
//   sort: 0, // 排序
//
//   // 促销信息
//   giftPoint: 0, // 赠送的积分
//   giftGrowth: 0, // 赠送的成长值
//   usePointLimit: 0, // 限制使用的积分数
//   previewStatus: 0, // 是否为预告商品：0->不是；1->是
//   publishStatus: 0, // 上架状态：0->下架；1->上架
//   newStatus: 0, // 新品状态:0->不是新品；1->新品
//   recommandStatus: 0, // 推荐状态；0->不推荐；1->推荐
//   serviceIds: "", // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
//   detailTitle: "", // 详情标题
//   detailDesc: "", // 详情描述
//   keywords: "", // 关键字
//   note: "", // 备注
//   promotionType: 0, // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
//   promotionPrice: 0, // 促销价格
//   promotionStartTime: 0, // 促销开始时间
//   promotionEndTime: 0, // 促销结束时间
//
//   // 属性信息
//   productAttributeCategoryId: 0, // 品牌属性分类id
//   pic: "", // 图片
//   albumPics: "", // 画册图片，连产品图片限制为5张，以逗号分割
//   detailHTML: "", // 电脑端详情
//   detailMobileHTML: "", // 移动端详情
//
//   // 状态
//   verifyStatus: 0, // 审核状态：0->未审核；1->审核通过
//   deleteStatus: 0, // 删除状态：0->未删除；1->已删除
//
//   // 其他
//   feightTemplateId: 0, // 运费模版id
//   sale: 0, // 销量
//   lowStock: 0, // 库存预警值
//   promotionPerLimit: 0, // 活动限购数量
//
//   // 冗余字段
//   brandName: "", // 品牌名称
//   productCategoryName: "", // 商品分类名称
//
//   // 关联
//   productLadders: [], // 商品阶梯价格设置
//   productFullReductions: [], // 商品满减价格设置
//   memberPrices: [], // 商品会员价格设置 {"memberLevelId":0,"memberPrice":0,"memberLevelName":""}
//   skuStocks: [], // 商品sku库存信息 {"lowStock":0,"pic":"","price":0,"sale":0,"skuCode":"","spData":"","stock":0}
//   productAttributeValues: [], // 商品属性设置 {"productAttributeId":0,"value":""}
//   subjectProductRelations: [], // 商品专题设置 {"subjectId":0}
//   prefrenceAreaProductRelations: [] // 商品优选设置 {"prefrenceAreaId":0}
// };

const productInfoDetailRules = reactive({
  name: [
    {required: true, message: '请输入商品名称', trigger: 'blur'},
    {min: 2, max: 140, message: '长度在 2 到 140 个字符', trigger: 'blur'}
  ],
  subTitle: [{required: true, message: '请输入商品副标题', trigger: 'blur'}],
  productCategoryId: [{required: true, message: '请选择商品分类', trigger: 'blur'}],
  brandId: [{required: true, message: '请选择商品品牌', trigger: 'blur'}],
  description: [{required: true, message: '请输入商品介绍', trigger: 'blur'}],
  requiredProp: [{required: true, message: '该项为必填项', trigger: 'blur'}]
});

// const nextStep = () => {
//   if (activeStep.value < showStatus.length - 1) {
//     activeStep.value = activeStep.value + 1;
//     hideAll();
//     showStatus[activeStep.value] = true;
//   }
// };
//
// const prevStep = () => {
//   if (activeStep.value > 0 && activeStep.value < showStatus.length) {
//     activeStep.value = activeStep.value - 1;
//     hideAll();
//     showStatus[activeStep.value] = true;
//   }
// };

// const hideAll = () => {
//   for (let i = 0; i < showStatus.length; i++) {
//     showStatus[i] = false;
//   }
// };

// const finishCommit = () => {
//   console.log("finishCommit");
// };

const activeStep = ref<number>(0);

// const showStatus = reactive<boolean[]>([true, false, false, false]);

// 是否显示弹框
const drawerVisible = ref(false);

interface DrawerProps {
  title: string; // 标题
  isView: boolean; // 是否查看
  row: Partial<Product.ProductModel>; // 数据
  api?: (params: any) => Promise<any>; // 提交api
  getTableList?: () => void; // 刷新列表
}

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
