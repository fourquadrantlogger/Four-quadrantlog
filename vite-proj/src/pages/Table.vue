<template>
    <el-row style="height:50px;">
        <el-col :span="3">
            <div class="grid-content bg-purple" />
            <el-date-picker
                v-model="ctimestartquery"
                type="datetime"
                placeholder="起始"
                align="right"
                value-format="YYYY-MM-DD HH:mm:ss"
                format="YYYY-MM-DD HH:mm:ss"
            ></el-date-picker>
        </el-col>
        <el-col :span="3">
            <div class="grid-content bg-purple-light" />
            <el-date-picker
                v-model="ctimeendquery"
                type="datetime"
                placeholder="结束"
                align="right"
                value-format="YYYY-MM-DD HH:mm:ss"
                format="YYYY-MM-DD HH:mm:ss"
            />
        </el-col>
        <el-col :span="3">
            <el-select v-model="quadrantquery" multiple placeholder="象限">
                <el-option v-for="item in QuadrantOptions" :key="item" :label="item" :value="item" />
            </el-select>
        </el-col>

        <el-col :span="3">
            <el-input v-model="atypequery" placeholder="类别" />
        </el-col>

        <el-col :span="4">
            <el-input v-model="titlequery" placeholder="标题" />
        </el-col>

        <el-col :span="4">
            <el-input v-model="detailquery" placeholder="详情" />
        </el-col>

        <el-col :span="4">
            <el-input v-model="reviewquery" placeholder="回顾" />
        </el-col>
    </el-row>

    <el-table
        :data="List"
        border
        show-summary
        :key="randomKey"
        @cell-dblclick="editData"
        class="maintable"
    >
        <el-table-column prop="quadrant" label="象限" width="100px" solt="size">
            <template #footer="{ scope }">
                <slot name="quadrant_slot" :scope="state">
                    <el-input
                        v-if="scope.row[scope.column.property + 'isShow']"
                        :ref="scope.column.property"
                        v-model="scope.row.quadrant"
                        @blur="alterData(scope.row, scope.column)"
                    ></el-input>
                    <span v-else>{{ scope.row.quadrant }}</span>
                </slot>
            </template>
        </el-table-column>>
        <el-table-column prop="ctime" label="时间" width="180" />
        <el-table-column prop="location" label="地址" />
        <el-table-column prop="atype" label="类别" />
        <el-table-column prop="title" label="标题" width="180" />
        <el-table-column prop="detail" label="详情" />
        <el-table-column prop="review" label="回顾" />
    </el-table>
    <el-pagination  v-model:currentPage="currentpage" background layout="prev, pager, next" :total="1000" @current-change="pagequery"/>
</template>
<style >
.maintable{
    height: calc(100% - 100px);
    width: auto;
}
</style>
<script >
import { getLogList, Quadrant } from '../apis/apis'
import { ElMessage } from 'element-plus'
 
export default {
    data() {
        return {
            randomKey: Math.random(),
            ctimestartquery: '',
            ctimeendquery: '',
            quadrantquery: '',
            atypequery: '',
            locationquery: '',
            titlequery: '',
            detailquery: '',
            reviewquery: '',
            currentpage:1,
            pagesize:20,
            List: [],
            QuadrantOptions: Quadrant,
        }
    },
    watch: {
        ctimeendquery: function (v) {
            this.listLog()
        },
        ctimestartquery: function (v) {
            this.listLog()
        },
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
        pagequery: function (v) {
            this.listLog()
        },
        pagesizequery: function (v) {
            this.listLog()
        },
    },
    mounted: function () {
        this.listLog();
    },
    methods: {
        async listLog() {
            let query = {}
            if (this.ctimestartquery != null) {
                query.start = this.ctimestartquery.toLocaleString()
            }
            if (this.ctimeendquery != null) {
                query.end = this.ctimeendquery.toLocaleString()
            }
            if (this.quadrantquery != null) {
                if (  this.quadrantquery instanceof Array){
                    query.quadrant = this.quadrantquery.join('/')
                }else{
                     query.quadrant = this.quadrantquery
                }
               
            }
            if (this.atypequery != undefined) {
                query.atype = this.atypequery
            }
            if (this.locationquery != undefined) {
                query.location = this.locationquery
            }
            if (this.titlequery != undefined) {
                query.title = this.titlequery
            }
            if (this.detailquery != undefined) {
                query.detail = this.detailquery
            } if (this.reviewquery != undefined) {
                query.review = this.reviewquery
            }
            
            query.offset=(this.currentpage-1)*this.pagesize
            query.limit=this.pagesize
            
            const loglist = await getLogList(query)
            console.log(loglist)
            this.List = [].concat(loglist)
            ElMessage.success('获取日志成功')
            return
        },
        editData(row, column) {
            console.log("......")
            row[column.property + "isShow"] = true
            //refreshTable是table数据改动时，刷新table的
            this.refreshTable()
            this.$nextTick(() => {
                this.$refs[column.property] && this.$refs[column.property].focus()
            })
        },
        alterData(row, column) {
            row[column.property + "isShow"] = false
            this.refreshTable()
        },

        refreshTable() {
            this.randomKey = Math.random()
        },

    },
}

</script>
