<template>
    <el-container class="alldiv" style="border: 1px solid #eee">
        <el-aside class="alldiv" width="200px">
            <div style=" display:flex;
    justify-content: center;width:100%;">
                <el-radio-group v-model="isCollapse" style="margin-bottom: 20px;">
                    <el-radio-button :label="false">+</el-radio-button>
                    <el-radio-button :label="true">-</el-radio-button>
                </el-radio-group>
            </div>

            <el-menu default-active="1-3" class="el-menu-vertical-demo" :collapse="isCollapse" @open="handleOpen"
                @close="handleClose">
                <el-sub-menu index="1">
                    <template #title>
                        <el-icon>
                            <location />
                        </el-icon>
                        <span>时空</span>
                    </template>
                    <el-menu-item-group>
                        <template #title><span>创建</span></template>
                        <el-menu-item index="1-1" @click="menuclick('/note')">文本</el-menu-item>
                         <el-menu-item index="1-2" @click="menuclick('/form')">文件</el-menu-item>
                    </el-menu-item-group>
                    <el-menu-item-group title="查询">
                         <el-menu-item index="1-2" @click="menuclick('/')">词频</el-menu-item>
                        <el-menu-item index="1-3" @click="menuclick('/table')">表格</el-menu-item>
                    </el-menu-item-group>
                </el-sub-menu>
                <el-menu-item index="2">
                    <el-icon>
                        <icon-menu />
                    </el-icon>
                    <template #title>身体</template>
                </el-menu-item>
                <el-menu-item index="3" disabled>
                    <el-icon>
                        <document />
                    </el-icon>
                    <template #title>信息知识</template>
                </el-menu-item>
                <el-menu-item index="4">
                    <el-icon>
                        <setting />
                    </el-icon>
                    <template #title>社会关系</template>
                </el-menu-item>
            </el-menu>
        </el-aside>

        <el-main style="height:100%" class="alldiv">
            <Layout style="width:100%,height:100%" class="alldiv">
                <router-view class="alldiv">
                </router-view>
            </Layout>
        </el-main>
    </el-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
} from '@element-plus/icons-vue'

import { useRouter } from "vue-router";


    const router = useRouter();
const menuclick = (path: string) => {
     console.log(path)
   router.push(path)

}
const isCollapse = ref(false)
const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
</script>

<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
}

html,
body,
#app,
.alldiv {
    height: 100%;
    margin: 0;
    padding: 0;
}

.el-header {
    background-color: #b3c0d1;
    color: var(--el-text-color-primary);
    line-height: 60px;
}

.el-aside {
    color: var(--el-text-color-primary);
}
</style>
