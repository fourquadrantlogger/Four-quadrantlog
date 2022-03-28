import {router} from './router.js'

import { createApp,h } from 'vue'

import App from './App.vue';

// 5. 创建并挂载根实例
const app  = createApp({
  render: ()=>h(App)
});

//确保 _use_ 路由实例使
//整个应用支持路由。
app.use(router)
// 注意：在服务器端，你需要手动跳转到初始地址。
router.isReady().then(() => app.mount('#app'))