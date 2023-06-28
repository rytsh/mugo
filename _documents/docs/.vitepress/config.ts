import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "mugo",
  description: "documentation",
  base: "/mugo/",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Documentation', link: '/introduction/getting-started.md' }
    ],

    sidebar: [
      {
        text: 'Introduction',
        collapsed: false,
        items: [
          { text: 'Getting Started', link: '/introduction/getting-started.md' },
          { text: 'Go Module', link: '/introduction/go-module.md' },
          { text: 'Command Line Interface', link: '/introduction/cli.md' },
        ]
      },
      {
        text: 'Functions',
        collapsed: false,
        items: [
          { text: 'Reference', link: '/functions/reference.md' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/rytsh/mugo' }
    ]
  }
})
