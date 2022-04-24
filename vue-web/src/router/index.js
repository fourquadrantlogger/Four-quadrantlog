import { Document, FolderOpened } from '@element-plus/icons'
import { createRouter } from '@naturefw/ui-core'

import home from '../views/Home.vue'

export default createRouter({
  /**
   * 基础路径
   */
  baseUrl: '/four-quadrant',
  /**
   * 首页
   */
  home: home,

  menus: [
    {
      menuId: '1000',
      title: '时空',
      naviId: '0',
      path: 'base',
      icon: FolderOpened,
      childrens: [
        {
          menuId: '11',
          title: '原生HTML',
          path: 'html',
          icon: Document,
          component: () => import('../views/base/data-list.vue')
        }
      ]
    }
  ]
})
