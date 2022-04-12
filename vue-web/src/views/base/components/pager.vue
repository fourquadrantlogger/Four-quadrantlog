<template>
  <el-pagination
    background
    layout="prev, pager, next"
    v-model:currentPage="state.pagerInfo.pagerIndex"
    :page-size="state.pagerInfo.pagerSize"
    :total="state.pagerInfo.count"
  >
  </el-pagination>
</template>

<script>
/* eslint-disable vue/no-v-model-argument */
  import { ElPagination } from 'element-plus'
  // 局部状态
  import { getListState } from '../controller/state-list.js'

</script>

<script setup>
  const state = getListState()
  const props = defineProps({
  })

</script>
