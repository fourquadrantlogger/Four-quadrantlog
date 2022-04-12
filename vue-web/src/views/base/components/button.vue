<template>
  <nf-button
    v-bind="buttonMeta"
    :componentKind="componentKind"
    :active="active"
    :moduleId="moduleId"
    :dataId="state.selection.dataId"
    :dataIds="state.selection.dataIds"
  >
  </nf-button>
</template>

<script>
  import { reactive, defineAsyncComponent } from 'vue'
  import { store } from '@naturefw/nf-state'

  // 局部状态
  import { getListState } from '../controller/state-list.js'

</script>

<script setup>
  const props = defineProps({
    moduleId: Number,
    active: Object
  })

  // 获取列表状态
  const state = getListState()

  const { buttonMeta } = store.meta[state.moduleId]

  buttonMeta.events = {
    deleteClick: (id, meta) => {
      // console.log('按了删除按钮：', id, meta)
      return true
    }
  }

  // 弹窗打开的组件的字典
  const componentKind = {
    form: defineAsyncComponent(() => import('./form.vue')),
    list: defineAsyncComponent(() => import('./pager.vue'))
  }

</script>
