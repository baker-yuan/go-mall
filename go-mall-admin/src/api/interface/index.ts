// 请求响应参数（不包含data）
export interface Result {
  code: string;
  msg: string;
}

// 请求响应参数（包含data）
export interface ResultData<T = any> extends Result {
  data: T;
}

// 分页响应参数
export interface ResPage<T> {
  data: T[];
  pageNum: number;
  pageSize: number;
  total: number;
}

// 分页请求参数
export interface ReqPage {
  pageNum: number;
  pageSize: number;
}

// 文件上传模块
export namespace Upload {
  export interface ResFileUrl {
    fileUrl: string;
  }
}

// 登录模块
export namespace Login {
  export interface ReqLoginForm {
    username: string;
    password: string;
  }
  export interface ResLogin {
    access_token: string;
  }
  export interface ResAuthButtons {
    [key: string]: string[];
  }
}

// 分类管理模块
export namespace Category {
  // 分类模型
  export interface CategoryModel {
    id: number, // 主键
    parentId: number, // 上级分类的编号：0表示一级分类
    name: string, // 名称
    level: number, // 分类级别：0->1级；1->2级
    productCount: number, // 商品数量
    productUnit: number, // 商品单位
    navStatus: number, // 是否显示在导航栏：0->不显示；1->显示
    showStatus: number, // 显示状态：0->不显示；1->显示
    sort: number, // 排序
    icon: string, // 图标
    keywords: string, // 关键字
    description: string, // 描述
  }
  // 分页查询分类
  export interface ReqCategoryListParams extends ReqPage {
    parentId: number, // 上级分类的编号：0表示一级分类
  }

}


// 品牌管理模块
export namespace Brand {
  // 品牌模型
  export interface BrandModel {
    id: number, // 主键
    name: string, // 名称
    firstLetter: string // 首字母
    sort: number // 排序
    factoryStatus: number // 是否为品牌制造商：0->不是；1->是
    showStatus: number // 是否显示
    productCount: number // 产品数量
    productCommentCount: number // 产品评论数量
    logo: string // 品牌logo
    bigPic: string // 专区大图
    brandStory: string // 品牌故事
  }

  // 分页查询商品品牌
  export interface ReqBrandListParams extends ReqPage {
    name: string, // 名称
  }

}


// 品牌管理模块
export namespace ProductAttributeCategory {
  // 品牌模型
  export interface ProductAttributeCategoryModel {
    id: number, // 主键
    name: string, // 类型名称
    attributeCount: number, // 属性数量
    paramCount: number // 参数数量
  }

  // 分页查询商品品牌
  export interface ReqProductAttributeCategoryListParams extends ReqPage {
    name: string, // 名称
  }

}


// 商品属性参数表管理模块
export namespace ProductAttribute {
  // 商品属性参数表模型
  export interface ProductAttributeModel {
    id: number // 编号
    productAttributeCategoryId: number // 产品属性分类表ID
    name: string // 属性名称
    selectType: number // 属性选择类型：0->唯一；1->单选；2->多选
    inputType: number // 属性录入方式：0->手工录入；1->从列表中选取
    inputList: string // 可选值列表，以逗号隔开
    sort: number // 排序字段
    filterType: number // 分类筛选样式：1->普通；1->颜色
    searchType: number // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
    relatedStatus: number // 相同属性产品是否关联；0->不关联；1->关联
    handAddStatus: number // 是否支持手动新增；0->不支持；1->支持
    type: number // 属性的类型；0->规格；1->参数
  }

  // 分页查询商品属性参数表
  export interface ReqProductAttributeListParams extends ReqPage {
  }

}

// 商品管理模块
export namespace Product {
  // 商品模型
  export interface ProductModel {
    id: number
  }

  // 分页查询商品
  export interface ReqProductListParams extends ReqPage {
  }

}
