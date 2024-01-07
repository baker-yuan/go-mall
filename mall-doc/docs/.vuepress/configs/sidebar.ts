// sidebar.ts
export default {
  '/document/': getDocumentSidebar(),
  '/table/': getTableSidebar(),
  '/recommend/': getRecommendSidebar(),
  '/code/': getCodeSidebar(),
}

function getCodeSidebar() {
  return [

  ]
}

// 推荐
function getRecommendSidebar () {
  return [
    {
      text: '项目介绍',
      link: '/recommend/lilishop/index.md',
    },
    {
      text: '页面效果',
      link: '/recommend/lilishop/show.md'
    },
    {
      text: '核心表设计',
      link: '/recommend/lilishop/',
      children: [
        { text: '分类', link: '/recommend/lilishop/table/category.md' },
        { text: '商品', link: '/recommend/lilishop/table/product.md' },
        { text: '购物车', link: '/recommend/lilishop/table/cart.md' },
        { text: '用户', link: '/recommend/lilishop/table/user.md' },
      ],
    },
  ]
}

// 文档搭建
function getDocumentSidebar() {
  return [
    {
      text: '搭建vuepress',
      link: '/document/index.md',
      children: [
        { text: '搭建文档', link: '/document/init/init.md' },
        { text: '导航栏', link: '/document/config/bar.md' },
        { text: '首页', link: '/document/config/home_page.md' },
        { text: '设置代码拷贝', link: '/document/config/code_copy.md' },
      ],
    },
    {
      text: '部署vuepress',
      children: [
        { text: '域名+nginx', link: '/document/deployment/nginx.md' },
        { text: 'github.io', link: '/document/deployment/github.md' },
      ],
    },
  ]
}

// 表设计
function getTableSidebar() {
  return [
    {
      isGroup: true,
      text: '商品(pms)',
      link: '/table/oms/oms.md',
      children: [
        { text: '商品分类', link: '/table/pms/pms_product_category.md' },
        { text: '商品品牌', link: '/table/pms/pms_brand.md' },
        { text: '商品类型', link: '/table/pms/pms_attribute.md' },
        { text: '商品', link: '/table/pms/pms_product.md' },
      ],
    },
    {
      isGroup: true,
      text: '订单(oms)',
      link: '/table/oms/oms.md',
      children: [
        { text: '订单', link: '/table/oms/oms_order.md' },
      ],
    },
  ]
}