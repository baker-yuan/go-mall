// sidebar.ts
export default {
  '/document/': getDocumentSidebar(),
  '/table/': getTableSidebar(),
  '/recommend/': getRecommendSidebar(),
}

// 推荐
function getRecommendSidebar () {
  return [
    {
      text: '商业版',
      link: '/recommend/index.md',
      children: [

      ],
    },
    {
      text: '学习类',
      children: [
        { text: 'java实现&适合学习', link: '/recommend/macrozheng_mall.md' },
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