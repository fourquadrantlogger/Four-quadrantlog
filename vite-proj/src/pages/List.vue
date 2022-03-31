<template>
 
    <p>log list</p>
    <template v-for="item in List" :key="item.msg">
        <li><p>{{ item.ctime }} {{ item.quadrant }} {{ item.atype }} </p> {{ item.title }}</li>
        <li class="divider" role="presentation"></li>
    </template>

</template>

<script>  
import getLogList from '../apis/apis'

import { computed } from '@vue/runtime-core'
export default {
    data(){
        return {
             List:[],
        }
    },
    mounted:function(){
        this.listLog();
    },
    methods:{
        async listLog(){
            const loglist= await getLogList()
            console.log(loglist)
            this.List=[].concat(loglist)
            return 
        },
    },
    computed: {
         
    },
}
 
</script>