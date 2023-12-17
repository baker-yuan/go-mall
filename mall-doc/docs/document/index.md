# 1、环境
```text
nvm use v18.16.1
```

# 2、启动
```text
pnpm docs:dev
pnpm docs:build
```

# 3、官方文档
```text
https://v2.vuepress.vuejs.org/zh/
```




# 4、标准目录结构
VuePress 2 是一个基于 Vue 3 的静态网站生成器，它的标准目录结构如下所示：
```
.
├── docs
│   ├── .vuepress
│   │   ├── components       # 存放 Vue 组件
│   │   ├── public           # 存放静态资源，如图片、样式表等
│   │   ├── styles           # 存放样式文件
│   │   │   ├── index.scss   # 全局样式文件
│   │   │   └── palette.scss # 覆盖默认颜色变量
│   │   ├── templates        # 存放 HTML 模板文件（可选）
│   │   ├── config.js        # VuePress 配置文件
│   │   └── theme            # 存放自定义主题（可选）
│   ├── README.md            # 首页内容
│   ├── guide                # 存放指南文档（自己的文档）
│   │   └── README.md        # 指南首页内容
│   └── config.md            # 其他页面内容
├── package.json             # 项目依赖和脚本
└── yarn.lock                # Yarn 锁文件（如果使用 npm，则为 package-lock.json）
```

这是一个典型的 VuePress 2 项目结构，其中：
- `docs` 目录是默认的文档源文件夹，你可以在 VuePress 配置中更改它。
- `.vuepress` 目录包含 VuePress 特定的配置、组件、静态资源和样式。
- `components` 目录用于存放 Vue 组件，这些组件可以在 Markdown 文件中直接使用。
- `public` 目录用于存放不需要 Webpack 处理的静态资源。
- `styles` 目录包含自定义样式文件，可以用来覆盖默认主题的样式。
- `templates` 目录（可选）用于存放 HTML 模板文件，可以自定义 HTML 输出。
- `config.js` 文件是 VuePress 的配置文件，用于配置站点的标题、描述、主题、插件等。
- `theme` 目录（可选）用于存放自定义主题。
- `README.md` 文件是文档的首页内容，每个目录下的 `README.md` 文件都会被视为该目录的索引页面。
- `guide` 和 `config.md` 是示例文档页面，你可以根据需要创建更多的文档页面和目录。
- `package.json` 文件包含项目的依赖和脚本，例如启动本地开发服务器的脚本。
- `yarn.lock` 或 `package-lock.json` 文件用于锁定依赖的版本，确保其他人在安装依赖时能够得到相同的版本。

请注意，这个结构是可扩展的，你可以根据项目的需要添加更多的目录和文件。例如，你可能会有一个 `blog` 目录来存放博客文章，或者一个 `api` 目录来存放 API 文档。


# 4、参考案例
```text
https://github.com/woow-wu7/8-divine-plus
https://github.com/LiuyangAce/tt-ui

https://github.com/it235/it235-vuepress
https://www.bilibili.com/video/BV17t41177cr
```