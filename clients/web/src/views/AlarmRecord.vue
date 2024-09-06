<template>
    <div class="button_list">
        <input class="pan" type="date" v-model="startTime" required>
        <input class="pan" type="date" v-model="endTime" required>
        <select class="pan" v-model="station">
            <option v-for="(item, index) in stations" :key="index" :value="item">{{ item.id }}</option>
        </select>
        <button class="pan" type="submit" @click="getAlarmByStation(), getAlarmUpperLower()">
            <span class="material-symbols-outlined"> search </span>
            Search
        </button>
        <Vue3JsonExcel :json-data="tableData" :fields="json_fields" name="AlarmRecord">
            <button class="pan">
                <span class="material-symbols-outlined"> exit_to_app </span>
                Export
            </button>
        </Vue3JsonExcel>
    </div>
    <!-- Alarm Set -->
    <div class="table-wrapper">
        <table class="fl-table">
            <thead>
                <tr>
                    <th>Station</th>
                    <th>Device</th>
                    <th>Upper Limit</th>
                    <th>Lower Limit</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <!-- pH -->
                <tr>
                    <td>{{ AlarmSetting.pH.station_name }}</td>
                    <td>{{ AlarmSetting.pH.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.pH.high_threshold.val">
                    </td>
                    <td>
                        <input type="number" v-model="AlarmSetting.pH.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.pH.high_threshold.uuid,
            AlarmSetting.pH.high_threshold.val,
            AlarmSetting.pH.low_threshold.uuid,
            AlarmSetting.pH.low_threshold.val)"> UPDATE </button>
                    </td>
                </tr>
                <!-- EC -->
                <tr>
                    <td>{{ AlarmSetting.EC.station_name }}</td>
                    <td>{{ AlarmSetting.EC.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.EC.high_threshold.val">
                    </td>
                    <td>
                        <input type="number" v-model="AlarmSetting.EC.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.EC.high_threshold.uuid,
            AlarmSetting.EC.high_threshold.val,
            AlarmSetting.EC.low_threshold.uuid,
            AlarmSetting.EC.low_threshold.val)">UPDATE</button>
                    </td>
                </tr>
                <!-- Temp -->
                <tr>
                    <td>{{ AlarmSetting.Temp.station_name }}</td>
                    <td>{{ AlarmSetting.Temp.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.Temp.high_threshold.val">
                    </td>
                    <td>
                        <input type="number" v-model="AlarmSetting.Temp.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.Temp.high_threshold.uuid,
            AlarmSetting.Temp.high_threshold.val,
            AlarmSetting.Temp.low_threshold.uuid,
            AlarmSetting.Temp.low_threshold.val)">UPDATE</button>
                    </td>
                </tr>
                <!-- COD -->
                <tr>
                    <td>{{ AlarmSetting.COD.station_name }}</td>
                    <td>{{ AlarmSetting.COD.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.COD.high_threshold.val">
                    </td>
                    <td>
                        <input type="number" v-model="AlarmSetting.COD.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.COD.high_threshold.uuid,
            AlarmSetting.COD.high_threshold.val,
            AlarmSetting.COD.low_threshold.uuid,
            AlarmSetting.COD.low_threshold.val)">UPDATE</button>
                    </td>
                </tr>
                <!-- Turbidity -->
                <tr>
                    <td>{{ AlarmSetting.Turbidity.station_name }}</td>
                    <td>{{ AlarmSetting.Turbidity.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.Turbidity.high_threshold.val">
                    </td>
                    <td>
                        <input type="number" v-model="AlarmSetting.Turbidity.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.Turbidity.high_threshold.uuid,
            AlarmSetting.Turbidity.high_threshold.val,
            AlarmSetting.Turbidity.low_threshold.uuid,
            AlarmSetting.Turbidity.low_threshold.val)">UPDATE</button>
                    </td>
                </tr>
                <!-- battery_temp -->
                <tr>
                    <td>{{ AlarmSetting.battery_temp.station_name }}</td>
                    <td>{{ AlarmSetting.battery_temp.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.battery_temp.high_threshold.val">
                    </td>
                    <td>
                        <input type="number" v-model="AlarmSetting.battery_temp.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.battery_temp.high_threshold.uuid,
            AlarmSetting.battery_temp.high_threshold.val,
            AlarmSetting.battery_temp.low_threshold.uuid,
            AlarmSetting.battery_temp.low_threshold.val)">UPDATE</button>
                    </td>
                </tr>
                <!-- battery -->
                <tr>
                    <td>{{ AlarmSetting.battery.station_name }}</td>
                    <td>{{ AlarmSetting.battery.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.battery.high_threshold.val">
                    </td>

                    <td>
                        <input type="number" v-model="AlarmSetting.battery.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.battery.high_threshold.uuid,
            AlarmSetting.battery.high_threshold.val,
            AlarmSetting.battery.low_threshold.uuid,
            AlarmSetting.battery.low_threshold.val)">UPDATE</button>
                    </td>
                </tr>
                <!-- signal_strength -->
                <tr>
                    <td>{{ AlarmSetting.signal_strength.station_name }}</td>
                    <td>{{ AlarmSetting.signal_strength.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.signal_strength.high_threshold.val">
                    </td>
                    <td>
                        <input type="number" v-model="AlarmSetting.signal_strength.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.signal_strength.high_threshold.uuid,
            AlarmSetting.signal_strength.high_threshold.val,
            AlarmSetting.signal_strength.low_threshold.uuid,
            AlarmSetting.signal_strength.low_threshold.val)">UPDATE</button>
                    </td>
                </tr>
                <!-- FCM_temp -->
                <tr>
                    <td>{{ AlarmSetting.FCM_temp.station_name }}</td>
                    <td>{{ AlarmSetting.FCM_temp.full_name }}</td>
                    <td>
                        <input type="number" v-model="AlarmSetting.FCM_temp.high_threshold.val">
                    </td>
                    <td>
                        <input type="number" v-model="AlarmSetting.FCM_temp.low_threshold.val">
                    </td>
                    <td>
                        <button class="pan" @click="setAlarmUpperLower(
            AlarmSetting.FCM_temp.high_threshold.uuid,
            AlarmSetting.FCM_temp.high_threshold.val,
            AlarmSetting.FCM_temp.low_threshold.uuid,
            AlarmSetting.FCM_temp.low_threshold.val)">
                            UPDATE</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
    <!-- Alarm Record -->
    <div class="table-wrapper">
        <table class="fl-table">
            <thead>
                <tr>
                    <th>Time</th>
                    <th>Status</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(tr_item, tr_index) in tableData" :key="tr_index">
                    <td>
                        {{ tr_item.occur_time }}
                    </td>
                    <td>
                        {{ tr_item.content }}
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
// 獲取當前日期，設置今天的值
const date = new Date();
// 調整日期為 UTC +8 時區
date.setUTCHours(date.getUTCHours() + 8)
// 獲取當前日並格式化為兩位數
const day = ("0" + date.getDate()).slice(-2);
// 獲取當前月並格式化為兩位數
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
// 定義 tableData 的 ref，未初始化，用於存放表格數據
const tableData = ref();
// 定義 json_fields 的 ref，初始值為對象，用於定義 JSON 欄位
const json_fields = ref({ occur_time: 'occur_time', content: 'content' })
// 定義 AlarmSetting 的 ref，初始值為對象，用於存放警報設置
const AlarmSetting = ref({
    pH: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
    EC: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
    Temp: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
    COD: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
    Turbidity: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
    battery_temp: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
    battery: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
    signal_strength: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
    FCM_temp: {
        station_name: "",
        full_name: "",
        high_threshold: {
            uuid: '',
            val: 0,
        },
        low_threshold: {
            uuid: '',
            val: 0,
        },
    },
})

// 定義 captureNumber 函數，接受輸入並提取數字部分
const captureNumber = (input) => {
    // 定義正則表達式，用於匹配數字部分
    const regex = /X([<>]=?|==)(-?\d+(\.\d+)?)/g;
    // const regex = /X[<=>]=?(\d+(\.\d+)?)/g;
    // 使用 exec 方法匹配规则
    const match = regex.exec(input);
    // 判断是否匹配到规则
    if (match) {
        // 匹配到规则时，返回捕获的数字部分
        return parseFloat(match[2]);
    } else {
        // 未匹配到规则时，返回null或其他标识
        return null;
    }
}

// 定義 getAllStation 函數，用於獲取所有站點數據
const getAllStation = async () => {
    await axios({
        method: "Get",
        url: "/api/v1/Station"
    }).then((response) => {
        // 請求成功後，將返回的數據推入 stations.value 中
        response.data.data.forEach(element => {
            stations.value.push(element)
        });
        // 設置 station 的值為返回數據中的第二個站點
        station.value = response.data.data[1]
    })
        .catch(function (error) {
            console.log(error);
        });
}

// 定義 getAlarmUpperLower 函數，用於獲取警報上下限設置
const getAlarmUpperLower = async () => {
    await axios({
        method: "Get",
        url: "/api/v1/Alarm/Setting/Station",
        params: {
            StationUUID: station.value.uuid,
        }
    }).then((response) => {
        // 請求成功後，遍歷返回的數據並設置警報上下限
        response.data.data.forEach(element => {
            switch (element.physical_quantity.name) {
                // 根據物理量名稱將警報設置數據保存到 AlarmSetting 中
                case "pH":
                    AlarmSetting.value.pH.station_name = station.value.name
                    AlarmSetting.value.pH.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        switch (ele.name) {
                            case "pH_high_threshold":
                                AlarmSetting.value.pH.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.pH.high_threshold.uuid = ele.uuid
                                break;
                            case "pH_low_threshold":
                                AlarmSetting.value.pH.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.pH.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                case "EC":
                    AlarmSetting.value.EC.station_name = station.value.name
                    AlarmSetting.value.EC.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        switch (ele.name) {
                            case "EC_high_threshold":
                                AlarmSetting.value.EC.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.EC.high_threshold.uuid = ele.uuid
                                break;
                            case "EC_low_threshold":
                                AlarmSetting.value.EC.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.EC.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                case "temp":
                    AlarmSetting.value.Temp.station_name = station.value.name
                    AlarmSetting.value.Temp.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        console.log(ele)
                        switch (ele.name) {
                            case "Temp_high_threshold":
                                AlarmSetting.value.Temp.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.Temp.high_threshold.uuid = ele.uuid
                                break;
                            case "Temp_low_threshold":
                                AlarmSetting.value.Temp.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.Temp.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                case "COD":
                    AlarmSetting.value.COD.station_name = station.value.name
                    AlarmSetting.value.COD.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        switch (ele.name) {
                            case "COD_high_threshold":
                                AlarmSetting.value.COD.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.COD.high_threshold.uuid = ele.uuid
                                break;
                            case "COD_low_threshold":
                                AlarmSetting.value.COD.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.COD.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                case "turbidity":
                    AlarmSetting.value.Turbidity.station_name = station.value.name
                    AlarmSetting.value.Turbidity.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        switch (ele.name) {
                            case "Turbidity_high_threshold":
                                AlarmSetting.value.Turbidity.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.Turbidity.high_threshold.uuid = ele.uuid
                                break;
                            case "Turbidity_low_threshold":
                                AlarmSetting.value.Turbidity.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.Turbidity.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                case "battery_temp":
                    AlarmSetting.value.battery_temp.station_name = station.value.name
                    AlarmSetting.value.battery_temp.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        switch (ele.name) {
                            case "battery_temperature_high_threshold":
                                AlarmSetting.value.battery_temp.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.battery_temp.high_threshold.uuid = ele.uuid
                                break;
                            case "battery_temperature_low_threshold":
                                AlarmSetting.value.battery_temp.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.battery_temp.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                case "battery":
                    AlarmSetting.value.battery.station_name = station.value.name
                    AlarmSetting.value.battery.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        switch (ele.name) {
                            case "battery_high_threshold":
                                AlarmSetting.value.battery.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.battery.high_threshold.uuid = ele.uuid
                                break;
                            case "battery_low_threshold":
                                AlarmSetting.value.battery.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.battery.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                case "signal_strength":
                    AlarmSetting.value.signal_strength.station_name = station.value.name
                    AlarmSetting.value.signal_strength.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        switch (ele.name) {
                            case "signal_strength_high_threshold":
                                AlarmSetting.value.signal_strength.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.signal_strength.high_threshold.uuid = ele.uuid
                                break;
                            case "signal_strength_low_threshold":
                                AlarmSetting.value.signal_strength.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.signal_strength.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                case "FCM_temp":
                    AlarmSetting.value.FCM_temp.station_name = station.value.name
                    AlarmSetting.value.FCM_temp.full_name = element.physical_quantity.full_name
                    element.alarm_settings.forEach(ele => {
                        console.log(ele);
                        switch (ele.name) {
                            case "FCM_temperature_high_threshold":
                                AlarmSetting.value.FCM_temp.high_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.FCM_temp.high_threshold.uuid = ele.uuid
                                break;
                            case "FCM_temperature_low_threshold":
                                AlarmSetting.value.FCM_temp.low_threshold.val = captureNumber(ele.boolean_expression)
                                AlarmSetting.value.FCM_temp.low_threshold.uuid = ele.uuid
                                break;
                            default:
                                break;
                        }
                    });
                    break;
                default:
                    break;
            }
        });
    }).catch(function (error) {
        // 請求處理失敗
        console.log(error);
    });
}
// 定義 setAlarmUpperLower 函數，用於設置警報上下限
const setAlarmUpperLower = async (high_uuid, high_val, low_uuid, low_val) => {
    // 發送 PATCH 請求以設置高閾值
    await axios({
        method: "PATCH",
        url: "/api/v1/Alarm/Setting/Expression",
        data: {
            uuid: high_uuid,
            expression: `X>=${high_val}`,
        }
    })
    // 發送 PATCH 請求以設置低閾值
    await axios({
        method: "PATCH",
        url: "/api/v1/Alarm/Setting/Expression",
        data: {
            uuid: low_uuid,
            expression: `X<=${low_val}`,
        },
    }).then((response) => {
        // console.log(response);
        alert(response.data.message)
    }).catch(function (error) {
        // 請求處理失敗
        // console.log(error);
        alert(error.data.message)
    });

}
// 定義 getAlarmByStation 函數，用於根據站點獲取警報信息
const getAlarmByStation = async () => {
    await axios({
        method: "POST",
        url: "/api/v1/Alarm/Station",
        data: {
            station_uuid: station.value.uuid,
            start_time: formatToLocalDateTime(startTime.value),
            end_time: endTime.value === today ? formatToLocalDateTime(date) : formatToLocalDateTime(endTime.value).replace('00:00:00', '23:59:59'),
        },
    }).then((response) => {
        console.log(response);
        tableData.value = []
        if (response.data.data) {
            response.data.data.forEach(element => {
                element.occur_time = formatDateTime(element.occur_time) // 格式化警報發生時間
            })
            tableData.value = response.data.data // 將警報數據保存到 tableData 中
        }
    }).catch(function (error) {
        // 請求處理失敗
        console.log(error);
    });
}
onMounted(async () => {
    await getAllStation() // 獲取所有站點數據
    await getAlarmByStation() // 獲取當前選中站點的警報數據
    await getAlarmUpperLower() // 獲取當前選中站點的警報上下限設置
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

/* Chrome, Safari, Edge, Opera */
input[type=number]::-webkit-outer-spin-button,
input[type=number]::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

/* Firefox */
input[type=number] {
    -moz-appearance: textfield;
}

input,
button {
    pointer-events: auto;
    cursor: pointer;
    background: #005AB5;
    border: none;
    padding: 0.5rem 1rem;
    margin: 0;
    font-family: inherit;
    font-size: 0.9rem;
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
    /* width: 80%; */
    box-shadow: 0px 35px 50px rgba(0, 0, 0, 0.2);
}

.fl-table {
    border-radius: 5px;
    font-size: 15px;
    font-weight: normal;
    border: none;
    border-collapse: collapse;
    width: 100%;
    max-width: 100%;
    /* white-space: nowrap; */
    background-color: white;
    table-layout: fixed;
}

.fl-table td,
.fl-table th {
    text-align: center;
    padding: 8px;
}

.fl-table td {
    border-right: 1px solid #f8f8f8;
    font-size: 1.2rem;
    overflow: visible;
    /* font-weight: 900; */
}

.fl-table td button .material-symbols-outlined {
    font-size: 1.3rem;
}


.fl-table td button span {
    font-size: 0.8rem;
}

.fl-table td:last-child {
    border-right: none;
}

.fl-table thead th {
    color: #ffffff;
    background: #005AB5;
}


.fl-table thead th:nth-child(odd) {
    color: #ffffff;
    background: #324960;
}

.fl-table tr:nth-child(even) {
    background: #F8F8F8;
}

td input {
    width: -webkit-fill-available;
    background: inherit;
    text-align: center;
    font-size: 1rem;
    font-weight: 900;
    padding: 0;
}



/* Responsive */
@media (max-width: 1024px) {}

@media (max-width: 720px) {
    .fl-table td button {
        max-width: 100%;
        font-size: 0.6rem;
        padding: 0.3rem 0.6rem;
    }
}

@media (max-width: 320px) {
    .fl-table td button {
        font-size: 0.4rem;
    }
}
</style>