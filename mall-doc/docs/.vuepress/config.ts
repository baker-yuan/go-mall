import { defaultTheme } from 'vuepress'
import navbar from './configs/navbar'
import sidebar from './configs/sidebar'
import { demoblockPlugin } from 'vuepress-plugin-demoblock-plus'

export default {
  base: '/mall-doc/',
  lang: 'zh-CN',
  title: 'mall-doc',
  description: 'go+vue实现电商系统',
  theme: defaultTheme({
    navbar,
    sidebar,
    editLinkText: '在 GitHub 上编辑此页',
    lastUpdatedText: '上次更新',
    contributorsText: '贡献者',
  }),

  plugins: [
    demoblockPlugin({
      // locales,
      customClass: 'demoblock-custom',
      theme: 'github-light',
      cssPreprocessor: 'scss',
      // customStyleTagName: 'style lang="scss"', // style标签会解析为<style lang="scss"><style>
      scriptReplaces: [
        {
          searchValue: /const ({ defineComponent as _defineComponent }) = Vue/g,
          replaceValue: 'const { defineComponent: _defineComponent } = Vue',
        },
      ],
    })
  ]
}
