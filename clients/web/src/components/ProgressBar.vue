<template>
    <div class="progress-bar__wrapper">
        <progress :value="props.value" :max='props.max'></progress>
    </div>
</template>

<script setup>
import { defineProps, watch } from 'vue';
import { ref } from 'vue';
// 定義兩個屬性，value 和 max，類型均為 Number
const props = defineProps({
    value: Number,
    max: Number,
})
// 定義一個響應式變量 bar_color 用於存儲進度條的顏色
const bar_color = ref()
// 使用 watch 監聽 props.value 的變化
watch(
    () => props.value,// 監聽的值
    (newVal, oldVal) => {
        // 檢查 newVal 是否不同於 oldVal，如果不同則進行顏色更新
        switch (newVal != oldVal) {
            // 根據 newVal 的範圍設置 bar_color 的值
            case (newVal <= 20):
                bar_color.value = 'rgb(189, 0, 2)'; //紅色
                break;
            case (newVal > 20 && newVal <= 40):
                bar_color.value = 'rgb(249, 79, 0)'; //橘紅色
                break;
            case (newVal > 40 && newVal <= 60):
                bar_color.value = 'rgb(247, 161, 15)'; //橘色
                break;
            case (newVal > 60 && newVal <= 80):
                bar_color.value = 'rgb(226, 229, 25)'; //黃色
                break;
            case (newVal > 80 && newVal <= 100):
                bar_color.value = 'rgb(46, 228, 44)'; //青色
                break;
            default:
                // 在這裡處理未匹配的情況
                break;
        }
    }
)
</script>

<style scoped>
.progress-bar__wrapper {
    grid-row-start: 3;
    grid-column-start: 1;
    grid-row-end: 4;
    grid-column-end: 3;
}

progress {
    width: 90%;
    height: 0.625rem;
    border-radius: 6.25rem;
    background-color: rgba(255, 255, 255, 0.1);
    transition: width 1000ms ease;
}

progress[value]::-webkit-progress-bar {
    width: 100%;
    height: 0.625rem;
    border-radius: 6.25rem;
    background-color: rgba(255, 255, 255, 0.1);
    transition: width 1000ms ease;
}

progress[value]::-webkit-progress-value {
    width: 0;
    border-radius: 6.25rem;
    /* background-color: v-bind('bar_color'); */
    background: linear-gradient(to right, #f63a0f, #86e01e);
    transition: width 1000ms ease;
}
</style>