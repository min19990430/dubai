<template>
    <div class="button_list">
        <select class="pan" @change="getCalibrationByDeviceUUID" v-model="device">
            <option v-for="(item, index) in devices" :key="index" :value="item">{{ item.id }}【 {{ item.station.name }} 】
            </option>
        </select>
    </div>

    <div class="table-wrapper">
        <table class="fl-table" :style="{ border: getBorderStyle(item.physical_quantity.status_code) }"
            v-for="(item, index) in data.calibration" :key="index">
            <thead>
                <tr>
                    <th class="status">
                        <select class="pan"
                            :style="{ backgroundColor: getBackgroundColor(item.physical_quantity.status_code) }"
                            @change="patchPhysicalQuantityStatus(item)" v-model="item.physical_quantity.status_code">
                            <option class="normal" value="10" selected>Normal</option>
                            <option class="calibrate" value="20">Calibrate </option>
                        </select>
                    </th>
                    <th class="title">
                        <span>
                            {{ item.full_name }}
                        </span>
                    </th>
                </tr>
                <tr>
                    <th> Point </th>
                    <th> Actual Value </th>
                    <th> Set Value </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(item_parameter, index_parameter) in item.calibration_parameter" :key="index_parameter">
                    <td>{{ index_parameter + 1 }}</td>
                    <td>
                        <input type="text" v-model.trim="item.calibration_parameter[index_parameter]" />
                    </td>
                    <td>
                        <input type="text" v-model.trim="item.calibration_value[index_parameter]" />
                    </td>
                </tr>
                <tr>
                    <td>
                        <div>
                            <button class="update" @click="patchCalibration(
            item.physical_quantity.uuid,
            item.calibration_parameter,
            item.calibration_value
        )">Update</button>
                        </div>
                    </td>

                    <td class=" data">
                        <span>{{ item.data }}</span>
                    </td>
                    <td class="value">
                        <span>{{ item.value }}</span>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, } from 'vue';
import axios from 'axios'

const devices = ref([])
const device = ref("")
const data = ref({})
const previousStatusCodes = ref({});

// 根據狀態碼獲取背景顏色
const getBackgroundColor = (statusCode) => {
    switch (statusCode) {
        case '10':
            return '#70ad47';
        case '20':
            return '#efbd09';
        default:
            return '#bebebe';
    }
};
// 根據狀態碼獲取邊框樣式
const getBorderStyle = (statusCode) => {
    switch (statusCode) {
        case '20':
            return '0.2rem solid #efbd09';
        default:
            return 'none';
    }
};
// 獲取所有設備站點數據
const getAllDeviceStation = async () => {
    await axios({
        method: "Get",
        url: "/api/v1/Device/Station"
    }).then((response) => {
        console.log(response.data.data);
        response.data.data.forEach(element => {
            devices.value.push(element)
        });
        device.value = response.data.data[1]
    })
        .catch(function (error) {
            console.log(error);
        });
}
// 根據設備 UUID 獲取校準數據
const getCalibrationByDeviceUUID = async () => {
    await axios({
        method: "Get",
        url: "/api/v1/Calibration",
        params: {
            DeviceUUID: device.value.uuid,
        }
    }).then((response) => {
        response.data.data.calibration = response.data.data.calibration.filter(val => {
            return (
                val.physical_quantity.source == "sensor"
            )
        })
        response.data.data.calibration.forEach((element) => {
            if (element.physical_quantity.status_code == "11") {
                element.physical_quantity.status_code = "10"
            }
            element.value = element.value.toFixed(2)
            element.calibration_parameter = element.calibration_parameter.split(',')
            element.calibration_value = element.calibration_value.split(',')
        });
        data.value = response.data.data

        // 校準數據並初始化
        response.data.data.calibration.forEach(item => {
            previousStatusCodes.value[item.physical_quantity.uuid] = item.physical_quantity.status_code;
        });
    }).catch(function (error) {
        console.log(error);
    });
}
// 根據設備 UUID 刷新校準數據
const RefreshCalibrationByDeviceUUID = async () => {
    await axios({
        method: "Get",
        url: "/api/v1/Calibration",
        params: {
            DeviceUUID: device.value.uuid,
        }
    }).then((response) => {
        response.data.data.calibration = response.data.data.calibration.filter(val => {
            return (
                val.physical_quantity.source == "sensor"
            )
        })
        response.data.data.calibration.forEach((element, index) => {
            data.value.calibration[index].value = element.value.toFixed(2)
            data.value.calibration[index].data = element.data
        });
    }).catch(function (error) {
        console.log(error);
    });
}
// 更新校準數據
const patchCalibration = async (uuid, parameter, value) => {
    if (confirm("Confirm the changes?") == true) {
        await axios({
            method: "PATCH",
            url: "/api/v1/Calibration",
            data: {
                uuid: uuid,
                calibration_enable: true,
                calibration_parameter: parameter.toString(),
                calibration_value: value.toString(),
            }
        }).then((response) => {
            alert(response.data.message)
        }).catch(function (error) {
            console.log(error);
        });
    }
}
// 更新物理量狀態碼
const patchPhysicalQuantityStatus = async (item) => {
    const { uuid, status_code } = item.physical_quantity
    if (!confirm("Confirm the changes Status?")) {
        item.physical_quantity.status_code = previousStatusCodes.value[uuid]
    } else {
        await axios({
            method: "PATCH",
            url: "/api/v1/PhysicalQuantity/Status",
            data: {
                uuid: uuid,
                status_code: status_code,
            }
        }).then((response) => {
            console.log(response);
            previousStatusCodes.value[uuid] = status_code;
        }).catch(function (error) {
            console.log(error);
        });
    }
}
// 定義校準計時器
let calibrationTimer = null

onMounted(async () => {
    await getAllDeviceStation()
    await getCalibrationByDeviceUUID()
    calibrationTimer = setInterval(RefreshCalibrationByDeviceUUID, 1000 * 5)
})

onBeforeUnmount(() => {
    clearInterval(calibrationTimer)
    calibrationTimer = null
})

</script>

<style scoped>
.button_list {
    text-align: center;
}

select {
    pointer-events: auto;
    cursor: pointer;
    background: #bebebe;
    border: none;
    padding: 0.5rem 1.0rem;
    margin: 0;
    font-family: inherit;
    font-size: 1rem;
    position: relative;
    display: inline-block;
    z-index: 1;
}

option {
    text-align: center
}

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

.normal {
    color: white;
    background: #70ad47;
}

.calibrate {
    color: white;
    background: #efbd09;
}

.pan {
    font-family: aktiv-grotesk-extended, sans-serif;
    font-weight: 700;
    border-radius: 3rem;
    overflow: hidden;
    color: black;
}

.update {
    font-family: aktiv-grotesk-extended, sans-serif;
    font-weight: 700;
    border-radius: 0.5rem;
    border: 0.2rem solid rgb(8, 61, 124);
    color: rgb(8, 61, 124);
    background: rgb(242, 242, 242);
    font-size: 1rem;
}

.pan:hover {
    color: white;
}

.pan::before {
    content: "";
    /* background: #BEBEBE; */
    transition: transform 0.3s cubic-bezier(0.7, 0, 0.2, 1);
    z-index: -1;
}

.pan:hover::before {
    transform: translate3d(0, -100%, 0);
}

.table-wrapper {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1rem;
    box-shadow: 0px 35px 50px rgba(0, 0, 0, 0.2);
}

input[type=checkbox] {
    position: relative;
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    width: 3rem;
    height: calc(3rem / 2);
    background-color: #ddd;
    border-radius: calc(3rem / 2);
    outline: none;
    transition: background 450ms ease;
    box-shadow: 0 0 0 3px #cacaca;
}

input[type=checkbox]:before,
input[type=checkbox]:after {
    position: absolute;
    display: block;
    content: "";
    border-radius: 100%;
    transition: background 450ms ease, transform 450ms ease;
}

input[type=checkbox]:before {
    width: calc(3rem / 2);
    height: calc(3rem / 2);
    background-color: black;
}

input[type=checkbox]:checked:before {
    background-color: #43A047;
    transform: translateX(100%);
}

table {
    border: 0.2rem solid #efbd09;
    margin-top: 1rem;
    background: #fff;
    width: 100%;
    border-radius: 0.8rem;
    padding: 1.8rem;
    text-align: center;
    box-shadow: 0 2rem 3rem rgba(132, 139, 200, 0.18);
    transition: all 300ms ease;
    table-layout: fixed
}

.title {
    /* padding-bottom: 2rem; */
    font-size: 1.5rem;
}

.status {
    font-size: 1.2rem;
}

table thead th {
    font-size: 1.3rem;
    color: #005AB5;
    width: 33%;
}

table thead th span {
    color: #363949;
    /* margin-right: 1rem; */
}

table tbody td {
    font-size: 1.2rem;
    font-weight: bold;
    height: 3.4rem;
    border-bottom: 1px solid rgba(132, 139, 200, 0.18);
    color: #677483;
}

table tbody tr .value span,
table tbody tr .data span {
    color: #111e88;
}

table tbody td input {
    font-size: 1.2rem;
    font-weight: bold;
    text-align: center;
    color: rgb(187, 23, 23);
    width: 100%;
}

table tbody tr:last-child td {
    border: none;
}


/* Responsive */
@media (max-width: 1024px) {
    .table-wrapper {
        grid-template-columns: repeat(2, 1fr);
        gap: 0.5rem;
    }

    .update {
        font-size: 0.9rem;
    }
}

@media (max-width: 720px) {
    .table-wrapper {
        grid-template-columns: 1fr;
    }

    .update {
        font-size: 0.8rem;
    }
}

@media (max-width: 320px) {}
</style>