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
  import { ElCol, ElRow, ElSlider } from 'element-plus'
  // import { createModel, lifecycle } from '@naturefw/ui-core'
  // 局部状态
  import { getListState } from '../controller/state-list.js'
  import { regFormState } from '../controller/state-form.js'

</script>

<script setup>
  const props = defineProps({
    moduleId: [Number, String],
    dataId: [Number, String],
    formMetaId: [Number, String],
    actionId: [Number, String],
    type: [Number, String],
    dialogInfo: [Object]
  })

  // 获取列表状态
  const state = getListState()

  // 注册表单状态
  const modelManage = regFormState()
  const model = modelManage.model

  // 获取表单控件需要的meta
  const { formMeta } = store.meta[state.moduleId]

  // 添加或者修改
  const save = () => {
    modelManage.submit()
  }

</script>
