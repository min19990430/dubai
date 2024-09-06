<template>
    <div class="layout">
        <div class="s01">
            <p>Monitoring Point Location {{ s01.id }}</p>
            <div class="location">
                <img src="@/assets/img/Vector.png">
                <h6>{{ s01.dms }}</h6>
            </div>
            <div class="content">
                <div class="ph">
                    <span class="material-symbols-outlined"> water_ph </span>
                    <h4>{{ s01.ph.full_name }}</h4>
                    <h5>{{ s01.ph.val }} </h5>
                    <ProgressBar :value="s01.ph.percent" :max="100" />
                </div>
                <div class="electrical">
                    <span class="material-symbols-outlined"> water_ec </span>
                    <h4>{{ s01.ec.full_name }}</h4>
                    <h5>{{ s01.ec.val }} µS/cm</h5>
                    <ProgressBar :value="s01.ec.percent" :max="100" />
                </div>
                <div class="temperature">
                    <span class="material-symbols-outlined"> dew_point </span>
                    <h4>{{ s01.temp.full_name }}</h4>
                    <h5>{{ s01.temp.val }} °C</h5>
                    <ProgressBar :value="s01.temp.percent" :max="100" />
                </div>
                <div class="cod">
                    <span class="material-symbols-outlined"> spo2 </span>
                    <h4>{{ s01.cod.full_name }}</h4>
                    <h5>{{ s01.cod.val }} mg/L</h5>
                    <ProgressBar :value="s01.cod.percent" :max="100" />
                </div>
                <div class="turbidity">
                    <span class="material-symbols-outlined"> water_do </span>
                    <h4>{{ s01.turbidity.full_name }}</h4>
                    <h5>{{ s01.turbidity.val }} NTU</h5>
                    <ProgressBar :value="s01.turbidity.percent" :max="100" />
                </div>
                <DateTimer :date_time="s01.update_time" />
            </div>
        </div>
        <div class="gis">
            <div class="insights">
                <GoogleMap class="google_map" :key="mapKey" :api-key="GOOGLE_MAP_API" language="en"
                    mapTypeId="satellite" :center="center" :zoom="13">
                    <Marker :options="{ position: s01.coordinate, icon: s01.google_map_icon }">
                        <InfoWindow>
                            <ul>
                                <li>{{ s01.battery_temp.full_name }} : {{ s01.battery_temp.val }} °C</li>
                                <li>{{ s01.battery.full_name }} : {{ s01.battery.val }} %</li>
                                <li>{{ s01.signal_strength.full_name }} : {{ s01.signal_strength.strength }}</li>
                                <li>{{ s01.fcm_temp.full_name }} : {{ s01.fcm_temp.val }} °C</li>
                            </ul>
                        </InfoWindow>
                    </Marker>
                    <Marker :options="{ position: s02.coordinate, icon: s02.google_map_icon }">
                        <InfoWindow>
                            <ul>
                                <li>{{ s02.battery_temp.full_name }} : {{ s02.battery_temp.val }} °C</li>
                                <li>{{ s02.battery.full_name }} : {{ s02.battery.val }} %</li>
                                <li>{{ s02.signal_strength.full_name }} : {{ s02.signal_strength.strength }}</li>
                                <li>{{ s02.fcm_temp.full_name }} : {{ s02.fcm_temp.val }} °C</li>
                            </ul>
                        </InfoWindow>
                    </Marker>
                </GoogleMap>
            </div>
        </div>
        <div class="s02">
            <p>Monitoring Point Location {{ s02.id }}</p>
            <div class="location">
                <img src="@/assets/img/Vector.png">
                <h6>{{ s02.dms }}</h6>
            </div>
            <div class="content">
                <div class="ph">
                    <span class="material-symbols-outlined"> water_ph </span>
                    <h4>{{ s02.ph.full_name }}</h4>
                    <h5>{{ s02.ph.val }} </h5>
                    <ProgressBar :value="s02.ph.percent" :max="100" />
                </div>
                <div class="electrical">
                    <span class="material-symbols-outlined"> water_ec </span>
                    <h4>{{ s02.ec.full_name }}</h4>
                    <h5>{{ s02.ec.val }} µS/cm</h5>
                    <ProgressBar :value="s02.ec.percent" :max="100" />
                </div>
                <div class="temperature">
                    <span class="material-symbols-outlined"> dew_point </span>
                    <h4>{{ s02.temp.full_name }}</h4>
                    <h5>{{ s02.temp.val }} °C</h5>
                    <ProgressBar :value="s02.temp.percent" :max="100" />
                </div>
                <div class="cod">
                    <span class="material-symbols-outlined"> spo2 </span>
                    <h4>{{ s02.cod.full_name }}</h4>
                    <h5>{{ s02.cod.val }} mg/L</h5>
                    <ProgressBar :value="s02.cod.percent" :max="100" />
                </div>
                <div class="turbidity">
                    <span class="material-symbols-outlined"> water_do </span>
                    <h4>{{ s02.turbidity.full_name }}</h4>
                    <h5>{{ s02.turbidity.val }} NTU</h5>
                    <ProgressBar :value="s02.turbidity.percent" :max="100" />
                </div>
                <DateTimer :date_time="s02.update_time" />
            </div>
        </div>
    </div>
    <div class="alarm">
        <div class="content">
            <div class="title">
                <h1>Immediate</h1>
                <h2>Alarm Notification</h2>
            </div>
            <div>
                <table>
                    <thead>
                        <tr>
                            <th>Point</th>
                            <th>Time</th>
                            <th>Status</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(item, index) in alarmData" :key="index">
                            <td>{{ item.station_name }}</td>
                            <td>{{ item.occur_time }}</td>
                            <td>{{ item.content }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script setup>
// 引入必要的 Vue 相關套件
import { ref, onMounted, onBeforeUnmount } from 'vue';
// 引入進度條元件
import ProgressBar from "@/components/ProgressBar.vue";
// 引入日期時間元件
import DateTimer from "@/components/DateTimerVue.vue";
// 引入 Axios 套件處理 HTTP 請求
import axios from 'axios'
// 引入日期時間格式化函數
import { formatDateTime } from '@/assets/js/formatDateTime';
// 引入經緯度轉換函數
import { convertLatLngToDMS } from '@/assets/js/convertLatLngToDMS';
// 引入 Google 地圖相關元件
import { GoogleMap, Marker, InfoWindow } from 'vue3-google-map'
// 定義 Google Maps API 金鑰
const GOOGLE_MAP_API = process.env.VUE_APP_GOOGLE_MAP_API
// 中心點經緯度，使用 ref 包裝成 reactive
const center = ref({ lng: 0, lat: 0 })
// Google 地圖金鑰狀態
const mapKey = ref(true)
const s01 = ref({
    id: '',
    update_time: '',
    coordinate: { lng: 55.3551, lat: 25.1976 },
    google_map_icon: { url: require('@/assets/img/FCM.png'), scaledSize: { width: 60, height: 60 } },
    dms: '', //度分秒
    ph: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    ec: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    cod: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    temp: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    turbidity: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    battery_temp: {
        full_name: '',
        val: 0,
    },
    battery: {
        full_name: '',
        val: 0,
    },
    signal_strength: {
        full_name: '',
        val: 0,
        strength: '',
    },
    fcm_temp: {
        full_name: '',
        val: 0
    },
})
const s02 = ref({
    id: '',
    update_time: '',
    coordinate: { lng: 46.674444, lat: 26.711389 },
    google_map_icon: { url: require('@/assets/img/FCM.png'), scaledSize: { width: 60, height: 60 } },
    dms: '', //度分秒
    ph: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    ec: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    cod: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    temp: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    turbidity: {
        full_name: '',
        val: 0,
        percent: 0,
        color: "#dfdfdf",
    },
    battery_temp: {
        full_name: '',
        val: 0,
    },
    battery: {
        full_name: '',
        val: 0,
    },
    signal_strength: {
        full_name: '',
        val: 0,
        strength: '',
    },
    fcm_temp: {
        full_name: '',
        val: 0
    },
})
const alarmData = ref()
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
// 取得最新資料函數
const getLastData = async () => {
    await axios
        .get("/api/v1/Last")
        .then((response) => {
            response.data.data.last.forEach(element => {
                switch (element.device.uuid) {
                    // M6288 FCM1
                    case 'f52132d2-8a8f-4acc-afbd-6ec556b12182':
                        s01.value.coordinate.lat = element.station.lat
                        s01.value.coordinate.lng = element.station.lon
                        s01.value.dms = convertLatLngToDMS(element.station.lat, element.station.lon)
                        s01.value.id = element.station.id
                        s01.value.update_time = element.device.update_time;
                        if (element.device.is_connected) {
                            s01.value.google_map_icon.url = require('@/assets/img/FCM-Normal.png')
                            console.log(element)
                        } else {
                            s01.value.google_map_icon.url = require('@/assets/img/FCM-StopFunctioning.png')
                        }
                        element.physical_quantities.forEach(ele => {
                            console.log(ele)
                            switch (ele.uuid) {
                                case '47d28dc4-b888-4880-ab23-3d1c46b24f06':
                                    s01.value.ph.full_name = ele.full_name
                                    s01.value.ph.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s01.value.ph.color = "#ffd700"
                                        s01.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    } else if (ele.status_code == 10) {
                                        s01.value.google_map_icon.url = require('@/assets/img/FCM-Normal.png')
                                    }
                                    break;
                                case '8df500a9-41b8-4aa0-b072-2fa2313f9677':
                                    s01.value.ec.full_name = ele.full_name
                                    s01.value.ec.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s01.value.ec.color = "#ffd700"
                                        s01.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case 'c9d30b7c-bcdd-466b-9c75-29dfcf3dc59f':
                                    s01.value.temp.full_name = ele.full_name
                                    s01.value.temp.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s01.value.temp.color = "#ffd700"
                                        s01.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case 'f17c29aa-9ee0-476b-9b89-4d624fd1af82':
                                    s01.value.cod.full_name = ele.full_name
                                    s01.value.cod.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s01.value.cod.color = "#ffd700"
                                        s01.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case '9019c4e1-e40d-4e77-b69e-87f70c4de96d':
                                    s01.value.turbidity.full_name = ele.full_name
                                    s01.value.turbidity.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s01.value.turbidity.color = "#ffd700"
                                        s01.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case '184c9a57-65bc-4271-97d6-b983431e511e':
                                    s01.value.battery_temp.full_name = ele.full_name
                                    s01.value.battery_temp.val = Math.floor(ele.value)
                                    break;
                                case 'a03598ab-8b1a-48ec-8ecf-d1e7838127df':
                                    s01.value.battery.full_name = ele.full_name
                                    s01.value.battery.val = Math.floor(ele.value)
                                    break;
                                case '675a68b9-e5f2-4b25-9ddb-d73cfb4c2d6a':
                                    s01.value.signal_strength.full_name = ele.full_name
                                    s01.value.signal_strength.val = Math.floor(ele.value)
                                    switch (true) {
                                        case ele.value <= 34:
                                            s01.value.signal_strength.strength = 'poor'
                                            break;
                                        case ele.value <= 67 && ele.value > 34:
                                            s01.value.signal_strength.strength = 'Fair'
                                            break;
                                        case ele.value <= 100 && ele.value > 67:
                                            s01.value.signal_strength.strength = 'Good'
                                            break;
                                        default:
                                            break;
                                    }
                                    break;
                                case '5c5d94c5-d1e2-4519-81ed-41ea88def09a':
                                    s01.value.fcm_temp.full_name = ele.full_name
                                    s01.value.fcm_temp.val = Math.floor(ele.value)
                                    break;
                                default:
                                    break;
                            }
                        })
                        break;
                    // M7106 FCM2
                    case '82d7329d-c750-4d97-9f03-c5771e599163':
                        s02.value.id = element.station.id
                        s02.value.update_time = element.device.update_time;
                        s02.value.coordinate.lat = element.station.lat
                        s02.value.coordinate.lng = element.station.lon
                        s02.value.dms = convertLatLngToDMS(element.station.lat, element.station.lon)
                        console.log(element);
                        if (element.device.is_connected) {
                            s02.value.google_map_icon.url = require('@/assets/img/FCM-Normal.png')
                        } else {
                            s02.value.google_map_icon.url = require('@/assets/img/FCM-StopFunctioning.png')
                        }
                        element.physical_quantities.forEach(ele => {
                            switch (ele.uuid) {
                                case '6b39a4db-3323-4013-a5ec-d2ee2bb7f652':
                                    s02.value.ph.full_name = ele.full_name
                                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                                    s02.value.ph.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s02.value.ph.color = "#ffd700"
                                        s02.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case '2f093694-38f7-4650-8ce4-d4c08a872d7f':
                                    s02.value.ec.full_name = ele.full_name
                                    s02.value.ec.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s02.value.ec.color = "#ffd700"
                                        s02.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case '7c35b464-7f7c-45b4-8251-ecb35241c361':
                                    s02.value.temp.full_name = ele.full_name
                                    s02.value.temp.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s02.value.temp.color = "#ffd700"
                                        s02.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case 'e9915e9b-1940-4a8b-aa45-d5114921b863':
                                    s02.value.cod.full_name = ele.full_name
                                    s02.value.cod.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s02.value.cod.color = "#ffd700"
                                        s02.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case '2a81912f-fa83-466d-a3f0-604b1cd78e73':
                                    s02.value.turbidity.full_name = ele.full_name
                                    s02.value.turbidity.val = parseFloat((ele.value).toFixed(2))
                                    if (ele.status_code == 20) {
                                        s02.value.turbidity.color = "#ffd700"
                                        s02.value.google_map_icon.url = require('@/assets/img/FCM-Calibrate.png')
                                    }
                                    break;
                                case '8bf89996-f693-40c5-8995-77160ed08fd8':
                                    s02.value.battery_temp.full_name = ele.full_name
                                    s02.value.battery_temp.val = Math.floor(ele.value)
                                    break;
                                case 'ee45cc81-6375-448e-b940-9a433b771cf9':
                                    s02.value.battery.full_name = ele.full_name
                                    s02.value.battery.val = Math.floor(ele.value)
                                    break;
                                case 'c9e0cee6-193e-4ba5-9e61-2edbf620b4cb':
                                    s02.value.signal_strength.full_name = ele.full_name
                                    s02.value.signal_strength.val = Math.floor(ele.value)
                                    switch (true) {
                                        case ele.value <= 34:
                                            s02.value.signal_strength.strength = 'poor'
                                            break;
                                        case ele.value <= 67 && ele.value > 34:
                                            s02.value.signal_strength.strength = 'Fair'
                                            break;
                                        case ele.value <= 100 && ele.value > 67:
                                            s02.value.signal_strength.strength = 'Good'
                                            break;
                                        default:
                                            break;
                                    }
                                    break;
                                case '995ec435-2c5f-41c6-9a16-c47c3b320364':
                                    s02.value.fcm_temp.full_name = ele.full_name
                                    s02.value.fcm_temp.val = Math.floor(ele.value)
                                    break;
                                default:
                                    break;
                            }
                        })
                        break;
                    default:
                        break;
                }
            })
            if (response.data.data.alarm) {
                response.data.data.alarm.forEach(element => {
                    element.occur_time = formatDateTime(element.occur_time)
                })
                alarmData.value = response.data.data.alarm.slice(0, 5)
            }
        })
        .catch(function (error) {
            console.log(error);
        });

    //getAlarmUpperLower S01
    await axios({
        method: "Get",
        url: "/api/v1/Alarm/Setting/Station",
        params: {
            StationUUID: 'e2232881-e996-46be-a0a2-95a2983fdb42',
        }
    }).then((response) => {
        response.data.data.forEach(element => {
            switch (element.physical_quantity.uuid) {
                //pH
                case "47d28dc4-b888-4880-ab23-3d1c46b24f06": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "dd867c32-fad2-4c59-917b-0f35458275d8":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "42aa4b4b-bda1-44e6-958b-61f878f8d049":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.ph.percent = ((s01.value.ph.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //EC
                case "8df500a9-41b8-4aa0-b072-2fa2313f9677": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "1b3a1bd3-6575-4f93-8821-8bf7e453e908":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "9576669f-c69e-4ad6-a244-f75e5c936b7f":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.ec.percent = ((s01.value.ec.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //TEMP.
                case "c9d30b7c-bcdd-466b-9c75-29dfcf3dc59f": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "5299c1ce-29ba-4819-840b-e8ce6895fe88":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "1b5bfea1-eaf3-4296-b774-8108df9a9b8f":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.temp.percent = ((s01.value.temp.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //COD
                case "f17c29aa-9ee0-476b-9b89-4d624fd1af82": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "3d806f62-869d-403e-9581-5107a9d983c5":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "c293e5d8-6117-433f-8527-8705f55f9741":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.cod.percent = ((s01.value.cod.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //Turbidity
                case "9019c4e1-e40d-4e77-b69e-87f70c4de96d": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "751cc912-cd23-4f3c-8a31-03cf41d713b6":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "8a8abeaf-7cc3-4843-bdae-14a7c431e6a6":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.turbidity.percent = ((s01.value.turbidity.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                default:
                    break;
            }
        });
    }).catch(function (error) {
        // 請求處理失敗
        console.log(error);
    });
    //getAlarmUpperLower s02
    await axios({
        method: "Get",
        url: "/api/v1/Alarm/Setting/Station",
        params: {
            StationUUID: '11465818-37f3-4c81-a093-296ea9dee685',
        }
    }).then((response) => {
        response.data.data.forEach(element => {
            switch (element.physical_quantity.uuid) {
                //pH
                case "6b39a4db-3323-4013-a5ec-d2ee2bb7f652": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "4bf608a5-7fc1-479d-b23f-42b8c4c472bf":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "ff552487-93b3-4b14-a6c4-2cc1a4e91d05":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.ph.percent = ((s02.value.ph.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //EC
                case "2f093694-38f7-4650-8ce4-d4c08a872d7f": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "e446d792-0903-4467-9854-5d8258dbff90":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "f5459d8c-df91-4352-b018-e33f9a6f330c":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.ec.percent = ((s02.value.ec.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //Temp
                case "7c35b464-7f7c-45b4-8251-ecb35241c361": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "0a62361b-f5eb-4de2-840e-a754bdbaf9ab":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "55772ca9-e49d-4d1b-8fd8-2a7f8b2905db":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.temp.percent = ((s02.value.temp.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //COD
                case "e9915e9b-1940-4a8b-aa45-d5114921b863": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "d41f5e72-d218-4db4-8b1e-b43b6c3ad3f8":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "34b06a39-30dd-4df9-b3ed-1b582a6d5f63":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.cod.percent = ((s02.value.cod.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //Turbidity
                case "2a81912f-fa83-466d-a3f0-604b1cd78e73": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "b9eaaed5-6ec3-4c05-b451-b47d7605e551":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "bbee8f51-22e2-4498-b776-dc63aea141c1":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.turbidity.percent = ((s02.value.turbidity.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                default:
                    break;
            }
        });
    }).catch(function (error) {
        // 請求處理失敗
        console.log(error);
    });

}
let LastDatTimer = null // 計時器
onMounted(() => {
    getLastData()

    // 取得 Google 地圖中心點
    axios.get("/api/v1/Last")
        .then((response) => {
            response.data.data.last.forEach(element => {
                center.value.lat += element.station.lat
                center.value.lng += element.station.lon
            })
            center.value.lat = center.value.lat / 2
            center.value.lng = center.value.lng / 2
            mapKey.value = !mapKey.value
        })
        .catch(function (error) {
            console.log(error);
        });
    // 設置定時器每 60 秒更新資料
    LastDatTimer = setInterval(getLastData, 1000 * 60)
})
onBeforeUnmount(() => {
    clearInterval(LastDatTimer)
    LastDatTimer = null
})
</script>

<style scoped>
.layout {
    display: grid;
    grid-template-columns: 1fr 2.5fr 1fr;
    text-align: center;
    row-gap: 0.8rem;
}

p {
    font-size: 2.5rem;
    color: #F1F1F1;
}

.location {
    display: flex;
    align-items: center;
    justify-content: center;
}

h6 {
    font-size: 1rem;
    color: #dfdfdf;
    margin-left: 0.6rem;
}

h5 {
    font-size: 1.5rem;
    color: #dfdfdf;
}

.s01 .content,
.s02 .content {
    display: grid;
    height: auto;
    grid-template-rows: repeat(6, 1fr);
}

.gis {
    display: grid;
    grid-template-columns: repeat(1, 1fr);
    background-color: rgb(26 26 26 / 30%);
    border-radius: 1rem;
    padding: 1rem;
}

.gis .insights {
    border-radius: 1.8rem;
}

.google_map {
    width: 100%;
    height: 100%;
}

.gm-style-iw {
    border-radius: 1rem;
    background-color: rgb(26 26 26 / 30%);
}

.label h6 {
    font-size: 1.5rem;
    color: white;
    margin-top: 1.8rem;
    text-align: center;
}

.ph,
.electrical,
.temperature,
.cod,
.turbidity {
    display: grid;
    grid-template-columns: 0.6fr 1fr;
    /* grid-template-rows: 1fr 1fr 1fr; */
    padding: 0.3rem;
    gap: 0.3rem;
}

.ph span,
.electrical span,
.temperature span,
.cod span,
.turbidity span {
    grid-row-start: 1;
    grid-column-start: 1;
    grid-row-end: 3;
    grid-column-end: 2;
    align-self: center;
}

.ph .material-symbols-outlined,
.electrical .material-symbols-outlined,
.temperature .material-symbols-outlined,
.cod .material-symbols-outlined,
.turbidity .material-symbols-outlined {
    color: #dfdfdf;
    font-size: 5rem;
}

.ph h4,
.electrical h4,
.temperature h4,
.cod h4,
.turbidity h4 {
    font-size: 1.8rem;
    color: #dfdfdf;
    align-self: end;
    text-align: center;
}

.s01 .ph h5 {
    font-size: 1.8rem;
    color: v-bind('s01.ph.color');
}

.s01 .electrical h5 {
    font-size: 1.8rem;
    color: v-bind('s01.ec.color');
}

.s01 .temperature h5 {
    font-size: 1.8rem;
    color: v-bind('s01.temp.color');
}

.s01 .cod h5 {
    font-size: 1.8rem;
    color: v-bind('s01.cod.color');
}

.s01 .turbidity h5 {
    font-size: 1.8rem;
    color: v-bind('s01.turbidity.color');
}

.s02 .ph h5 {
    font-size: 1.8rem;
    color: v-bind('s02.ph.color');
}

.s02 .electrical h5 {
    font-size: 1.8rem;
    color: v-bind('s02.ec.color');
}

.s02 .temperature h5 {
    font-size: 1.8rem;
    color: v-bind('s02.temp.color');
}

.s02 .cod h5 {
    font-size: 1.8rem;
    color: v-bind('s02.cod.color');
}

.s02 .turbidity h5 {
    font-size: 1.8rem;
    color: v-bind('s02.turbidity.color');
}

.alarm {
    height: auto;
    grid-column-start: 1;
    grid-column-end: 4;
    background-color: rgb(26 26 26 / 30%);
    border-radius: 1.8rem;
    padding: 0.6rem 1.2rem;
    margin-top: 1rem;
}

.alarm .content {
    display: grid;
    gap: 1rem;
    grid-template-columns: 0.3fr 1fr;
    width: 100%;
    height: 100%;
}

.alarm h1 {
    color: #E5E5E5;
    font-size: 4rem;
    text-align: center;
    /* line-height: 1rem; */
    white-space: nowrap;
}

.alarm h2 {
    color: #E5E5E5;
    font-size: 2.9rem;
    text-align: center;
    /* line-height: 4rem; */
}

table {
    width: 100%;
    height: 95%;
    border-collapse: collapse;
}

table thead tr {
    height: 0.15rem;

}

table th {
    font-size: 1.5rem;
    color: #E5E5E5;
    vertical-align: top;
}

table tbody {
    background: rgb(255 255 255 / 70%);
}

table tbody td {
    border-bottom: 0.06rem solid #989898;
    text-align: center;
    color: rgb(187, 23, 23);
    font-weight: 600;
}

table tbody tr:first-child td:first-child {
    border-top-left-radius: 1.25rem;
}

table tbody tr:first-child td:last-child {
    border-top-right-radius: 1.25rem;
}

table tbody tr:last-child td:first-child {
    border-bottom-left-radius: 1.25rem;
}

table tbody tr:last-child td:last-child {
    border-bottom-right-radius: 1.25rem;
}

table tbody tr:last-child td {
    border: none;
}

@media (max-width: 1024px) {
    .layout {
        grid-template-columns: 1fr 1fr;
        grid-template-rows: repeat(2, 1fr);
    }

    .gis {
        grid-column-start: 1;
        grid-column-end: 3;
        grid-row-start: 2;
        grid-row-end: 3;
    }

    .alarm .content {
        grid-template-columns: 1fr;
        grid-template-rows: 0.3fr 1fr;
    }

    .alarm .title {
        margin-bottom: 1rem;
    }
}

@media (max-width: 720px) {
    .layout {
        grid-template-columns: 1fr;
        /* grid-template-rows: repeat(2, 1fr); */
    }

    .gis {
        display: none;
    }

    .alarm {
        display: none;
    }
}

@media (max-width: 320px) {}
</style>