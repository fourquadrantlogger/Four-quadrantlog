<template>
  <el-table
    ref="gridDom"
    size="mini"
    style="width: 100%"
    :data="dataList"
    :height="height"
    :stripe="stripe"
    :border="border"
    :fit="fit"
    :highlight-current-row="highlightCurrentRow"
    :current-row-key="idName"
    :row-key="idName"
    @selection-change="selectionChange"
    @current-change="currentChange"
  >
    <!--显示选择框-->
    <el-table-column
      type="selection"
      width="55">
    </el-table-column>
    <!--显示字段列表-->
    <el-table-column
      v-for="(id, index) in colOrder"
      :key="'grid_list_' + index + '_' + id"
      :colId="id"
      :column-key="'col_' + id"
      :fixed="id < fixedIndex"
      v-bind="itemMeta[id]"
      :prop="itemMeta[id].colName"
      :min-width="50"
    >
    </el-table-column>
    <!--右面的操作列-->
  </el-table>
</template>

<script>
  import { onMounted, ref } from 'vue'
  import { ElTable, ElTableColumn } from 'element-plus'
  // 表单控件的属性 

/**
 * 表单控件需要的属性propsForm
 */
const gridProps = {
  moduleId: { // 101 gridID，必填
    type: Number,
    required: true
  },
  height: { // table的高度
    type: Number,
    default: 300
  },
  stripe: { // 斑马纹
    type: Boolean,
    default: true
  },
  border: { // 纵向边框
    type: Boolean,
    default: true
  },
  highlightCurrentRow: { // 要高亮当前行
    type: Boolean,
    default: true
  },
  colOrder: { // 字段显示的顺序
    type: Array,
    default: () => []
  },
  fixedIndex: { //  锁定的列数
    type: Number,
    default: 0
  },
  idName: { // 主键字段的名称
    type: String,
    default: 'ID'
  },
  itemMeta: Object, // table的列的 meta
  selection: { // 选择的情况
    type: Object,
    default: () => {
      return {
        dataId: '', // 单选ID number 、string
        row: {}, // 单选的数据对象 {}
        dataIds: [], // 多选ID []
        rows: [] // 多选的数据对象 []
      }
    }
  },
  'data-list': { // 绑定的数据
    type: Array,
    default: () => []
  }
}

export { gridProps }

/**
 * 列表的单选和多选的事件
 */
const choiceManage = (props, gridDom) => {
  // 是否单选触发
  let isCurrenting = false
  // 是否多选触发
  let isMoring = false

  // 单选
  const currentChange = (row) => {
    if (isMoring) return // 多选代码触发
    if (row === null) return // 清空

    if (gridDom.value) {
      isCurrenting = true
      gridDom.value.clearSelection() // 清空多选
      gridDom.value.toggleRowSelection(row) // 设置复选框
      gridDom.value.setCurrentRow(row) // 设置单选
      // 记录
      props.selection.dataId = row[props.idName]
      props.selection.dataIds[0] = row[props.idName]
      props.selection.row = row
      props.selection.rows = row

      isCurrenting = false
    }
    
  }

  // 多选
  const selectionChange = (rows) => {
    if (isCurrenting) return
    // 记录
    if (typeof props.selection.dataIds === 'undefined') {
      props.selection.dataIds = []
    }
    props.selection.dataIds.length = 0 // 清空
    // 设置多选
    rows.forEach((item) => {
      if (typeof item !== 'undefined' && item !== null) {
        props.selection.dataIds.push(item[props.idName])
      }
    })
    props.selection.rows = rows
    // 设置单选
    switch (rows.length) {
      case 0:
        // 清掉单选
        gridDom.value.setCurrentRow()
        props.selection.dataId = ''
        props.selection.row = {}
        break
      case 1:
        isMoring = true
        // 设置新单选
        gridDom.value.setCurrentRow(rows[0])
        isMoring = false
        props.selection.row = rows[0]
        props.selection.dataId = rows[0][props.idName]
        break
      default:
        // 清掉单选
        gridDom.value.setCurrentRow()
        props.selection.row = rows[rows.length - 1]
        props.selection.dataId = props.selection.row[props.idName]
    }
  }

  return {
    currentChange, // 单选
    selectionChange // 多选
  }
}

/**
 * * height: 250, // 高度
 * * colOrder: [101], // 显示顺序
 * * fixedIndex: 0 // 锁定的列数
 * * itemMeta: {
 * * --101: {
 * * ----id: 101,
 * * ----label: '编号11',
 * * ----prop: 'code',
 * * ----width: 120
 * * --}
 */
export default {
  name: 'nf-el-grid-list',
  inheritAttrs: false,
  components: {
    ElTable, ElTableColumn
  },
  props: {
    ...gridProps
  },
  setup (props, ctx) {
    // console.log('\ngrid---props', props) //

    // 获取 table 的 dom 和其他，
    const gridDom = ref(null)
    onMounted(() => {
    })

    // 列表选项的事件
    const {
      currentChange, // 单选
      selectionChange // 多选
    } = choiceManage(props, gridDom)

    // 格式化
    const myformatter = (row, column, cellValue, index) => {
      // console.log('\nrow:', row, 'col:', column) //
      // console.log('cellValue:', cellValue, 'index:', index) //
      if (Array.isArray(cellValue)) {
        return cellValue.join(',')
      }
      if (typeof cellValue === 'boolean') {
        return cellValue.toString()
      }
      return row[column.property] // realWidth
    }

    return {
      selectionChange, // 多选
      currentChange, // 单选
      gridDom, // table 的 dom
      myformatter // 格式化
    }
  }
}
</script>
