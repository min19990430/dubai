<template>
    <div class="button_list">
        <input class="pan" type="date" lang="en" v-model="startTime" required>
        <input class="pan" type="date" lang="en" v-model="endTime" required>
        <select class="pan" v-model="station">
            <option v-for="(item, index) in stations" :key="index" :value="item.uuid">{{ item.id }}</option>
        </select>
        <select class="pan" v-model="interval">
            <option>5m</option>
            <option>10m</option>
        </select>
        <button class="pan" type="submit" @click="searchTimeSeriesByStation()">
            <span class="material-symbols-outlined"> search </span> Search </button>
    </div>
    <div class="insights">
        <div>
            <Line :options="chartOptions" width="1000px" :data="chartData" v-if="loaded" />
        </div>
    </div>
</template>

<script setup>
// 從 'vue' 中引入 ref 和 onMounted
import { ref, onMounted } from 'vue';
// 引入 axios 用於發送 HTTP 請求
import axios from 'axios'
// 引入 vue-chartjs 的相關組件
import { Line } from 'vue-chartjs'
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from 'chart.js'
// 引入 chartjs-plugin-zoom 用於圖表縮放
import zoomPlugin from 'chartjs-plugin-zoom';
// 註冊 chart.js 所需的組件和插件
ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, zoomPlugin)
// 引入時間格式化相關的函數
import { extractTimeIfNotZero, formatToLocalDateTime } from '@/assets/js/formatDateTime';
// 獲取當天日期
const date = new Date();
const day = ("0" + date.getDate()).slice(-2);
const month = ("0" + (date.getMonth() + 1)).slice(-2);
const today = date.getFullYear() + "-" + month + "-" + day;
// date.setUTCHours(date.getUTCHours() + 8)

const station = ref("")
const stations = ref([])
const startTime = ref(today)
const endTime = ref(today)
const interval = ref("5m")
// 宣告 ref 變數來存儲圖表資料
const chartData = ref({
    labels: [],
    datasets: [
        {
            label: 'pH',
            backgroundColor: '#FF0000',
            borderColor: '#FF0000', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
        {
            label: 'EC',
            backgroundColor: '#FF7F00',
            borderColor: '#FF7F00', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
        {
            label: 'COD',
            backgroundColor: '#FFFF00',
            borderColor: '#FFFF00', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
        {
            label: 'NTU',
            backgroundColor: '#00FF00',
            borderColor: '#00FF00', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
        {
            label: 'TEMP.',
            backgroundColor: '#00FFFF',
            borderColor: '#00FFFF', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
        {
            label: 'BATTERY TEMP.',
            backgroundColor: '#0000FF',
            borderColor: '#0000FF', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
        {
            label: 'BATTERY CAPACITY',
            backgroundColor: '#8B00FF',
            borderColor: '#8B00FF', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
        {
            label: 'SIGNAL STRENGTH',
            backgroundColor: '#BEBEBE',
            borderColor: '#BEBEBE', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
        {
            label: 'FCM DEVICE TEMP.',
            backgroundColor: '#000000',
            borderColor: '#000000', // 设置线条颜色
            borderWidth: 2, // 设置线条宽度
            data: []
        },
    ]
})
const loaded = ref(false)
// 宣告 ref 變數來存儲圖表選項
const chartOptions = ref({
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
        legend: {
            position: 'left', // 将图例放置在左侧
            labels: {
                // 自定義生成圖例標籤的函數
                generateLabels(chart) {
                    const original = ChartJS.defaults.plugins.legend.labels.generateLabels;
                    const labels = original(chart);
                    console.log(labels);
                    labels.unshift({
                        text: 'SELECT ALL',
                        fillStyle: 'transparent',
                        hidden: false,
                        lineCap: 'butt',
                        lineDash: [],
                        lineDashOffset: 0,
                        lineJoin: 'miter',
                        lineWidth: 0,
                        strokeStyle: 'transparent',
                        pointStyle: 'rect',
                        datasetIndex: -1, // 自定義屬性，用於識別自定義標籤
                    });
                    return labels;
                },
            },
            //處理圖例項目的點擊事件
            onClick(e, legendItem, legend) {
                const chart = legend.chart;
                if (legendItem.datasetIndex === -1) {
                    // Custom handling for "全选"
                    const allVisible = chart.data.datasets.every((dataset, index) => !chart.getDatasetMeta(index).hidden);
                    chart.data.datasets.forEach((dataset, index) => {
                        chart.getDatasetMeta(index).hidden = allVisible;
                    });
                    chart.update();
                    console.log(chart.update());
                } else {
                    // Default click handling
                    const index = legendItem.datasetIndex;
                    chart.getDatasetMeta(index).hidden = chart.isDatasetVisible(index);
                    chart.update();
                }
            },
        },
        zoom: {
            pan: {
                enabled: true,
                mode: 'xy',
            },
            zoom: {
                wheel: {
                    enabled: true,
                },
                pinch: {
                    enabled: true,
                },
                mode: 'xy',
            },
        },
    }
})
// 獲取所有測站資訊的函數
const getAllStation = async () => {
    await axios({
        method: "Get",
        url: "/api/v1/Station"
    }).then((response) => {
        response.data.data.forEach(element => {
            stations.value.push(element)
        });
        station.value = response.data.data[1].uuid
    })
        .catch(function (error) {
            console.log(error);
        });
}
// 根據測站查詢時間序列資料的函數
const searchTimeSeriesByStation = () => {
    loaded.value = false
    axios({
        method: "POST",
        url: "/api/v1/TimeSeries/Station",
        data: {
            station_uuid: station.value,
            start_time: formatToLocalDateTime(startTime.value),
            end_time: endTime.value === today ? date : formatToLocalDateTime(endTime.value).replace('00:00:00', '23:59:59'),
            interval: interval.value,
        }
    }).then((response) => {
        loaded.value = false
        chartData.value.labels = []
        // console.log(chartData.value.labels);
        chartData.value.datasets.forEach((ele) => {
            ele.data = []
        })
        response.data.data.data.forEach((element) => {
            let newTime = extractTimeIfNotZero(element[0])
            chartData.value.labels.push(newTime)
            chartData.value.datasets[0].data.push(element[1]) // ph
            chartData.value.datasets[1].data.push(element[2]) // ec
            chartData.value.datasets[2].data.push(element[3]) // COD
            chartData.value.datasets[3].data.push(element[4]) // NTU
            chartData.value.datasets[4].data.push(element[5]) // 溫度
            chartData.value.datasets[5].data.push(element[6]) // 電池溫度
            chartData.value.datasets[6].data.push(element[7]) // 電池電量
            chartData.value.datasets[7].data.push(element[8]) // 訊號強度
            chartData.value.datasets[8].data.push(element[9]) // FCM箱內溫度
        })
        loaded.value = true
    })
        .catch(function (error) {
            console.log(error);
        });

}

onMounted(async () => {
    await getAllStation()
    searchTimeSeriesByStation()

});

</script>
<style scoped>
.button_list {
    text-align: center;
}

.pan {
    margin: 0 0.5rem 0.5rem 0;
}

.pan:last-child {
    margin-right: 0;
}

input[type="date"]::-webkit-calendar-picker-indicator {
    position: absolute;
    padding-left: 100%;
    right: 1rem;
}

input[type="date"]::-webkit-calendar-picker-indicator:hover {
    filter: invert(1);
}

input,
button {
    pointer-events: auto;
    cursor: pointer;
    background: #005AB5;
    border: none;

    padding: 0.5rem 1.5rem;
    margin: 0;
    font-family: inherit;
    font-size: inherit;
    position: relative;
    display: inline-block;
    z-index: 1;
}

select {
    pointer-events: auto;
    cursor: pointer;
    background: #BEBEBE;
    border: none;
    padding: 0.6rem 1.5rem;
    margin: 0;
    font-family: inherit;
    font-size: inherit;
    position: relative;
    display: inline-block;
    z-index: 1;
}


input::before,
input::after,
select::before,
select::after,
button::before,
button::after {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
}

.pan {
    font-family: aktiv-grotesk-extended, sans-serif;
    font-weight: 700;
    border-radius: 3rem;
    overflow: hidden;
    color: #000000;
}

.pan:hover {
    color: white;
}

.pan::before {
    content: "";
    background: #BEBEBE;
    transition: transform 0.3s cubic-bezier(0.7, 0, 0.2, 1);
    z-index: -1;
}

.pan:hover::before {
    transform: translate3d(0, -100%, 0);
}

.material-symbols-outlined {
    vertical-align: middle;
    color: black;
    font-size: 1.2rem;
    margin-right: 0.3rem;
}

.material-symbols-outlined:hover {
    color: white;
}

/* =============== insights =============== */
.insights>div {
    background: #fff;
    padding: 1.8rem;
    border-radius: 2rem;
    margin-top: 1rem;
    margin-bottom: 1rem;
    box-shadow: 0 2rem 3rem rgba(132, 139, 200, 0.18);
    transition: all 300ms ease;
}

.insights>div:hover {
    box-shadow: none;
}


/* Responsive */
@media (max-width: 1024px) {}

@media (max-width: 720px) {}

@media (max-width: 320px) {}
</style>