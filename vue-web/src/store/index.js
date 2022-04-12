// 加载状态的类库
import { createStore } from '@naturefw/nf-state'

import { loadMeta } from '../storage/service-meta.js'

const _baseUrl = (document.location.host.includes('.gitee.io')) ? '/four-quadrant/' : '/'

/**
 * 统一注册全局状态
 */
export default createStore({
  // 读写状态，直接使用 reactive
  store: {
    // 纯 state
    web: {
      state: () => {
        return {
          ver: 1.0,
          baseUrl: _baseUrl,
          load: 'start',
          dbHelp: null, // indexedDB 的help
          name: 'jyk' //
        }
      }
    },
    // 项目运行需要的 meta 信息
    meta: {
      0: {
        moduleId: 0,
        grid: {},
        pager: {},
        form: {},
        find: {},
        button: {}
      }
    },
    // 有 getters、actions
    log_list: {
      state: [],
      getters: {
      },
      actions: {
        async loadData (val, state) {
          const foo = await testPromie()
          state.name = foo
          this.name = foo
        }
      },
      options: {
        isLocal: false, // true：局部状态；false：全局状态（默认属性）；
        isLog: true, // true：做记录；false：不用做记录（默认属性）；
        /**
         * 1：宽松，可以各种改变属性，适合弹窗、抽屉、多tab切换等。
         * 2：一般，不能通过属性直接改状态，只能通过内置函数、action 改变状态
         * 3：严格，不能通过属性、内置函数改状态，只能通过 action 改变状态
         * 4：超严，只能在指定组件内改变状态，比如当前用户的状态，只能在登录组件改，其他组件完全只读
        */
        level: 1
      }
    }
  },

  // 可以给全局状态设置初始状态，同步数据可以直接在上面设置，如果是异步数据，可以在这里设置。
  async init (store) {
    await loadMeta('m_log')
    store.web.load = 'end'
  }
})
