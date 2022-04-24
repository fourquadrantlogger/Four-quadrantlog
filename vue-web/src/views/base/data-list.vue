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
      :moduleId="moduleId"
    >
    </component>
  </div>
  <!--列表控件-->
  <my-grid></my-grid>
  <!--分页-->
  <my-pager></my-pager>
  列表数据的状态：
  <div v-for="(item, key) in state">
    {{key}}： 
    <div style="padding-left: 20px;"
      v-if="key === 'selection'">
      <div v-for="(item2, key2) in item">
        {{key2}}： {{item2}}
      </div>
    </div>
    <span v-else>
      {{item}}
    </span>
  </div>
</template>

<script>
/**
 * 数据列表的基础页面
 */
  import { reactive, watch, defineAsyncComponent } from 'vue'
  // import { ElPagination, ElInput } from 'element-plus'

  // 局部状态，注册列表数据的状态
  import { regListState } from './controller/state-list.js'

  // 分页组件
  import myPager from './components/pager.vue'
  // 列表组件
  import myGrid from './components/grid.vue'
  // 查询组件 -- 延迟加载
  const _find = defineAsyncComponent(() => import('./components/find.vue'))
  // 按钮组件 -- 延迟加载
  const _button = defineAsyncComponent(() => import('./components/button.vue'))

</script>

<script setup>
  // 注册状态
  const state = regListState()
  state.moduleId = 'm_log'

  // 当前激活的模块
  // const { active } = useActive(moduleId)

  // 延迟加载
  const delay = {
    load: '',
    find: _find,
    button: _button
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
