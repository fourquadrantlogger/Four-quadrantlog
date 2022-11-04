<template>
    <el-row style="height: 50px">
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
      <el-input v-model="atypequery" placeholder="类别" />
    </el-col>

  </el-row>

<div style="display: flex " class="all">
   <CloudTag style="width:800px;height:800px;" :datas="q1" :title='肉体象限'></CloudTag>
  <CloudTag style="width:800px;height:800px;" :datas="q2" :title='知识信息象限'></CloudTag>
  <CloudTag style="width:800px;height:800px;" :datas="q3" :title='社会关系象限'></CloudTag>
  <CloudTag style="width:800px;height:800px;" :datas="q4" :title='持有物象限'></CloudTag>
</div>

</template>
<style lang="css">
</style>
<script>
import CloudTag from '../layouts/CloudTag.vue'
import { getTagList, Quadrant } from '../apis/apis'

import { ElMessage } from 'element-plus'

export default {

  components: { CloudTag },
  data: function () {
    return {
      q1: [],
      q2: [],
      q3: [],
      q4: [],
      ctimestartquery: "",
      ctimeendquery: "",
      atypequery:"",
      limit: 100,
    }
  },
  mounted: function () {
    this.tagLog()
  },
  methods: {
    async tagLog() {
      let query = {}
      query.limit = 50
      let start = new Date()
      start.setMonth(start.getMonth() - 3);
      query.start = start.toLocaleString()
      if (this.ctimestartquery != null) {
          query.start = this.ctimestartquery.toLocaleString()
      }
      if (this.ctimeendquery != null) {
          query.end = this.ctimeendquery.toLocaleString()
      }

      if (this.ctimeendquery != null) {
          query.atype = this.atypequery
      }
      {
        query.quadrant = '肉体'
        const taglist = await getTagList(query)
        this.q1 = [].concat(taglist)
      }
      {
        query.quadrant = '信息'
        const taglist = await getTagList(query)
        this.q2 = [].concat(taglist)
      }
      {
        query.quadrant = '社会关系'
        const taglist = await getTagList(query)
        this.q3 = [].concat(taglist)
      }
      {
        query.quadrant = '持有物'
        const taglist = await getTagList(query)
        this.q4 = [].concat(taglist)
      }

      ElMessage.success('获取tags成功')
      return
    },
  },
}
</script>