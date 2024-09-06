<template>
    <div class="wrapper">
        <div class="login">
            <img src="../assets/img/Logo.png" alt="">
            <p>OverSeas IOT Monitor</p>
            <div>
                <input v-model="userName" type="text" @keyup.enter="login" placeholder="Account" />
            </div>
            <div>
                <input v-model="password" type="password" @keyup.enter="login" placeholder="Password" />
            </div>
            <div>
                <input v-model="captcha" type="text" @keyup.enter="login" placeholder="Captcha" />
            </div>
            <img class="captchaImg" :src="captchaImg" />
            <button class="pan" type="button" @click="login()">
                Login
            </button>
        </div>
    </div>
</template>

<script setup>
// 從 Vue 庫中引入 ref 和 onMounted 函數
import { ref, onMounted } from 'vue';
// 從 axios 庫中引入 axios 物件，用於發送 HTTP 請求
import axios from 'axios';
// 從路由中引入 router，用於導航
import router from '@/router';

// 定義 userName 的 ref，用於存放用戶名
const userName = ref()
// 定義 password 的 ref，用於存放密碼
const password = ref()
// 定義 captcha 的 ref，用於存放驗證碼
const captcha = ref('')
// 定義 captcha_id 的 ref，用於存放驗證碼 ID
const captcha_id = ref()
// 定義 captchaImg 的 ref，用於存放驗證碼圖片
const captchaImg = ref('')
// const store = useStore(); （此行被註解掉，可能是因為沒有用到 Vuex）

// 定義 login 函數，用於處理登入操作
const login = () => {
    // 發送 POST 請求到指定 URL，傳遞用戶名、密碼、驗證碼和驗證碼 ID
    axios({
        method: "POST",
        url: "/api/v1/Login",
        data: {
            username: userName.value,
            password: password.value,
            captcha: captcha.value,
            captcha_id: captcha_id.value,
        },
    }).then((response) => {
        // 請求成功後處理返回的數據
        let token = response.data.data;
        // 將 token 存入 sessionStorage
        sessionStorage.setItem("token", token)
        // 彈出提示訊息
        alert(response.data.message)
        // 導航至指定路由 '/WaterMonitoring'
        router.push('/WaterMonitoring')
    }).catch(function (error) {
        // 請求處理失敗時的處理
        // 輸出錯誤訊息
        console.log(error);
        // 彈出錯誤訊息
        alert(error.data.message)
        // 重新載入當前頁面
        router.go(0)
    });
};

// 組件掛載完成後執行的函數
onMounted(async () => {
    // 獲取驗證碼
    await axios
        .get("/api/v1/Captcha")
        .then((response) => {
            // 請求成功後處理返回的數據
            console.log(response);
            // 將返回的驗證碼 ID 賦值給 captcha_id
            captcha_id.value = response.data.data.captcha_id
            // 將返回的驗證碼圖片資料賦值給 captchaImg
            captchaImg.value = response.data.data.captcha_data
        })
        .catch(function (error) {
            // 請求處理失敗時的處理
            console.log(error);
        });
})
</script>

<style scoped>
.wrapper {
    height: 100vh;
    /* background: rgb(255 255 255 / 60%); */
}

.login {
    width: 18rem;
    height: 25rem;
    padding: 1.3rem;
    text-align: center;
    background: rgb(255 255 255 / 60%);
    margin: 0 auto;
    position: relative;
    top: 25%;
    border-radius: 1.5rem;
}

.login p {
    font-size: 1.3rem;
    color: #0071BB;
}

input {
    width: 100%;
    height: 2.25rem;
    margin-top: 1rem;
    background: rgb(0 0 0 / 20%);
    text-align: center;
    border-radius: 0.5rem;
}

::placeholder {
    font-size: 1rem;
    text-align: center;
    color: #666666;
}

.captchaImg {
    width: 100%;
    margin: 0.5rem 0;
}


button {
    width: 100%;
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
    border-radius: 0.5rem;
    overflow: hidden;
    color: white;
}

.pan::before {
    content: "";
    background: #000000;
    transition: transform 0.3s cubic-bezier(0.7, 0, 0.2, 1);
    z-index: -1;
}

.pan:hover::before {
    transform: translate3d(0, -100%, 0);
}
</style>