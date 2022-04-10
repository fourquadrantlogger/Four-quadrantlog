<template>
  <div
    style="position:fixed;width: 220px;left: 10px;top:100px; border:0px solid #123123;z-index: 100;"
  >
    <el-menu
      ref="domMenu"
      class="el-menu-vertical-demo"
      default-active="1"
      v-bind="$attrs"
      :background-color="backgroundColor"
      @select="select"
      @open="open"
      @close="close"
    >
      <!--递归调用，支持n级菜单 myMenu-->
      <mySubMenu :subMenu="menus"/>
    </el-menu>
  </div>
</template>

<script>
  // 模块菜单
  import { reactive, ref, watch } from 'vue'
  import { ElMenu } from 'element-plus'
  // 状态
  // import { store } from '@naturefw/nf-state'
  // 展示子菜单
  import mySubMenu from './menus-sub.vue'

</script>

<script setup>
 const props = defineProps({
    'background-color': {
      tyoe: String,
      default: '#ece5d9'
    }
  })

  // 当前菜单、文档内容，菜单列表
  // const { current, menuList } = store
  
  // 自动展开菜单用
  const domMenu = ref(null)
 
  const menus = [
    {
      "menuId": "1000",
      "text": "时空",
      "description": "描述",
      "icon": "FolderOpened",
      "children": [
        {
          "menuId": "11_quick_start_1",
          "text": "选项一",
          "description": "",
          "icon": "Document",
          "children": []
        },
        {
          "menuId": "11_quick_start_2",
          "text": "选项二",
          "description": "",
          "icon": "Document"
        }
      ]
    },
    {
      "menuId": "2000",
      "text": "肉体",
      "description": "描述",
      "icon": "FolderOpened" 
    },
    {
      "menuId": "3000",
      "text": "信息知识",
      "description": "描述",
      "icon": "FolderOpened",
      "children": [
        {
          "menuId": "11_quick_start_1",
          "text": "选项一",
          "description": "",
          "icon": "Document",
          "children": []
        },
        {
          "menuId": "11_quick_start_2",
          "text": "选项二",
          "description": "",
          "icon": "Document"
        }
      ]
    },
    {
      "menuId": "4000",
      "text": "社会关系",
      "description": "描述",
      "icon": "FolderOpened",
      "children": [
        {
          "menuId": "11_quick_start_1",
          "text": "选项一",
          "description": "",
          "icon": "Document",
          "children": []
        },
        {
          "menuId": "11_quick_start_2",
          "text": "选项二",
          "description": "",
          "icon": "Document"
        }
      ]
    },
    {
      "menuId": "5000",
      "text": "持有物",
      "description": "描述",
      "icon": "FolderOpened",
      "children": [
        {
          "menuId": "11_quick_start_1",
          "text": "选项一",
          "description": "",
          "icon": "Document",
          "children": []
        },
        {
          "menuId": "11_quick_start_2",
          "text": "选项二",
          "description": "",
          "icon": "Document"
        }
      ]
    }
  ]

  
  /**
   * 二级菜单被选中
   * @param {*} index 
   * @param {*} indexPath 
   */
  const select = (index, indexPath) => {

  }

</script>

<style>
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    max-width: 320px;
    min-height: 400px;
  }
</style>
