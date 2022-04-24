import { createApp } from 'vue'
import App from './App.vue'

// 状态
import store from './store'

// 简易路由
import router from './router'

// 设置icon
import installIcon from './icon/index.js'

// UI库
import ElementPlus from 'element-plus'
// import 'element-plus/lib/theme-chalk/index.css'
// import 'dayjs/locale/zh-cn'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

// 二次封装
import { nfElementPlus, dialogDrag } from '@naturefw/ui-elp'

// 数据库
import db from './storage'

db()

createApp(App).use(router) // 路由
  .use(store)
  .use(installIcon) // 注册全局图标
  .use(ElementPlus, { locale: zhCn, size: 'small' }) // UI库
  .use(nfElementPlus) // 二次封装的组件
  .mount('#app')
