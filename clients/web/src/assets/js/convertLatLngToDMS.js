// 定義一個函數 convertToDMS 用於將十進制度數轉換為度分秒（DMS）格式
const convertToDMS = (deg) => {
    // 取絕對值以處理負數情況
    const absDeg = Math.abs(deg);
    // 取整數部分作為度數
    const degrees = Math.floor(absDeg);
    // 計算餘下的小數部分轉換為分鐘數
    const minutes = Math.floor((absDeg - degrees) * 60);
    // 計算餘下的小數部分轉換為秒數，並保留一位小數
    const seconds = ((absDeg - degrees - (minutes / 60)) * 3600).toFixed(1);
    // 返回度分秒格式的字符串
    return degrees + "°" + minutes + "'" + seconds + '"';
}

// 定義並導出一個函數 convertLatLngToDMS 用於將緯度和經度轉換為度分秒（DMS）格式
export const convertLatLngToDMS = (latitude, longitude) => {
    // 使用 convertToDMS 將緯度轉換為 DMS 格式
    const latDMS = convertToDMS(latitude);
    // 使用 convertToDMS 將經度轉換為 DMS 格式
    const longDMS = convertToDMS(longitude);
    // 判斷緯度的方向是北（N）還是南（S）
    const latDirection = latitude >= 0 ? "N" : "S";
    // 判斷經度的方向是東（E）還是西（W）
    const longDirection = longitude >= 0 ? "E" : "W";
    // 返回轉換後的緯度和經度的 DMS 格式字符串，包含方向
    return latDMS + " " + latDirection + " " + longDMS + " " + longDirection;
}