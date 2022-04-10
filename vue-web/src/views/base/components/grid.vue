<template>
   <nf-grid
    v-grid-drag="gridMeta"
    :dataList="dataList"
    v-bind="gridMeta"
    :selection="selection"
  >
  </nf-grid>
  列表组件： {{state}}
</template>

<script>
  import { reactive, watch, onMounted, nextTick } from 'vue'
  import { createDataList, createModel, lifecycle } from '@naturefw/ui-core'
  
  // 局部状态
  import { getListState } from '../controller/state-list.js'

  import _gridMeta from '../../json/list1.json'
  import _formMeta from '../../json/form.json'

</script>

<script setup>
  const props = defineProps({
    moduleId: Number,
    selection: Object,
    pageInfo: Object,
    active: Object
  })
 
  const state = getListState()
 
  // 获取表单控件需要的meta
  const gridMeta = reactive(_gridMeta)
  // 改为插槽
  // formMeta.itemMeta[101].controlType = 1
  
  // 根据 meta 创建表单的 model
  const _dataList = createDataList(_formMeta.itemMeta, props.pageInfo.count).reverse()

  // 绑定控件的列表
  const dataList = reactive([])
  // 监听页号，设置列表数据
  watch(() => props.pageInfo.pageIndex, (index) => {
    // console.log(index)
    dataList.length = 0
    const cpp = (props.pageInfo.pageIndex - 1) * props.pageInfo.pageSize
    for(let i=0; i< props.pageInfo.pageSize; i++) {
      dataList.push(_dataList[cpp + i])
    }
  })

  setTimeout(() => {
    for(let i = 0; i< props.pageInfo.pageSize; i++) {
      dataList.push(_dataList[i])
    }
  }, 1)

</script>