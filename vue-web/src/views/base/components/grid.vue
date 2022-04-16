<template>
  <nf-grid
    v-grid-drag="gridMeta"
    :dataList="state.dataList"
    v-bind="gridMeta"
    :selection="state.selection"
  >
  </nf-grid>
</template>

<script>
  import { store } from '@naturefw/nf-state'
  // 局部状态
  import { getListState } from '../controller/state-list.js'

</script>

<script setup>
  // 获取列表状态
  const state = getListState()

  // 获取表单控件需要的meta
  const { gridMeta } = store.meta[state.moduleId]
</script>
