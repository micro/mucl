import { defineAdditionalConfig, type DefaultTheme } from 'vitepress'


// // https://vitepress.dev/reference/site-config
// export default defineConfig({
//   title: "mucl",
//   description: "Developer tools for go-micro",
//   themeConfig: {
//     // https://vitepress.dev/reference/default-theme-config
//     nav: [
//       { text: 'Home', link: '/' },
//       { text: 'Examples', link: '/markdown-examples' }
//     ],

//     sidebar: [
//       {
//         text: 'Examples',
//         items: [
//           { text: 'Markdown Examples', link: '/markdown-examples' },
//           { text: 'Runtime API Examples', link: '/api-examples' }
//         ]
//       }
//     ],

//     socialLinks: [
//       { icon: 'github', link: 'https://github.com/micro/mucl' }
//     ]
//   }
// })



export default defineAdditionalConfig({
  description: 'Developer tools for go-micro',
  title: 'mucl',
  cleanUrls: true,
  base: '/mucl/',


  themeConfig: {
    nav: nav(),

    sidebar: {
      '/guide/': { base: '/guide/', items: sidebarGuide() },
      '/reference/': { base: '/reference/', items: sidebarReference() }
    },

    editLink: {
      pattern: 'https://github.com/micro/mucl/edit/main/docs/:path',
      text: 'Edit this page on GitHub'
    },

    footer: {
      message: 'mucl is part of go-micro.',
      copyright: 'Copyright © 2015-present Asim Aslam'
    }
  }
})

function nav(): DefaultTheme.NavItem[] {
  return [
    {
      text: 'Guide',
      link: '/guide/what-is-mucl',
      activeMatch: '/guide/'
    },
    {
      text: 'Reference',
      link: '/reference/mucl',
      activeMatch: '/reference/'
    },
    {
      text: "mucl",
      items: [
        {
          text: 'Contributing',
          link: 'https://github.com/micro/mucl/blob/main/CONTRIBUTING.md'
        }
      ]
    }
  ]
}

function sidebarGuide(): DefaultTheme.SidebarItem[] {
  return [
    {
      text: 'Introduction',
      collapsed: false,
      items: [
        { text: 'What is mucl?', link: 'what-is-mucl' },
        { text: 'Install mucl', link: 'installation' },
        { text: 'Getting Started', link: 'getting-started' }
      ]
    },

    { text: 'MuCL Reference', base: '/reference/', link: 'mucl' }
  ]
}

function sidebarReference(): DefaultTheme.SidebarItem[] {
  return [
    {
      text: 'Reference',
      items: [
        { text: 'MuCL', link: 'mucl' },
      ]
    }
  ]
}
