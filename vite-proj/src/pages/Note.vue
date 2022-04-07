<template>
    <h2>{{logid}}</h2>
    <snote ref="inote" :log=log></snote>
    <el-button type="success" @click="createlog()">记录</el-button>
</template>
<script>
import { ElMessage } from 'element-plus'
import { createlog,updatelog,getlog} from '../apis/apis'
import StickyNote from "../layouts/StickyNote.vue"

export default {
    data() {
        return {
         log:{},
        }
    },
    props:{
         logid:undefined,
    },
    components: {
        'snote': StickyNote,
    },
   
    mounted:function () {
 
        if (this.logid!=0&&this.logid!=undefined){
             this.showlog()  
        }
    },
       
    methods: {
        async showlog(){
            let l=await getlog(this.logid)
            this.log=l
            l.quadrant=l.quadrant.split('/')
            console.log(l)
        },
        async createlog() {
            if (this.$refs.inote.log.ctime === null) {
                ElMessage.error("时间必填")
                return
            }
            if (this.$refs.inote.log.title === null) {
                ElMessage.error("标题必填")
                return
            }
            if (this.$refs.inote.log.atype === null) {
                ElMessage.error("分类必填")
                return
            }
            if (this.$refs.inote.log.quadrant === null) {
                ElMessage.error("象限必填")
                return
            }
            var l =   this.$refs.inote.log 
            if ( l.quadrant   instanceof Array){
                 l.quadrant=l.quadrant.join('/')
            }
           
            if (this.logid === null||this.logid==undefined) {
                const clog =await createlog(l)
                ElMessage.success("创建成功")
                this.$router.push('/note/'+clog.data.id)
            } else {
                var ll=l
                ll.id=this.logid
                await updatelog(ll)
                ElMessage.success("更新成功")
            }

        },
    },
}
</script>
