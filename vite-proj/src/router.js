import { createRouter, createWebHistory } from 'vue-router'
// 3. 创建路由实例并传递 `routes` 配置
// 你可以在这里输入更多的配置，但我们在这里
// 暂时保持简单
export const router =  createRouter({
  history: createWebHistory(),
  routes: [
    { 
      path: '/note',
      component: () => import('./pages/Note.vue'),
    },
    { 
      path: '/md',
      component: () => import('./pages/MD.vue'),
    },
    { 
      path: '/note/:logid',
      props: true,
      component: () => import('./pages/Note.vue'),
    },
    { 
      path: '/table',
      component: () => import('./pages/Table.vue'),
    },
  ],
})

 