<template>
  <el-pagination
    background
    layout="prev, pager, next"
    v-model:currentPage="state.pagerInfo.pagerIndex"
    :total="state.pagerInfo.count"
    :active="active"
  >
  </el-pagination>
</template>


<script>
  import { ElPagination } from 'element-plus'
  // 局部状态
  import { getListState } from '../controller/state-list.js'

</script>

<script setup>
  const state = getListState()
 
  const props = defineProps({
    active: Object
  })

</script>