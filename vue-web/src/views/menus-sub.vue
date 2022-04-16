<template>
  <template v-for="(item, index) in subMenu">
    <!--树枝-->
    <template v-if="item.children && item.children.length > 0">
      <el-sub-menu
        :key="item.menuId + '_' + index"
        :index="item.menuId"
        style="vertical-align: middle;"
      >
        <template #title>
          <component
            :is="$icon[item.icon]"
            style="width: 1.5em; height: 1.5em; margin-right: 8px;vertical-align: middle;"
          >
          </component>
          <span>{{item.text}}</span>
        </template>
        <!--递归子菜单-->
        <my-sub-menu2
          :subMenu="item.children"
        />
      </el-sub-menu>
    </template>
    <!--树叶-->
    <el-menu-item v-else
      :index="item.menuId"
      :key="item.menuId + 'son_' + index"
    >
      <template #title>
        <span style="float: left;">
          <component
            :is="$icon[item.icon]"
            style="width: 1.5em; height: 1.5em; margin-right: 8px;vertical-align: middle;"
          >
          </component>
          <span >{{item.text}}</span>
        </span>
      </template>
    </el-menu-item>
  </template>
</template>

<script>
  import { ElMenuItem, ElSubMenu } from 'element-plus'
  // 展示子菜单 - 递归
  import mySubMenu2 from './menus-sub.vue'

</script>

<script setup>

  const props = defineProps({
    subMenu: Array // 要显示的菜单，可以n级
  })

</script>
