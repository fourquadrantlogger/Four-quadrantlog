<template>
  <!--表单-->
  <el-row :gutter="20">
    <el-col :span="23">
      <nf-form
        v-form-drag="formMeta"
        :model="model"
        v-bind="formMeta"
      >
      </nf-form>
    </el-col>

    <el-col :span="8">
    </el-col>
    <el-col :span="8">
      <el-button type="" @click="save">保存</el-button>
    </el-col>
  </el-row>
</template>

<script>
  import { reactive, watch, onMounted, nextTick } from 'vue'
  import { store } from '@naturefw/nf-state'
  import { ElCol, ElRow } from 'element-plus'
  // import { createModel, lifecycle } from '@naturefw/ui-core'
  // 列表局部状态
  import { getListState } from '../controller/state-list.js'
  // 表单状态
  import { regFormState } from '../controller/state-form.js'

  export default {
    name: 'form-log',
    components: {
      ElCol,
      ElRow
    },
    props: {
      moduleId: [Number, String],
      dataId: [Number, String],
      formMetaId: [Number, String],
      actionId: [Number, String],
      type: [Number, String],
      dialogInfo: [Object]
    },
    setup () {
      // 获取列表状态
      const listState = getListState()

      // 注册表单状态
      const model = regFormState()

      // 获取表单控件需要的meta
      const { formMeta } = store.meta[listState.moduleId]

      // 监听 dataID 的变化
      watch(() => listState.selection.dataId, (id) => {
        console.log('记录ID的变化：', id)
        console.log('props：', props)
      })

      // 添加或者修改
      const save = () => {
        // 添加
        model.addNew()
        // 更新列表
        listState.loadData(true)
      }

      return {
        model,
        formMeta,
        save
      }
    }
  }
</script>
