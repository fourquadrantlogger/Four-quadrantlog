module.exports = {
  env: {
    browser: false,
    es2021: true
  },
  extends: [
    'plugin:vue/essential',
    'standard'
  ],
  parserOptions: {
    ecmaVersion: 12,
    sourceType: 'module'
  },
  plugins: [
    'vue'
  ],
  rules: {
    // 'no-v-model-argument': 'off',
    'import/no-absolute-path': 'off',
    'vue/script-indent': ['error', 2, { baseIndent: 1, switchCase: 0 }], // 缩进
    'vue/no-multiple-template-root': 'off', // template 里面可以有多个根
    'no-v-model-argument': 'off', // 分页的参数
    'no-absolute-path': 'off', // 不判断import 的路径，因为不识别别名
    'no-unused-vars': 'off', // 不判断是否被使用，因为 setup 里面没有return 了
    'no-undef': 'off' // dayjs
  },
  overrides: [
    {
      files: ['*.vue'],
      rules: {
        indent: 'off'
      }
    }
  ]
}
