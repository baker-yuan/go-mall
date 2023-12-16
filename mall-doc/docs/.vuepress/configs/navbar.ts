// navbar.ts
export default [
  { text: '表设计', link: '/table/table.md' },
  { text: '环境搭建', link: '/environment/index.md' },
  { text: '代码讲解', link: '/code/index.md' },
  { text: '文档搭建', link: '/document/index.md' },
  {
    text: '其他',
    children: [
      {
        text: 'gitee',
        link: 'https://gitee.com/baker-yuan'
      },
      {
        text: 'github',
        link: 'https://github.com/baker-yuan'
      }
    ]
  },
]