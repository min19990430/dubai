import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
axios.defaults.baseURL = process.env.VUE_APP_API_PATH; //配置路由請求路徑
axios.defaults.withCredentials = true; // axios請求開啟cookie，支持跨域請求攜帶cookie
axios.defaults.timeout = 10000; // 請求超時時間 當請求時間超過5秒還未取得結果時 提示用戶請求超時

axios.interceptors.response.use(
    response => {
        // 如果返回的狀態碼為200，說明接口請求成功，可以正常拿到數據
        // 否則的話拋出錯誤
        if (response.status === 200) {
            return Promise.resolve(response);
        } else {
            return Promise.reject(response);
        }
    },
    error => {
        if (error.response.status) {
            switch (error.response.status) {
                // 401: 未登錄
                // 未登錄則跳轉登錄頁面，並攜帶當前頁面的路徑
                // 在登錄成功後返回當前頁面，這一步需要在登錄頁操作。
                case 401:
                    console.log("401")
                    router.replace({
                        path: '/login',
                        query: {
                            redirect: router.currentRoute.fullPath
                        }
                    });
                    break;
                case 503:
                    console.log("503")
                    router.replace({
                        path: '/login',
                        query: {
                            redirect: router.currentRoute.fullPath
                        }
                    });
                    break;
                default:
            }
            return Promise.reject(error.response);
        }
    }
)

createApp(App).use(store).use(router).mount('#app')
