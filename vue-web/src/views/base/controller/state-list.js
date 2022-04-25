
import { watch } from 'vue'

import { defineStore, useStore, store } from '@naturefw/nf-state'

const flag = Symbol('pager001')
// const flag = 'pager001'

/**
 * 注册局部状态，用 provide 注入
 * 数据列表用
 * @returns
 */
const regListState = () => {
  // 定义 列表用的状态
  const state = defineStore(flag, {
    state: () => {
      return {
        moduleId: 0, // 模块ID
        dataList: [], // 数据列表
        findValue: {}, // 查询条件的精简形式
        findArray: [], // 查询条件的对象形式
        pagerInfo: { // 分页信息
          pagerSize: 5,
          count: 20, // 总数
          pagerIndex: 1 // 当前页号
        },
        selection: { // 列表里选择的记录
          dataId: '', // 单选ID number 、string
          row: {}, // 单选的数据对象 {}
          dataIds: [], // 多选ID []
          rows: [] // 多选的数据对象 []
        },
        query: {} // 查询条件
      }
    },
    getters: {
      nameTest () {
        return this.name + '_通过 getter 获取'
      }
    },
    actions: {
      /**
       * 加载数据，
       * @param {*} isReset true：设置总数，页号设置为1；false：仅翻页
       */
      async loadData (isReset = false) {
        const dbHelp = store.web.dbHelp
        const pager = {
          count: this.pagerInfo.pagerSize,
          start: (this.pagerInfo.pagerIndex - 1) * this.pagerInfo.pagerSize
        }
        const list = await dbHelp.listPager('db_fq_log', this.findValue, pager)
        this.dataList.length = 0
        if (list.dataList.length > 0) {
          this.dataList.push(...list.dataList)
        }
        if (isReset) {
          this.pagerInfo.count = list.allCount === 0 ? 1 : list.allCount
          this.pagerInfo.pagerIndex = 1
        }
      }
    }
  },
  { isLocal: true }
  )

  // 初始化
  state.loadData(true)

  // 分页功能
  watch(() => state.pagerInfo.pagerIndex, (index) => {
    state.loadData()
  })

  // 更换查询条件
  watch(state.findValue, () => {
    state.loadData(true)
  })

  return state
}

/**
 * 用 inject 获取状态
 * @returns
 */
const getListState = () => {
  return useStore(flag)
}

export {
  getListState,
  regListState
}
