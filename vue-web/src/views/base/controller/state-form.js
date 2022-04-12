
import { watch } from 'vue'

import { defineStore, useStore, store, deepClone } from '@naturefw/nf-state'


const flag = Symbol('form001')
// const flag = 'pager001'

/**
 * 注册局部状态，用 provide 注入。
 * 表单用
 * @returns 
 */
const regFormState = () => {
  const dbHelp = store.web.dbHelp

  // 定义 表单用的状态
  const state = defineStore(flag, {
    state: {
      moduleId: 0,
      model: {
        id: '1', // 
        title: '标题', // 100 标题
        quadrant: '', // 101 象限
        atype: '分类', // 102 分类
        detail: '内容', // 103 内容
        review: '思考', // 104 思考
        blob: '', // 105 图片
        location: '', // 106 
        ctime: new Date() // 110 
      },
      state: 'add' // add、update、、
    },
    getters: {
      nameTest() {
        return this.name + '_通过 getter 获取'
      }
    },
    actions: {
      submit(val) {
        // 判断是添加还是修改
        this.name = val
        // 添加
        console.log(dbHelp)
        this.model.id = new Date().valueOf()
        const _model = deepClone({}, this.model)
        _model.quadrant = this.model.quadrant.join(',')

        dbHelp.addModel('db_fq_log', _model).then((res) => {
          console.log('添加成功', res)
        })
      }
    }
  },{isLocal: true})
  
  return state

}

/**
 * 用 inject 获取状态
 * @returns 
 */
const getFormState = () => {
  return useStore(flag)
}


export {
  getFormState,
  regFormState
}