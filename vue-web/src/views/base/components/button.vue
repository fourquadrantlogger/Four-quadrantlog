<template>
  <nf-button
    v-bind="buttonMeta"
    :componentKind="componentKind"
    :active="active"
    :moduleId="moduleId"
    :dataId="selection.dataId"
    :dataIds="selection.dataIds"
  >
  </nf-button>
</template>

<script>
  import { reactive, defineAsyncComponent } from 'vue'
  import _buttonMeta from '../../json/button.json'

</script>

<script setup>
  const props = defineProps({
    moduleId: Number,
    selection: Object,
    active: Object
  })

  const buttonMeta = reactive(_buttonMeta)

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