// navbar.ts
export default [
  {
    text: '首页',
    link: '/',
  },
  { text: '表设计', link: '/table/index.md' },
  { text: '环境搭建', link: '/environment/index.md' },
  { text: '代码讲解', link: '/code/index.md' },
  { text: '文档搭建', link: '/document/index.md' },
  { text: '商城推荐', link: '/recommend/index.md' },
  {
    text: '其他',
    children: [
      {
        text: 'gitee',
        link: 'https://gitee.com/baker-yuan/go-mall'
      },
      {
        text: 'github',
        link: 'https://github.com/baker-yuan/go-mall'
      },
      {
        text: '个人网站',
        link: 'http://www.baker-yuan.cn/'
      },
      {
        text: 'csdn',
        link: 'https://yuanyu.blog.csdn.net/'
      }
    ]
  },
]