import { createRouter, createWebHistory } from 'vue-router'
const routes = [
  {
    path: '/login',
    name: 'LogIn',
    component: () => import('@/views/LogIn.vue'),
  },
  {
    path: '/WaterMonitoring',
    name: 'WaterMonitoring',
    component: () => import('@/views/WaterMonitoring.vue'),
  },
  {
    path: '/StationMonitoring',
    name: 'StationMonitoring',
    component: () => import('@/views/StationMonitoring.vue'),
  },
  {
    path: '/Historical',
    name: 'Historical',
    component: () => import('@/views/HistoricalData.vue'),
  },
  {
    path: '/TrendChart',
    name: 'TrendChart',
    component: () => import('@/views/TrendChart.vue'),
  },
  {
    path: '/Alarm',
    name: 'Alarm',
    component: () => import('@/views/AlarmRecord.vue'),
  },
  {
    path: '/Calibrate',
    name: 'Calibrate',
    component: () => import('@/views/CalibrationSet.vue'),
  },
  {
    path: '/Download',
    name: 'Download',
    component: () => import('@/views/DownloadFile.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  // 配置router active的類名
  // 每個router都會經過/所以會被激活，通過exact可以精準選擇所在路由
  // linkActiveClass: '', // 模糊比對
  linkExactActiveClass: 'active', // 準確比對
  routes
})

router.beforeEach((to) => {
  // to：使用者要跳轉的路由
  // from：使用者前一個訪問的路由
  // 回傳 false 取消跳轉，true / undefined（預設）容許跳轉
  // next 參數在 Vue Router 4 並非必須

  const token = sessionStorage.getItem("token");

  if (to.fullPath == '/login') return; // 登入頁不用驗證

  if (token === null) {
    router.push('/login')
  }
})

export default router
