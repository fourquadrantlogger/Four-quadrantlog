<template>
    <div id="mywordcloud" :style="{ width: '100%', height: '300px' }" :data="worddata"></div>

</template>
<script lang="javascript">

import * as echarts from "echarts";
import "echarts-wordcloud/dist/echarts-wordcloud";
import "echarts-wordcloud/dist/echarts-wordcloud.min";

import { getTagList, Quadrant } from '../apis/apis'

import { ElMessage } from 'element-plus'

export default {
    name: "VueWordCloud",
    data() {
        return {
            msg: 'Welcome to Your Vue.js App',
            worddata: [

            ],
        }
    },

    mounted() {
        this.tagLog().then(
            x => {
                this.initChart();
            }
        )  //  输出结果 1， 1，

        this.tagLog();

    },
    methods: {
        initChart() {
            var tagcom = this
            console.log("tagcom.worddata", tagcom.$data.worddata)
            var myChart = echarts.init(document.getElementById("mywordcloud"));
            const option = {
                title: {
                    text: "词云",
                    x: "center"
                },
                backgroundColor: "#fff",
                // tooltip: {
                //   pointFormat: "{series.name}: <b>{point.percentage:.1f}%</b>"
                // },
                series: [
                    {
                        type: "wordCloud",
                        //用来调整词之间的距离
                        gridSize: 10,
                        //用来调整字的大小范围
                        // Text size range which the value in data will be mapped to.
                        // Default to have minimum 12px and maximum 60px size.
                        sizeRange: [14, 60],
                        // Text rotation range and step in degree. Text will be rotated randomly in range [-90,                                                                             90] by rotationStep 45
                        //用来调整词的旋转方向，，[0,0]--代表着没有角度，也就是词为水平方向，需要设置角度参考注释内容
                        // rotationRange: [-45, 0, 45, 90],
                        // rotationRange: [ 0,90],
                        rotationRange: [0, 0],
                        //随机生成字体颜色
                        // maskImage: maskImage,
                        textStyle: {
                            normal: {
                                color: function () {
                                    return (
                                        "rgb(" +
                                        Math.round(Math.random() * 255) +
                                        ", " +
                                        Math.round(Math.random() * 255) +
                                        ", " +
                                        Math.round(Math.random() * 255) +
                                        ")"
                                    );
                                }
                            }
                        },
                        //位置相关设置
                        // Folllowing left/top/width/height/right/bottom are used for positioning the word cloud
                        // Default to be put in the center and has 75% x 80% size.
                        left: "center",
                        top: "center",
                        right: null,
                        bottom: null,
                        width: "200%",
                        height: "200%",
                        //数据
                        data: tagcom.$data.worddata
                    }
                ]
            };
            console.log("tagcom.worddata", tagcom.$data.worddata)
            myChart.setOption(option);
            myChart.on('click', function (params) {
                console.log('myChart----click---:', params, '------', params.data)
            });
        },
        async tagLog() {
            let query = {}
            query.limit = 20
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


            const taglist = await getTagList(query)
            this.worddata = [].concat(taglist)
            ElMessage.success('获取tags成功')
            return
        },
    }
}
</script>