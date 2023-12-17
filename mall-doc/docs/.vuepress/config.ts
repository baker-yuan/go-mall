// docs/.vuepress/config.ts

import { defaultTheme } from 'vuepress'
import navbar from './configs/navbar'
import sidebar from './configs/sidebar'
import { demoblockPlugin } from 'vuepress-plugin-demoblock-plus'

export default {
  // 基础路径，用于部署到非根 URL。
  base: '',
  // 网站使用的语言
  lang: 'zh-CN',
  // 网站的标题
  title: 'baker-mall',
  // 网站的描述
  description: 'go+vue+uni-app实现完整的电商系统',
  // 主题配置
  theme: defaultTheme({
    // 导航栏配置
    navbar,
    // 侧边栏配置
    sidebar,
    // 编辑链接文本，显示在每个页面底部，提示用户编辑当前页面。
    editLinkText: '在 GitHub 上编辑此页',
    // 最后更新文本，显示在每个页面底部，表示页面的最后更新时间。
    lastUpdatedText: '上次更新',
    // 贡献者文本，显示在每个页面底部，列出对当前页面做出贡献的人员。
    contributorsText: '贡献者',
  }),
  // 插件配置数组，可以添加多个插件。
  plugins: [
    // demoblock 插件，用于在 VuePress 文档中展示 Vue 组件的示例和代码块。
    demoblockPlugin({
      // 主题，这里使用的是 GitHub 的浅色主题。
      theme: 'github-light',
      // CSS 预处理器，这里指定使用 SCSS。
      cssPreprocessor: 'scss',
    })
  ],
}
