export const formatDateTime = (dateTime) => {
    let date = new Date(dateTime)
    let year = date.getFullYear();
    let month = ("0" + (date.getMonth() + 1)).slice(-2); // Months are zero based
    let day = ("0" + date.getDate()).slice(-2);
    let hours = ("0" + date.getHours()).slice(-2);
    let minutes = ("0" + date.getMinutes()).slice(-2);
    let seconds = ("0" + date.getSeconds()).slice(-2);
    let formattedDate = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedDate
}

export const extractTimeIfNotZero = (dateTime) => {
    // 將日期時間字串轉換為 Date 物件
    const date = new Date(dateTime);

    // 獲取時、分、秒
    let year = date.getFullYear();
    let month = ("0" + (date.getMonth() + 1)).slice(-2); // Months are zero based
    let day = ("0" + date.getDate()).slice(-2);
    let hours = ("0" + date.getHours()).slice(-2);
    let minutes = ("0" + date.getMinutes()).slice(-2);
    let seconds = ("0" + date.getSeconds()).slice(-2);

    // 檢查是否為 00:00:00
    if (hours === '00' && minutes === '00' && seconds === '00') {
        return `${year}-${month}-${day} 00:00`;
    } else {
        return `${hours}:${minutes}`;
    }
}

export const formatToLocalDateTime = (dateString) => {
    let isoDateTime = new Date(dateString).toISOString()
    let date = new Date(dateString)
    //取得時區偏移量（以分鐘為單位）
    const offsetMinutes = date.getTimezoneOffset();
    // 將偏移量轉換為小時和分鐘
    const offsetHours = Math.abs(offsetMinutes / 60);
    const offsetSign = offsetMinutes > 0 ? "-" : "+";
    // 格式化偏移量為符合ISO 8601的字串格式
    const offsetString = `${offsetSign}${String(offsetHours).padStart(2, "0")}:00`;
    // 建立新的日期字串，加入時區偏移量
    const newTime = isoDateTime.replace("Z", `${offsetString}`);
    return newTime
}
export const formatStringOffsetToISO8601 = (dateTime) => {
    // 將字串轉換為時間格式
    var date = new Date(dateTime);
    // 調整時區
    var timezoneOffset = 8; // 台灣時區 GMT+8
    //傳回當地時區與UTC時間的時區差，傳回的值是以分鐘為單位。以台灣來說，時區是UTC +8，傳回的值是-480。
    var offsetInMinutes = date.getTimezoneOffset() + timezoneOffset * 60;
    var adjustedDate = new Date(date.getTime() + offsetInMinutes * 60 * 1000);
    // 格式化時間字串
    var formattedDateString = adjustedDate.toISOString();
    // 在字串末尾添加時區資訊
    var finalDateString = formattedDateString.replace("Z", "+08:00");
    return finalDateString
}