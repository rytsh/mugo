export default {
  title: "mugo",
  description: "documentation",
  base: "/mugo/",
  markdown: {
    theme: "dracula",
  },
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
          { text: 'List', link: '/functions/list.md' },
          { text: 'Reference', link: '/functions/reference.md' },
        ]
      },
      {
        text: 'Go Template',
        collapsed: false,
        items: [
          { text: 'Intro', link: '/templates/intro.md' },
          { text: 'Examples', link: '/templates/examples.md' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/rytsh/mugo' }
    ]
  }
}
