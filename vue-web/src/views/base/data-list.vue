<template>
 <!--查询控件 延迟加载-->
  <div style="height: 65px;">
    <component
      :is="delay[delayKind.find]"
      :moduleId="moduleId"
    >
    </component>
  </div>
  <!--操作按钮 延迟加载-->
  <div style="height: 30px;">
    <component
      :is="delay[delayKind.button]"
      :selection="selection"
      :moduleId="moduleId"
    >
    </component>
  </div>
  <!--列表控件-->
  <my-grid
    :moduleId="moduleId"
    :selection="selection"
    :pageInfo="pagerInfo"
  >
  </my-grid>
  <!--分页-->
  <my-pager
    :pageInfo="pagerInfo"
  >
  </my-pager>
  列表页面：{{state}}
</template>

<script>
/**
 * 数据列表的基础页面
 */
  import { reactive, watch, defineAsyncComponent } from 'vue'
  // import { ElPagination, ElInput } from 'element-plus'
  import { createDataList, mykeypager, lifecycle } from '@naturefw/ui-core'
  // import { loading } from '/nf-ui-elp'
  import _gridMeta from '../json/list1.json'
  import _formMeta from '../json/form.json'

  import myPager from './components/pager.vue'
  import myGrid from './components/grid.vue'

  // 局部状态
  import { getListState,
  regListState } from './controller/state-list.js'

</script>

<script setup>

  const state = regListState()
  state.moduleId = 100
  
  // 分页信息
  const pagerInfo = reactive({
    pageSize: 10,
    count: 20, // 总数
    pageIndex: 1 // 当前页号
  })

  const moduleId = new Date().valueOf()

  // 列表里选择的记录
  const selection = reactive({
    dataId: '', // 单选ID number 、string
    row: {}, // 单选的数据对象 {}
    dataIds: [], // 多选ID []
    rows: [] // 多选的数据对象 []
  }) 

  // 当前激活的模块
  // const { active } = useActive(moduleId)

  // 延迟加载
  const delay = {
    load: '',
    find: defineAsyncComponent(() => import('./components/find.vue')),
    button: defineAsyncComponent(() => import('./components/button.vue'))
  }
  const delayKind = reactive({
    find: 'load',
    button: 'load'
  })

  setTimeout(() => {
    delayKind.find = 'find'
    delayKind.button = 'button'
  }, 1)
 

</script>

