<template>
    <snote ref="inote"></snote>
     <el-button type="success" @click="createlog()">记录</el-button>
</template>
<script>
import { ElMessage } from 'element-plus'
import { createlog } from '../apis/apis'
import StickyNote from "../layouts/StickyNote.vue"
export default {
    data() {
        return
    },
    components: {
        'snote': StickyNote,
    },
     methods: {
        createlog() {
            
            if (this.$refs.inote.ctime === null) {
                ElMessage("时间必填")
                return
            }
            if ( this.$refs.inote.title ===  null) {
                ElMessage("标题必填")
                return
            }
            if ( this.$refs.inote.atype ===  null) {
                ElMessage("分类必填")
                return
            }
            if ( this.$refs.inote.quadrant ===  null) {
                ElMessage("象限必填")
                return
            }
            var l = {
                ctime:  this.$refs.inote.ctime,
                location:  this.$refs.inote.location,
                title:  this.$refs.inote.title,
                atype:  this.$refs.inote.atype,
                quadrant:  this.$refs.inote.quadrant.join('/'),
                detail:  this.$refs.inote.detail,
                review:  this.$refs.inote.review,
            }
            console.log(l)
            createlog(l).then(function (res) {
                if (res.data.code == 200) {
                    that.user = res.data.data;
                    console.log(res.data.data.username);
                    // console.log(that.user.password);
                }
                console.log(res.data.msg);
            }).catch(function (error) {
                console.log(error);
            });
        },
    },
}
</script>