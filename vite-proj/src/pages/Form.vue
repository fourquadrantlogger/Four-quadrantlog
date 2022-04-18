<template>
    <h2>{{ logid }}</h2>
    <snote ref="inote" :log=log></snote>
     <input type="file"  ref="uploadFiles" placeholder="选择文件">
    <el-button type="success" @click="createblob()">记录</el-button>
    <el-input v-model="blobrename" placeholder="重命名为" clearable />
</template>
<script>
import { ElMessage } from 'element-plus'
import { createblob, updatelog, getlog } from '../apis/apis'
import StickyNote from "../layouts/StickyNote.vue"
 
export default {
    data() {
        return {
            log: {},
            blobrename:undefined,
        }
    },
    props: {
        logid: undefined,
    },
    components: {
        'snote': StickyNote,
    },

    mounted: function () {
        this.log = {};
        if (this.logid != 0 && this.logid != undefined) {
            this.showblob()
        }
    },

    methods: {
        async showblob() {
            let l = await getlog(this.logid)
            this.log = l
            l.quadrant = l.quadrant.split('/')
            console.log(l)
        },
        async createblob() {
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
            var formData = new FormData();
            formData.append('title', this.$refs.inote.log.title);
            formData.append('atype', this.$refs.inote.log.atype);
            formData.append('location', this.$refs.inote.log.location);
            formData.append('ctime', this.$refs.inote.log.ctime);

            formData.append('detail', this.$refs.inote.log.detail);
            formData.append('review', this.$refs.inote.log.review);
             if (this.blobrename != undefined) {
                 formData.append('blobrename', this.blobrename);
            }
            
            if (this.$refs.uploadFiles.files.length == 0) {
                ElMessage.warning("请选择文件")
                return
            }

            let blob=this.$refs.uploadFiles.files[0]

            formData.append('blob', blob);


            if (this.$data.log.quadrant instanceof Array) {
                let quadrant = this.$data.log.quadrant.join('/')
                formData.append('quadrant', quadrant);
            } else {
                formData.append('quadrant', quadrant);
            }

            if (this.logid === null || this.logid == undefined) {
                const clog = await createblob(formData)
                ElMessage.success("创建成功")
                this.$router.push('/note/' + clog.data.id)
            }  
        },
    },
}
</script>
