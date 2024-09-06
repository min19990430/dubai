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
            <span class="material-symbols-outlined"> search </span>
            Search
        </button>
        <Vue3JsonExcel :json-data="tableData" :fields="json_fields" name="Historical">
            <button class="pan">
                <span class="material-symbols-outlined"> exit_to_app </span>
                Export
            </button>
        </Vue3JsonExcel>
    </div>
    <div class="table-wrapper">
        <table class="fl-table">
            <thead>
                <tr>
                    <th v-for="(item, index) in tableHead" :key="index">
                        {{ item.full_name }}
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(tr_item, tr_index) in tableData" :key="tr_index">
                    <td v-for="(td_item, td_index) in tr_item" :key="td_index">
                        {{ td_item }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup>
// 從 Vue 庫中引入 ref 和 onMounted 函數
import { ref, onMounted } from 'vue';
// 從 axios 庫中引入 axios 物件，用於發送 HTTP 請求
import axios from 'axios'
// 從 vue3-json-excel 庫中引入 Vue3JsonExcel 元件，用於 Excel 匯出
import { Vue3JsonExcel } from 'vue3-json-excel';
// 從自定義模組中引入 formatDateTime 和 formatToLocalDateTime 函數，用於日期格式化
import { formatDateTime, formatToLocalDateTime } from '@/assets/js/formatDateTime';
// 獲取當前日期
const date = new Date();
// 獲取當前日，並格式化為兩位數
const day = ("0" + date.getDate()).slice(-2);
// 獲取當前月，並格式化為兩位數
const month = ("0" + (date.getMonth() + 1)).slice(-2);
// 獲取今天的日期，格式為 yyyy-MM-dd
const today = date.getFullYear() + "-" + month + "-" + day;
// 定義 stations 的 ref，初始值為空數組，用於存放站點數據
const stations = ref([])
// 定義 station 的 ref，初始值為空字符串，用於存放選定的站點
const station = ref("")
// 定義 startTime 的 ref，初始值為今天，用於存放開始時間
const startTime = ref(today)
// 定義 endTime 的 ref，初始值為今天，用於存放結束時間
const endTime = ref(today)
// 定義 interval 的 ref，初始值為 "5m"，用於存放時間間隔
const interval = ref("5m")
// 定義 tableHead 的 ref，未初始化，用於存放表格標題
const tableHead = ref();
// 定義 tableData 的 ref，未初始化，用於存放表格數據
const tableData = ref();
// 定義 json_fields 的 ref，未初始化，用於存放 JSON 欄位
const json_fields = ref({})
// 定義 getAllStation 函數，用於獲取所有站點數據
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
// 定義 searchTimeSeriesByStation 函數，用於根據站點搜尋時間序列數據
const searchTimeSeriesByStation = async () => {
    await axios({
        method: "POST",
        url: "/api/v1/TimeSeries/Station",
        data: {
            station_uuid: station.value,
            start_time: formatToLocalDateTime(startTime.value),
            end_time: endTime.value === today ? date : formatToLocalDateTime(endTime.value).replace('00:00:00', '23:59:59'),
            interval: interval.value,
            reverse: true,
        }
    }).then((response) => {
        response.data.data.columns.forEach((element, index) => {
            json_fields.value[element.full_name] = index.toString()
        });
        tableHead.value = response.data.data.columns
        response.data.data.data.forEach(element => {
            let newTime = formatDateTime(element[0])
            element[0] = newTime.slice(0, 16)
        })
        tableData.value = response.data.data.data
    })
        .catch(function (error) {
            console.log(error);
        });
}
onMounted(async () => {
    await getAllStation()
    await searchTimeSeriesByStation()
})
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
    color: black;
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


.table-wrapper {
    margin: 20px auto;
    box-shadow: 0px 35px 50px rgba(0, 0, 0, 0.2);
    height: 80vh;
    overflow: auto;
}


.fl-table {
    margin: 0 auto;
    border-radius: 5px;
    font-size: 15px;
    font-weight: normal;
    border: none;
    border-collapse: collapse;
    width: 100%;
    max-width: 100%;
    background-color: white;
    table-layout: fixed;
}

.fl-table td,
.fl-table th {
    text-align: center;
    padding: 8px;
}

.fl-table th {
    position: sticky;
    top: 0;
    color: #ffffff;
    background: #005AB5;
    font-size: 1rem;
}

.fl-table td {
    border-right: 1px solid #f8f8f8;
    font-size: 0.8rem;
}

.fl-table td:last-child {
    border-right: none;
}

.fl-table thead th:nth-child(odd) {
    color: #ffffff;
    background: #324960;
}

.fl-table tr:nth-child(even) {
    background: #F8F8F8;
}

/* Responsive */
@media (max-width: 1024px) {
    .fl-table th {
        word-wrap: break-word;
        /* 允许 <th> 元素的内容换行 */
    }

    .fl-table td {
        white-space: break-spaces;
        /* 禁止 <td> 元素的内容换行 */
    }
}

@media (max-width: 720px) {
    .fl-table {
        display: block;
        width: 100%;
    }

    .table-wrapper:before {
        content: "";
        display: block;
        text-align: right;
        font-size: 11px;
        color: white;
        padding: 0 0 10px;
    }

    .fl-table thead,
    .fl-table tbody,
    .fl-table thead th {
        display: block;
    }

    .fl-table thead th:last-child {
        border-bottom: none;
    }

    .fl-table thead {
        float: left;
    }

    .fl-table tbody {
        width: auto;
        position: relative;
        overflow-y: auto;
    }

    .fl-table td,
    .fl-table th {
        padding: 20px .625em .625em .625em;
        height: 60px;
        vertical-align: middle;
        box-sizing: border-box;
        overflow-x: hidden;
        overflow-y: auto;
        width: 120px;
        font-size: 13px;
        text-overflow: ellipsis;

    }

    .fl-table thead th {
        text-align: center;
        border-bottom: 1px solid #f7f7f9;
    }

    .fl-table tbody tr {
        display: table-cell;
    }

    .fl-table tbody tr:nth-child(odd) {
        background: none;
    }

    .fl-table tr:nth-child(even) {
        background: transparent;
    }

    .fl-table tr td:nth-child(odd) {
        background: #F8F8F8;
        border-right: 1px solid #E6E4E4;
    }

    .fl-table tr td:nth-child(even) {
        border-right: 1px solid #E6E4E4;
    }

    .fl-table tbody td {
        display: block;
        text-align: center;
    }
}

@media (max-width: 320px) {}
</style>