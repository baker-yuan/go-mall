// sidebar.ts
export default {
  '/document/': getDocumentSidebar(),
  '/table/': getTableSidebar(),
}

// 文档搭建
function getDocumentSidebar() {
  return [
    {
      text: '搭建vuepress',
      link: '/document/document.md',
      children: [
        { text: '设置代码拷贝', link: '/document/code_copy.md' },
        { text: '头部导航栏', link: '/document/nav_bar.md' },
        { text: '测边导航栏', link: '/document/side_bar.md' },
      ],
    },
    {
      text: '部署vuepress',
      children: [
        { text: '使用自己的域名', link: '/document/deployment/nginx.md' },
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