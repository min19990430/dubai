<template>
    <nav>
        <router-link to="/WaterMonitoring">
            <span class="material-symbols-outlined"> water_drop </span>
            <h3>Water</h3>
        </router-link>
        <router-link to="/StationMonitoring">
            <span class="material-symbols-outlined"> cast </span>
            <h3>Station</h3>
        </router-link>
        <router-link to="/Historical">
            <span class="material-symbols-outlined"> quick_reference_all </span>
            <h3>Historical</h3>
        </router-link>
        <router-link to="/TrendChart">
            <span class="material-symbols-outlined"> insights </span>
            <h3>Trend</h3>
        </router-link>
        <router-link to="/Alarm">
            <span class="material-symbols-outlined"> release_alert </span>
            <h3>Alarm</h3>
        </router-link>
        <router-link to="/Calibrate">
            <span class="material-symbols-outlined"> settings </span>
            <h3>Calibrate</h3>
        </router-link>
        <router-link to="/Download">
            <span class="material-symbols-outlined"> file_save </span>
            <h3>Download</h3>
        </router-link>
        <button @click="logout">
            <span class="material-symbols-outlined"> start </span>
            <h3>Logout</h3>
        </button>
    </nav>
</template>

<script setup>
import axios from 'axios'
import router from '@/router';

const logout = () => {
    let access_token = sessionStorage.getItem('token')
    axios({
        method: "POST",
        url: "/api/v1/Logout",
        headers: {
            'Authorization': `Bearer ${access_token}`
        },
    }).then((response) => {
        console.log(response)
        sessionStorage.removeItem('token')
        router.push("/login")
    }).catch(function (error) {
        // 請求處理失敗
        console.log(error.response);
    });
}
</script>

<style scoped>
nav {
    position: sticky;
    top: 0;
    display: flex;
    z-index: 1;
    /* flex-wrap: wrap; */
    /* 允許換行 */
    height: 3.5rem;
    line-height: 3.5rem;
    background: #FFF;
    justify-content: center;
    /* margin-bottom: 1rem; */
}

nav a {
    display: flex;
    align-items: center;
    margin-right: 3rem;
    color: #024D81;
    white-space: nowrap;
    background: inherit;
    cursor: pointer;
}

nav button {
    display: flex;
    align-items: center;
    /* margin-right: 3rem; */
    color: #024D81;
    white-space: nowrap;
    background: inherit;
    cursor: pointer;
}

.material-symbols-outlined {
    line-height: 0;
    margin-right: 0.3rem;
}

/* 桌面端樣式 */
@media (min-width: 1025px) {}

/* 平板端樣式 */
@media (max-width: 1024px) {
    h3 {
        display: none;
    }
}

/* 手機端樣式 */
@media (max-width: 600px) {
    nav a {
        margin-right: 1rem;
    }
}
</style>