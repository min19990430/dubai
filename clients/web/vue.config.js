const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  //修改網頁title
  chainWebpack: config => {
    config.plugin('define').tap((definitions) => {
      Object.assign(definitions[0], {
        __VUE_OPTIONS_API__: 'true',
        __VUE_PROD_DEVTOOLS__: 'false',
        __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: 'false'
      })
      return definitions
    })
    config
      .plugin('html')
      .tap(args => {
        args[0].title = "Dubai - Piloting Project 2024";
        return args
      })
  },
  transpileDependencies: true
})
