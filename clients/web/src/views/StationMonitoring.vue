<template>
    <div class="layout">
        <div class="s01">
            <div class="title">
                <p>Monitoring Point Location {{ s01.id }}</p>
            </div>
            <div class="location">
                <img src="@/assets/img/Vector.png" width="">
                <h6>{{ s01.coordinate }}</h6>
            </div>
            <div class="content">
                <div class="station_img">
                    <img src="@/assets/img/FCM.png">
                </div>
                <div class="date_time">
                    <DateTimer :date_time="s01.update_time" />
                </div>
                <div class="battery_temp">
                    <img src="../assets/img/BatteryTemperature.png" alt="">
                    <h4>{{ s01.battery_temp.full_name }}</h4>
                    <h3>{{ s01.battery_temp.val + '\xa0' }}°C</h3>
                    <ProgressBar :value="s01.battery_temp.percent" :max="100" />
                </div>
                <div class="battery_value">
                    <img src="../assets/img/BatteryCharging80.png" alt="">
                    <h4>{{ s01.battery.full_name }}</h4>
                    <h3>{{ s01.battery.val + '\xa0' }}%</h3>
                    <ProgressBar :value="s01.battery.percent" :max="100" />
                </div>

                <div class="signal_strength">
                    <img src="../assets/img/SignalStrength.png" alt="">
                    <h4>{{ s01.signal_strength.full_name }}</h4>
                    <h3>{{ s01.signal_strength.strength }}</h3>
                    <ProgressBar :value="s01.signal_strength.percent" :max="100" />
                </div>
                <div class="fcm_temp">
                    <img src="../assets/img/BatteryTemperature.png" alt="">
                    <h4>{{ s01.fcm_temp.full_name }}</h4>
                    <h3>{{ s01.fcm_temp.val + '\xa0' }}°C</h3>
                    <ProgressBar :value="s01.fcm_temp.percent" :max="100" />
                </div>

            </div>
        </div>
        <div class="s02">
            <div class="title">
                <p>Monitoring Point Location {{ s02.id }}</p>
            </div>
            <div class="location">
                <img src="@/assets/img/Vector.png">
                <h6>{{ s02.coordinate }}</h6>
            </div>
            <div class="content">
                <div class="station_img">
                    <img src="@/assets/img/FCM.png">
                </div>
                <div class="date_time">
                    <DateTimer :date_time="s02.update_time" />
                </div>
                <div class="battery_temp">
                    <img src="../assets/img/BatteryTemperature.png" alt="">
                    <h4>{{ s02.battery_temp.full_name }}</h4>
                    <h3>{{ s02.battery_temp.val + '\xa0' }}°C</h3>
                    <ProgressBar :value="s02.battery_temp.percent" :max="100" />
                </div>

                <div class="battery_value">
                    <img src="../assets/img/BatteryCharging80.png" alt="">
                    <h4>{{ s02.battery.full_name }}</h4>
                    <h3>{{ s02.battery.val + '\xa0' }}%</h3>
                    <ProgressBar :value="s02.battery.percent" :max="100" />
                </div>

                <div class="signal_strength">
                    <img src="../assets/img/SignalStrength.png" alt="">
                    <h4>{{ s02.signal_strength.full_name }}</h4>
                    <h3>{{ s02.signal_strength.strength }}</h3>
                    <ProgressBar :value="s02.signal_strength.percent" :max="100" />
                </div>

                <div class="fcm_temp">
                    <img src="../assets/img/BatteryTemperature.png" alt="">
                    <h4>{{ s02.fcm_temp.full_name }}</h4>
                    <h3>{{ s02.fcm_temp.val + '\xa0' }}°C</h3>
                    <ProgressBar :value="s02.fcm_temp.percent" :max="100" />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';
import ProgressBar from "@/components/ProgressBar.vue";
import DateTimer from "@/components/DateTimerVue.vue";
import axios from 'axios'
// 引入 convertLatLngToDMS 函數，用於將經緯度轉換為度分秒表示法
import { convertLatLngToDMS } from '@/assets/js/convertLatLngToDMS';
const s01 = ref({
    id: '',
    update_time: '',
    coordinate: { lng: 55.3551, lat: 25.1976 },
    battery_temp: {
        full_name: '',
        val: 0,
        percent: 0,
    },
    battery: {
        full_name: '',
        val: 0,
        percent: 0,
    },
    signal_strength: {
        full_name: '',
        val: 0,
        percent: 0,
        strength: '',
    },
    fcm_temp: {
        full_name: '',
        val: 0,
        percent: 0,
    },
})
const s02 = ref({

    id: '',
    update_time: '',
    coordinate: { lng: 46.674444, lat: 24.711389 },
    battery_temp: {
        full_name: '',
        val: 0,
        percent: 0,
    },
    battery: {
        full_name: '',
        val: 0,
        percent: 0,
    },
    signal_strength: {
        full_name: '',
        val: 0,
        percent: 0,
        strength: '良好',
    },
    fcm_temp: {
        full_name: '',
        val: 0,
        percent: 0,
    },
})
// 定義一個捕獲數字的函數
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
// 定義一個獲取最新數據的函數
const getLastData = async () => {
    await axios
        .get("/api/v1/Last")
        .then((response) => {
            response.data.data.last.forEach(element => {
                switch (element.device.uuid) {
                    //FCM01
                    case 'f52132d2-8a8f-4acc-afbd-6ec556b12182':
                        s01.value.id = element.station.id
                        s01.value.update_time = element.device.update_time;
                        s01.value.coordinate = convertLatLngToDMS(element.station.lat, element.station.lon)
                        element.physical_quantities.forEach(ele => {
                            switch (ele.uuid) {
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
                    //FCM02
                    case '82d7329d-c750-4d97-9f03-c5771e599163':
                        s02.value.id = element.station.id
                        s02.value.update_time = element.device.update_time;
                        s02.value.coordinate = convertLatLngToDMS(element.station.lat, element.station.lon)
                        element.physical_quantities.forEach(ele => {
                            switch (ele.uuid) {
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
        })
        .catch(function (error) {
            console.log(error);
        });
    //getAlarmUpperLower
    await axios({
        method: "Get",
        url: "/api/v1/Alarm/Setting/Station",
        params: {
            StationUUID: 'e2232881-e996-46be-a0a2-95a2983fdb42',
        }
    }).then((response) => {
        response.data.data.forEach(element => {
            switch (element.physical_quantity.uuid) {
                //battery_temp
                case "184c9a57-65bc-4271-97d6-b983431e511e": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "278c4213-86f3-4108-8b16-6efa50c8c1cd":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "1a1dbb62-37cf-4e5b-bbe4-aa8731fcf124":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.battery_temp.percent = ((s01.value.battery_temp.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //battery
                case "a03598ab-8b1a-48ec-8ecf-d1e7838127df": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "75b3a6ee-11a4-4474-911c-a624780b1cdf":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "f14f3e02-ea9c-4c38-87bd-16deedac53c5":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.battery.percent = ((s01.value.battery.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //signal_strength
                case "675a68b9-e5f2-4b25-9ddb-d73cfb4c2d6a": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "e054c9ba-f567-44d5-a98a-796b8195f47f":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "e9be06e5-1d38-4c77-be7c-05be21c37aab":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.signal_strength.percent = ((s01.value.signal_strength.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //FCM_temp
                case "5c5d94c5-d1e2-4519-81ed-41ea88def09a": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {

                        switch (ele.uuid) {
                            case "6c7f7d6c-2433-47a3-b5c9-d6b6531ff4d3":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "c4ce3d77-df41-4d20-85b5-b5afd048b536":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s01.value.fcm_temp.percent = ((s01.value.fcm_temp.val - low_threshold) / (high_threshold - low_threshold)) * 100
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
    //s02
    await axios({
        method: "Get",
        url: "/api/v1/Alarm/Setting/Station",
        params: {
            StationUUID: '11465818-37f3-4c81-a093-296ea9dee685',
        }
    }).then((response) => {
        response.data.data.forEach(element => {
            switch (element.physical_quantity.uuid) {
                //battery_temp
                case "8bf89996-f693-40c5-8995-77160ed08fd8": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "a8039623-b898-4a6f-aa82-8a9ba9c2b6ca":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "07b906a1-6d10-4aa0-9976-5dc62037b779":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.battery_temp.percent = ((s02.value.battery_temp.val - low_threshold) / (high_threshold - low_threshold)) * 100

                    break;
                }
                //battery
                case "ee45cc81-6375-448e-b940-9a433b771cf9": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "ecad9cd2-117b-4dbb-9855-7f667468d861":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "0c756f74-a9d2-4365-b9e5-baa02f11430d":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.battery.percent = ((s02.value.battery.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //signal_strength
                case "c9e0cee6-193e-4ba5-9e61-2edbf620b4cb": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "cc72a1ab-825f-4c1e-8b56-c308052a6efb":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "232609ce-4023-4208-9f33-3b74cfb76126":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.signal_strength.percent = ((s02.value.signal_strength.val - low_threshold) / (high_threshold - low_threshold)) * 100
                    break;
                }
                //FCM_temp
                case "995ec435-2c5f-41c6-9a16-c47c3b320364": {
                    let high_threshold
                    let low_threshold
                    element.alarm_settings.forEach(ele => {
                        switch (ele.uuid) {
                            case "d6e75edf-15ac-4b6c-b8b0-d940d93c54d6":
                                high_threshold = captureNumber(ele.boolean_expression)
                                break;
                            case "0e811b44-8978-4f40-8a93-58e3742d147a":
                                low_threshold = captureNumber(ele.boolean_expression)
                                break;
                            default:
                                break;
                        }
                    });
                    // 將原始數值轉換成百分比 ((原始數值 - min) / (max - min)) * 100;
                    s02.value.fcm_temp.percent = ((s02.value.fcm_temp.val - low_threshold) / (high_threshold - low_threshold)) * 100
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
    grid-template-columns: 1fr 1fr;
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
    margin: 0.8rem;
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

.title {
    display: flex;
    justify-content: space-around;
    align-items: center;
}

.s01 .content,
.s02 .content {
    display: grid;
    height: auto;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: repeat(4, 1fr);
}

.content .station_img {
    grid-row-start: 1;
    grid-column-start: 1;
    grid-row-end: 4;
    grid-column-end: 2;
    align-self: center;
}

.content .station_img img {
    width: 65%;
    height: auto;
}


.label h6 {
    font-size: 1.5rem;
    color: white;
    margin-top: 1.8rem;
    text-align: center;
}

.date_time {
    grid-row-start: 4;
    grid-column-start: 1;
    grid-row-end: 5;
    grid-column-end: 2;
    align-content: center;
}

.battery_temp,
.fcm_temp,
.battery_value,
.signal_strength {
    display: grid;
    grid-template-columns: 0.6fr 1fr;
    /* grid-template-rows: 30% 30% 40%; */
    padding: 0.4rem;
    gap: 0.3rem;
}


.battery_temp img,
.fcm_temp img,
.battery_value img,
.signal_strength img {
    grid-row-start: 1;
    grid-column-start: 1;
    grid-row-end: 3;
    grid-column-end: 2;
    align-self: center;
    margin: 0 auto;
}

.battery_temp span,
.fcm_temp span,
.battery_value span,
.signal_strength span {
    grid-row-start: 1;
    grid-column-start: 1;
    grid-row-end: 3;
    grid-column-end: 2;
    align-self: center;
}

.battery_temp .material-symbols-outlined,
.fcm_temp .material-symbols-outlined,
.battery_value .material-symbols-outlined,
.signal_strength .material-symbols-outlined {
    color: #dfdfdf;
    font-size: 5rem;
}

.battery_temp h4,
.fcm_temp h4,
.battery_value h4,
.signal_strength h4 {
    font-size: 1.7rem;
    color: #dfdfdf;
    align-self: end;
    text-align: center;
}

.battery_temp h3,
.fcm_temp h3,
.battery_value h3,
.signal_strength h3 {
    font-size: 2rem;
    color: #dfdfdf;
    align-self: start;
    text-align: center;
}

@media (max-width: 1024px) {

    .s01 .content,
    .s02 .content {
        grid-template-columns: 1fr;
    }
}

@media (max-width: 720px) {
    .layout {
        grid-template-columns: 1fr;
    }
}

@media (max-width: 320px) {}
</style>