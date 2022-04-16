<template>
  <nf-find
    :findValue="state.findValue"
    :findArray="state.findArray"
    :active="active"
    :moduleId="moduleId"
    v-bind="findMeta"
  >
  </nf-find>
</template>

<script>
  import { reactive, watch, onMounted, nextTick } from 'vue'
  import { store } from '@naturefw/nf-state'
  // 局部状态
  import { getListState } from '../controller/state-list.js'

</script>

<script setup>

  const props = defineProps({
    moduleId: Number,
    active: Object
  })

  const state = getListState()

  // 获取表单控件需要的meta
  const { findMeta } = store.meta[state.moduleId]

</script>
