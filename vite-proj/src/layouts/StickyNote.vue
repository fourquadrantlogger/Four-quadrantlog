<template>

    <el-input v-model="log.title" placeholder="title" />
    <el-input v-model="log.atype" placeholder="分类" />
    <el-input v-model="log.location" placeholder="location" />
    <div class="block">
        <el-date-picker
            v-model="log.ctime"
            type="datetime"
            placeholder="日期时间"
            align="right"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
        ></el-date-picker>
    </div>
    <div style="display: inline-block">
        <el-select v-model="log.quadrant" multiple placeholder="Select" style="width: 240px">
            <el-option v-for="item in options" :key="item" :label="item" :value="item" />
        </el-select>
    </div>

    <el-input v-model="log.detail" :rows="5" type="textarea" placeholder="内容" />
    <el-input v-model="log.review" :rows="5" type="textarea" placeholder="思考" />
</template>
<script>
import {Quadrant} from '../apis/apis'
export default {
    props: {
        log: {
            type: Object,
            // 对象或数组的默认值必须从一个工厂函数返回
            default() {
                return {
                    ctime: '',
                    location: '',
                    title: '',
                    atype: '',
                    quadrant: [''],
                    detail: '',
                    review: '',
                }
            }
        },
    },
    data() {
        return {
            options: Quadrant,

        }
    },

}
</script>