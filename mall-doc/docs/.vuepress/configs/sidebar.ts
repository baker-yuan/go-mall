// sidebar.ts
export default {
  '/document/': getDocumentSidebar(),
  '/table/': getTableSidebar(),
}

// 文档搭建
function getDocumentSidebar() {
  return [
    // {
    //   link: '/document/document.md',
    // },
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
        { text: '商品品牌表', link: '/table/pms/pms_brand.md' },
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