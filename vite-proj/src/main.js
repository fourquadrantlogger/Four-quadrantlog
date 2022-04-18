import {router} from './router.js'

import { createApp,h } from 'vue'

import App from './App.vue';
 

import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

import 'element-plus/dist/index.css'
 
// 引入组件库
import {createPinia} from 'pinia'


// 5. 创建并挂载根实例
const app  = createApp({
  render: ()=>h(App)
});
const pinia=createPinia()
app.use(pinia)
//确保 _use_ 路由实例使
//整个应用支持路由。
app.use(router)
app.use(ElementPlus, {
  locale: zhCn,
})



// 注意：在服务器端，你需要手动跳转到初始地址。
router.isReady().then(() => app.mount('#app'))