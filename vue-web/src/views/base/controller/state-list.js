
import { watch } from 'vue'

import { defineStore, useStore } from '@naturefw/nf-state'

const flag = Symbol('pager001')
// const flag = 'pager001'

/**
 * 注册局部状态，用 provide 注入
 * @returns 
 */
const regListState = () => {

  // 定义 列表用的状态
  const state = defineStore(flag, {
    state: {
      moduleId: 0,
      dataList: [],
      findValue: {},
      findArray: [],
      pagerInfo: {
        pagerSize: 10,
        count: 20, // 总数
        pagerIndex: 1 // 当前页号
      },
      query: {} // 查询条件
    },
    getters: {
      nameTest() {
        return this.name + '_通过 getter 获取'
      }
    },
    actions: {
      updateName(val) {
        this.name = val
      },
      setPager(info) {
        this.$patch(info)
      }
    }
  },{isLocal: true})
  
  // 设置联动

  // 分页功能

  watch(() => state.pagerInfo.pageIndex, (index) => {
    // console.log(index)
    const pager = state.pagerInfo
    state.dataList.length = 0
    const cpp = (pager.pageIndex - 1) * pager.pageSize
    for(let i=0; i< pager.pageSize; i++) {
      state.dataList.push({})
    }
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