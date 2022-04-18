<template>
    <el-row style="height:50px;">
        <el-col :span="3">
            <div class="grid-content bg-purple" />
            <el-date-picker v-model="ctimestartquery" type="datetime" placeholder="起始" align="right"
                value-format="YYYY-MM-DD HH:mm:ss" format="YYYY-MM-DD HH:mm:ss"></el-date-picker>
        </el-col>
        <el-col :span="3">
            <div class="grid-content bg-purple-light" />
            <el-date-picker v-model="ctimeendquery" type="datetime" placeholder="结束" align="right"
                value-format="YYYY-MM-DD HH:mm:ss" format="YYYY-MM-DD HH:mm:ss" />
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

    <el-row>
        <el-scrollbar :style="maintable">
            <el-table class="table-content" @cell-dblclick="celldbclick" :data="List" border show-summary>
                <el-table-column prop="quadrant" label="象限" width="60px">
                    <template #footer="{ scope }">
                        <slot name="quadrant_slot" :scope="state">
                            <el-input v-if="scope.row[scope.column.property + 'isShow']" :ref="scope.column.property"
                                v-model="scope.row.quadrant" @blur="alterData(scope.row, scope.column)"></el-input>
                            <span v-else>{{ scope.row.quadrant }}</span>
                        </slot>
                    </template>
                </el-table-column>>
                <el-table-column prop="ctime" label="时间" width="110px"> </el-table-column>
                <el-table-column prop="location" label="地址" width="200px" />
                <el-table-column prop="atype" label="类别" width="120px" />
                <el-table-column prop="title" label="标题" width="200px" />
                <el-table-column prop="detail" label="详情" width="calc(100% - 790px)" />
                <el-table-column prop="review" label="回顾" width="100px" />
            </el-table>
        </el-scrollbar>
    </el-row>

    <el-row :gutter="20">
        <el-col :span="4">
            <div class="grid-content bg-purple" />
        </el-col>
        <el-col :span="16">
            <div class="grid-content bg-purple" />

            <div class="grid-content bg-purple">
                <el-pagination v-model:currentPage="currentpage" background layout="prev, pager, next" :total="total"
                    style="width:100%;" />

            </div>
        </el-col>
        <el-col :span="4">
            <div class="grid-content bg-purple" />
        </el-col>


    </el-row>
</template>
<style  >
</style>
<script >
import { getLogList, Quadrant } from '../apis/apis'
import { ElMessage } from 'element-plus'
import {GetZhWeekDay} from '../xutil/xtime'
export default {
    created() {
        this.resetheight()
        window.addEventListener('resize', this.resetheight);
    },
    destroyed() {
        window.removeEventListener('resize', this.resetheight)
    },
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
            currentpage: 1,
            pagesize: 20,
            List: [],
            QuadrantOptions: Quadrant,
            maintable: {
                height: '600px',
                width: '100%',
            },
            total: 0,
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
        currentpage: function (v) {
            this.listLog()
        },
        pagesize: function (v) {
            this.listLog()
        },
    },
    mounted: function () {
        console.log(this.$route.query)
        this.atypequery = this.$route.query.atypequery;
        this.listLog();
    },
    methods: {

        celldbclick(row, column, cell, event) {

            console.log(row, column, cell, event);
            
            let ext=row.title.substring(row.title.lastIndexOf("."))
            if (ext != '') {
               window.open('/api/blob/'+row.id,'_blank');
            } else {
                this.$router.push(
                    {
                        path: "/note/" + row.id,
                    }
                )
            }
        },
        resetheight() {
            this.maintable.height = window.innerHeight - 150 + 'px'
        },
        async listLog() {
            let query = {}
            if (this.ctimestartquery != null) {
                query.start = this.ctimestartquery.toLocaleString()
            }
            if (this.ctimeendquery != null) {
                query.end = this.ctimeendquery.toLocaleString()
            }
            if (this.quadrantquery != null) {
                if (this.quadrantquery instanceof Array) {
                    query.quadrant = this.quadrantquery.join('/')
                } else {
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

            query.offset = (this.currentpage - 1) * this.pagesize
            query.limit = this.pagesize

            const loglist = await getLogList(query)
            console.log(loglist)
            this.List = [].concat(loglist.data)
            for (let i=0;i<this.List.length;i++) {
                let weekday = GetZhWeekDay(new Date(this.List[i].ctime));
                this.List[i].ctime=["星期"+weekday,this.List[i].ctime];
            }
            this.total = loglist.total
            ElMessage.success('获取日志成功')
            return
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
