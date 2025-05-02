import { defineAdditionalConfig, type DefaultTheme } from 'vitepress'


// // https://vitepress.dev/reference/site-config
// export default defineConfig({
//   title: "mu",
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
//       { icon: 'github', link: 'https://github.com/micro/mu' }
//     ]
//   }
// })



export default defineAdditionalConfig({
  description: 'Developer tools for go-micro',
  title: 'mu',
  cleanUrls: true,
  base: '/mu/',


  themeConfig: {
    nav: nav(),

    sidebar: {
      '/guide/': { base: '/guide/', items: sidebarGuide() },
      '/reference/': { base: '/reference/', items: sidebarReference() }
    },

    editLink: {
      pattern: 'https://github.com/micro/mu/edit/main/docs/:path',
      text: 'Edit this page on GitHub'
    },

    footer: {
      message: 'mu is part of go-micro.',
      copyright: 'Copyright © 2015-present Asim Aslam'
    }
  }
})

function nav(): DefaultTheme.NavItem[] {
  return [
    {
      text: 'Guide',
      link: '/guide/what-is-mu',
      activeMatch: '/guide/'
    },
    {
      text: 'Reference',
      link: '/reference/mucl',
      activeMatch: '/reference/'
    },
    {
      text: "mu",
      items: [
        {
          text: 'Contributing',
          link: 'https://github.com/micro/mu/blob/main/CONTRIBUTING.md'
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
        { text: 'What is mu?', link: 'what-is-mu' },
        { text: 'Install mu', link: 'installation' },
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