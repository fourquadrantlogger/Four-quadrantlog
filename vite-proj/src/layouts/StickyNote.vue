<template>
    <el-input v-model="title" placeholder="title" />
    <el-input v-model="atype" placeholder="分类" />
    <el-input v-model="location" placeholder="location" />
    <div class="block">
        <span class="demonstration">带快捷选项</span>
        <el-date-picker
            v-model="ctime"
            type="datetime"
            placeholder="选择日期时间"
            align="right"
            format="YYYY-MM-DD HH:mm:ss"
            :picker-options="pickerOptions"
        ></el-date-picker>
    </div>
    <div style="display: inline-block">
        <el-select v-model="quadrant" multiple placeholder="Select" style="width: 240px">
            <el-option v-for="item in options" :key="item" :label="item" :value="item" />
        </el-select>
    </div>

    <el-input v-model="detail" :rows="5" type="textarea" placeholder="内容" />
    <el-input v-model="review" :rows="5" type="textarea" placeholder="思考" />
   
</template>
<script>

export default {
    props: {
        ctime: {
            type: Date,
            required: true
        },
        location: {
            type: String,
            required: true
        },
        title: {
            type: String,
            required: true
        },
        atype: {
            type: String,
            required: true
        },
        quadrant:{
            type: Array,
            required: true
        },
        detail: '',
        review: '',
    },
    data() {
        return {
            options: [
                '肉体',
                '信息',
                '知识',
                '时空',
                '社会关系',
                '持有物'
            ],
            pickerOptions: {
                shortcuts: [{
                    text: '今天',
                    onClick(picker) {
                        picker.$emit('pick', new Date());
                    }
                }, {
                    text: '昨天',
                    onClick(picker) {
                        const date = new Date();
                        date.setTime(date.getTime() - 3600 * 1000 * 24);
                        picker.$emit('pick', date);
                    }
                }, {
                    text: '一周前',
                    onClick(picker) {
                        const date = new Date();
                        date.setTime(date.getTime() - 3600 * 1000 * 24 * 7);
                        picker.$emit('pick', date);
                    }
                }]
            },

        }
    },
  
}
</script>