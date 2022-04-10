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
  </el-row>
   <el-row :gutter="20">
    <el-col :span="8">
    </el-col>
    <el-col :span="8">
      <el-button type="" @click="save">保存</el-button>
    </el-col>
  </el-row>
</template>


<script>
  import { reactive, watch, onMounted, nextTick } from 'vue'
  import { ElCol, ElRow, ElSlider } from 'element-plus'
  import { createModel, lifecycle } from '@naturefw/ui-core'
  import _formMeta from '../../json/form.json'

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

  // 获取表单控件需要的meta
  const formMeta = reactive(_formMeta)
  // 改为插槽
  // formMeta.itemMeta[101].controlType = 1

  // 根据 meta 创建表单的 model
  const model = createModel(formMeta.itemMeta, formMeta.colOrder)

  const addNew = () => {
    const a = {}
    Object.assign(a, model)
    dataList.unshift(a)
  }

  const save = () => {

  }

</script>