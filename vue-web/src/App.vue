<template>
  <el-container>
    <el-header>
      <!--导航-->
      <div style="float: left;">
        <!--写网站logo、标题等-->
        <h1>四象限——人生日志管理系统1</h1>
      </div>
      <div style="float: right;min-width: 100px;height: 40px;padding-top: 3px;">
        <!--写网站logo、标题等-->
      </div>
      <div style="float: right;min-width: 600px;height: 60px;">
        <!--网站导航-->
      </div>
    </el-header>
    <el-container>
      <!--左侧边栏-->
      <el-aside width="230px">
        <!--菜单-->
        <menus></menus>
      </el-aside>
      <el-main>
        <!--编辑区域-->
        <component
          :is="delay[delayKind]"
        >
        </component>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
  import { defineAsyncComponent, watch, ref } from 'vue'
  import { ElContainer, ElHeader, ElAside, ElMain, ElSwitch } from 'element-plus'
  // 访问状态
  import { store } from '@naturefw/nf-state'
  // 左侧菜单导航
  import menus from './views/menus.vue'

  export default {
    components: {
      menus,
      ElContainer,
      ElHeader,
      ElAside,
      ElMain,
      ElSwitch
    },
    setup () {
      // 延迟加载
      const delay = {
        load: '',
        list: defineAsyncComponent(() => import('./views/base/data-list.vue'))
      }

      const delayKind = ref('')

      const { web } = store

      watch(() => store.web.load, (load) => {
        if (load === 'end') {
          delayKind.value = 'list'
        }
      })

      return {
        delay,
        delayKind
      }
    }
  }
</script>

<style>
body {
  background-color:#ece5d9;
}
  .el-header {
    background-color: #ece5d9;
    color: var(--el-text-color-primary);
    text-align: left;
    height: 80px;
    line-height: 30px;
    border: 1px solid #ccc!important;padding:14px;border-radius: 16px!important;
  }

  .el-aside {
    background-color: #ece5d9;
    color: var(--el-text-color-primary);
    text-align: left;
    line-height: 30px;
    min-height: 500px;
    border:1px solid rgba(159, 196, 238, 0.315);
  }

  .el-main {
    background-color: #ece5d9;
    color: var(--el-text-color-primary);
    text-align: left;
    line-height: 18px;
  }

  body > .el-container {
    margin-bottom: 10px;
  }

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin-top: 5px;
}
</style>
