<template>
    <div class="demo-date-picker">
        <div class="block">
            <span class="demonstration">Default</span>
            <el-date-picker
                v-model="ctimestart"
                type="daterange"
                range-separator="To"
                start-placeholder="Start date"
                end-placeholder="End date"
            />
        </div>
        <div class="block">
            <span class="demonstration">With quick options</span>
            <el-date-picker
                v-model="ctimeend"
                type="daterange"
                unlink-panels
                range-separator="To"
                start-placeholder="Start date"
                end-placeholder="End date"
            />
        </div>
    </div>
    <el-input v-model="quadrantquery" placeholder="象限" />
    <el-input v-model="atypequery" placeholder="类别" />
    <el-input v-model="titlequery" placeholder="标题" />
    <el-input v-model="detailquery" placeholder="详情" />
    <el-input v-model="reviewquery" placeholder="回顾" />

    <el-table :data="List" height="1000px" style="width: 100%">
        <el-table-column prop="quadrant" label="象限" width="180" />
        <el-table-column prop="ctime" label="时间" width="180" />
        <el-table-column prop="location" label="地址" />
        <el-table-column prop="atype" label="类别" />
        <el-table-column prop="title" label="标题" width="180" />
        <el-table-column prop="detail" label="详情" />
        <el-table-column prop="review" label="回顾" />
    </el-table>
</template>

<script >
import {getLogList} from '../apis/apis'
import { ElMessage } from 'element-plus'

export default {
    data() {
        return {
            ctimestart: '',
            ctimeend: '',
            quadrantquery: '',
            atypequery: '',
            locationquery: '',
            titlequery: '',
            detailquery: '',
            reviewquery: '',
            List: [],
        }
    },
    watch: {
        quadrantquery: function (v) {
            this.listLog()
        },
        atypequery: function (v) {
            this.listLog()
        },
        locationquery: function (v) {
            this.listLog()
        },
        titlequery: function (v) {
            this.listLog()
        },
        detailquery: function (v) {
            this.listLog()
        },
        reviewquery: function (v) {
            this.listLog()
        },
    },
    mounted: function () {
        this.listLog();
    },
    methods: {
        async listLog() {
            const loglist = await getLogList({

                quadrant: this.quadrantquery,
                atype: this.atypequery,
                location: this.locationquery,
                title: this.titlequery,
                detail: this.detailquery,
                review: this.reviewquery,
            })
            console.log(loglist)
            this.List = [].concat(loglist)
            ElMessage.success('获取日志成功')
            return
        },
    },
}

</script>
