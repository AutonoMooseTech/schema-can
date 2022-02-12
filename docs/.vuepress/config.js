module.exports = {
	// site
	lang: 'en-US',
	title: 'SchemaCAN',
	description: 'An open standard for a CAN Bus',
	base: '/schema-can/',

	// theme
	theme: '@vuepress/theme-default',

	// plugins
	plugins: [
    [
      '@vuepress/plugin-search',
      {
        locales: {
          '/': {
            placeholder: 'Search',
          },
          '/zh/': {
            placeholder: '搜索',
          },
        },
      },
    ],
  ],
}