<template>
<div style="display: flex " class="all">
   <CloudTag style="width:400px;height:400px;" :datas="q1" :title='肉体象限'></CloudTag>
  <CloudTag style="width:400px;height:400px;" :datas="q2" :title='知识信息象限'></CloudTag>
  <CloudTag style="width:400px;height:400px;" :datas="q3" :title='社会关系象限'></CloudTag>
  <CloudTag style="width:400px;height:500px;" :datas="q4" :title='持有物象限'></CloudTag>
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
      // if (this.ctimestartquery != null) {
      //     query.start = this.ctimestartquery.toLocaleString()
      // }
      // if (this.ctimeendquery != null) {
      //     query.end = this.ctimeendquery.toLocaleString()
      // }
      // if (this.quadrantquery != null) {
      //     if (this.quadrantquery instanceof Array) {
      //         query.quadrant = this.quadrantquery.join('/')
      //     } else {
      //         query.quadrant = this.quadrantquery
      //     }

      // }

      // if (this.locationquery != undefined) {
      //     query.location = this.locationquery
      // }
      // if (this.titlequery != undefined) {
      //     query.title = this.titlequery
      // }
      // if (this.detailquery != undefined) {
      //     query.detail = this.detailquery
      // } if (this.reviewquery != undefined) {
      //     query.review = this.reviewquery
      // }

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