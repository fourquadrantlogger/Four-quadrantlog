<template>
<div v-html="cdiv+cid+cdiv_"></div>
</template>
<script lang="javascript">

import * as echarts from "echarts";
import "echarts-wordcloud/dist/echarts-wordcloud";
import "echarts-wordcloud/dist/echarts-wordcloud.min";


export default {
    name: "VueWordCloud",
    data: function () {
        return {
            cdiv: ` <div id="`,
            cid: Math.random().toString(36).slice(-6),
            cdiv_: `" style="width:100%;height:100%;"></div>`,
        }
    },
    props: {
        datas: [],
        title:'',
    },
    watch: {
        datas: { //深度监听，可监听到对象、数组的变化
            handler(newV, oldV) {
                // do something, 可使用this
                console.log(newV, oldV)
                this.initChart()
            },
            deep: true
        },
    },
    methods: {
        initChart() {
            var tagcom = this
            console.log("datas", tagcom.$props.datas)
            var chartquadrant = echarts.init(document.getElementById(tagcom.$data.cid));

            var series = {
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
                                Math.round(Math.random() * 160) +
                                ", " +
                                Math.round(Math.random() * 160) +
                                ", " +
                                Math.round(Math.random() * 160) +
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
                data: tagcom.$props.datas,
            }

            let option = {
                title: {
                    text: tagcom.$data.title,
                    x: "center"
                },
                backgroundColor: "#fff",
                // tooltip: {
                //   pointFormat: "{series.name}: <b>{point.percentage:.1f}%</b>"
                // },
                series: [
                    series
                ]
            };


            chartquadrant.setOption(option);

            var chartclick = function (ps) {
                console.log('myChart----click---:', ps, '------', ps.data)
                tagcom.$router.push({
                    path: "/table",
                    query: {
                        'atypequery': ps.data.name,
                    }
                })
            }
            chartquadrant.on('click', chartclick);
        },
    }
}
</script>