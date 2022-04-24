<template>
<el-form :model="form" label-width="120px">
    <el-form-item label="Activity name">
      <el-input v-model="form.name" />
    </el-form-item>
</template>
</template>
<style  >
</style>
<script >
import { getLogList, Quadrant } from '../apis/apis'
import { ElMessage } from 'element-plus'

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
            form:{
 ctime: '',
            quadrant: '',
            atype: '',
            location: '',
            title: '',
            detail: '',
            review: '',
            blob:null,
            },
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
        currentpage: function (v) {
            this.listLog()
        },
        pagesize: function (v) {
            this.listLog()
        },
    },
    mounted: function () {
        this.listLog();
    },
    methods: {
        celldbclick(row, column, cell, event){

            console.log(row, column, cell, event);

            if (column.property=='blob'){

            }else{
                this.$router.push(
                    {
                        path:"/note/"+row.id,
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
            this.List = [].concat(loglist)
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
