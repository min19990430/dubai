<template>
    <div id="clock">
        <p class="date">{{ date }}</p>
        <p class="time">{{ time }}</p>
    </div>
</template>

<script setup>
import { ref, watch, defineProps } from 'vue';
// 使用 defineProps 定義接收的屬性
// 定義一個名為 date_time 的屬性，其類型為 String
const props = defineProps({
    date_time: String,
})
// 定義一個響應式變量 date，初始值為 "2024-03-07 THU"
const date = ref("2024-03-07 THU")
// 定義一個響應式變量 time，初始值為 "16:33:07"
const time = ref("16:33:07")
// 定義一個陣列 week，包含一周七天的縮寫
const week = ['SUN', 'MON', 'TUE', 'WED', 'THU', 'FRI', 'SAT'];

// 使用 watch 監聽 props.date_time 的變化
watch(
    () => props.date_time, (newVal) => {  // 當監聽的值發生變化時執行的回調函數
        // 將新值轉換為 Date 對象
        var cd = new Date(newVal);
        // 更新 date 變量的值，格式為 "YYYY-MM-DD DDD"
        date.value = zeroPadding(cd.getFullYear(), 4) + '-' + zeroPadding(cd.getMonth() + 1, 2) + '-' + zeroPadding(cd.getDate(), 2) + ' ' + week[cd.getDay()];
        // 更新 time 變量的值，格式為 "HH:MM:SS"
        time.value = zeroPadding(cd.getHours(), 2) + ':' + zeroPadding(cd.getMinutes(), 2) + ':' + zeroPadding(cd.getSeconds(), 2);
    }
)
// 定義一個函數 zeroPadding，用於在數字前補零，使其達到指定的位數
const zeroPadding = (num, digit) => {
    // 定義一個空字串 zero
    var zero = '';
    // 使用迴圈生成指定位數的零
    for (var i = 0; i < digit; i++) {
        zero += '0';
    }
    // 將零和數字拼接後，取其最後的指定位數，返回結果
    return (zero + num).slice(-digit);
}
</script>

<style scoped>
#clock {
    font-family: "Share Tech Mono", monospace;
    color: #ffffff;
    text-align: center;
    color: #daf6ff;
    text-shadow: 0 0 20px #0aafe6, 0 0 20px rgba(10, 175, 230, 0);
}

#clock .time {
    letter-spacing: 0.05em;
    font-size: 4rem;
}

#clock .date {
    letter-spacing: 0.1em;
    font-size: 1.5rem;
}
</style>