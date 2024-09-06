<style>
table {
  font-family: 'Courier New', 'Microsoft JhengHei', Arial, sans-serif;
  width: 100%;
}
table > thead {
  color: black;
  background-color: #D3D3D3;
}
pre {
  font-family: 'Courier New', 'Microsoft JhengHei', Arial, sans-serif;
  background-color: WhiteSmoke;
}
pre > font {
  color: #36BF36;
  font-weight: 700;
}
pre > var {
  color: #FF4D40;
  font-weight: 700;
}
.query {
  color: black;
  background-color: #ECECEC;
}
</style>

# Record/Sensor/JSON API

## Request
### HTTP request
<pre>POST https://iot.fenri.com.tw/api/v1/Record/Sensor/JSON</pre>


### Request body
<table>
    <thead>
        <tr>
            <th>Parameter</th>
            <th>Value</th>
            <th>Description</th>
        </tr>
    </thead>
      <tbody>
    <tr>
      <td>device_uuid</td>
      <td>string</td>
      <td>device UUID</td>
    </tr>
    <tr>
      <td>start_time</td>
      <td>string</td>
      <td>RFC3339<p>
      ex. "2021-09-21T02:19:00+04:00","2023-11-21T02:19:00Z"<p>
      **WARNING** should not exceed 90 days
      </td>
    </tr>
    <tr>
      <td>end_time</td>
      <td>string</td>
      <td>RFC3339<p>
      ex. "2021-09-21T02:19:00+04:00","2023-11-21T02:19:00Z"<p>
      **WARNING** should not exceed 90 days
      </td>
    </tr>
    <tr>
      <td>time_zone</td>
      <td>string</td>
      <td>ex. "+00:00","+04:00","+08:00","-07:00"</td>
    </tr>
  </tobdy>
<table>

---
## Response
<table>
  <thead>
    <tr>
      <th>Parameter</th>
      <th>Value</th>
      <th>Description</th>
    </tr> 
  </thead>
  <tbody>
    <tr>
      <td>success</td>
      <td>bool</td>
      <td>Indicates whether the request was successful</td>
    </tr>
    <tr>
      <td>message</td>
      <td>string</td>
      <td>Message or error information</td>
    </tr>
    <tr>
      <td>data</td>
      <td>objects(JSON)</td>
      <td>records</td>
    </tr>
  </tobdy>
</table>

<table>
  <thead>
    <tr>
      <th>record</th>
    </tr>
    <tr>
      <th>Parameter</th>
      <th>Value</th>
      <th>Description</th>
    </tr> 
  </thead>
  <tbody>
    <tr>
      <td>device_id</td>
      <td>string</td>
      <td>device ID</td>
    </tr>
    <tr>
      <td>datetime</td>
      <td>string</td>
      <td>RFC3339<p>
      ex. "2021-09-21T02:19:00+08:00","2023-11-21T02:19:00Z"<p>
      </td>
    </tr>
    <tr>
      <td>pH</td>
      <td>string</td>
      <td>pH data value after calibration</td>
    </tr>
    <tr>
      <td>pH_status</td>
      <td>string</td>
      <td>"10":"normal"<p>
      "11":"alarm"<p>
      "20":"maintain"<p>
      </td>
    </tr>
    <tr>
      <td>EC</td>
      <td>string</td>
      <td>EC data value after calibration</td>
    </tr>
    <tr>
      <td>EC_status</td>
      <td>string</td>
      <td>"10":"normal"<p>
      "11":"alarm"<p>
      "20":"maintain"<p>
      </td>
    </tr>
    <tr>
      <td>COD</td>
      <td>string</td>
      <td>COD data value after calibration</td>
    </tr>
    <tr>
      <td>COD_status</td>
      <td>string</td>
      <td>"10":"normal"<p>
      "11":"alarm"<p>
      "20":"maintain"<p>
      </td>
    </tr>
    <tr>
      <td>turbidity</td>
      <td>string</td>
      <td>Turbidity data value after calibration</td>
    </tr>
    <tr>
      <td>turbidity_status</td>
      <td>string</td>
      <td>"10":"normal"<p>
      "11":"alarm"<p>
      "20":"maintain"<p>
      </td>
    </tr>
    <tr>
      <td>temp</td>
      <td>string</td>
      <td>temp data value after calibration</td>
    </tr>
    <tr>
      <td>temp_status</td>
      <td>string</td>
      <td>"10":"normal"<p>
      "11":"alarm"<p>
      "20":"maintain"<p>
      </td>
    </tr>
  </tobdy>
</table>

## example
### Status 200
#### Request body
<pre>
{
  <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
  <font>"start_time"</font>: "2024-05-09T00:00:00+04:00",
  <font>"end_time"</font>: "2024-05-10T15:00:00+04:00",
  <font>"time_zone"</font>: "+04:00"
}
</pre>

#### Response
<pre>
{
  <font>"data"</font>: [
    {
        <font>"COD"</font>: "9.560000",
        <font>"COD_status"</font>: "10",
        <font>"EC"</font>: "842.540000",
        <font>"EC_status"</font>: "10",
        <font>"datetime"</font>: "2024-05-09T06:12:35+04:00",
        <font>"device_id"</font>: "FCM_113002",
        <font>"pH"</font>: "7.730000",
        <font>"pH_status"</font>: "20",
        <font>"temp"</font>: "31.010000",
        <font>"temp_status"</font>: "10",
        <font>"turbidity"</font>: "55.640000",
        <font>"turbidity_status"</font>: "10",
    },
    {
        <font>"COD"</font>: "9.420000",
        <font>"COD_status"</font>: "10",
        <font>"EC"</font>: "843.290000",
        <font>"EC_status"</font>: "10",
        <font>"datetime"</font>: "2024-05-09T06:17:22+04:00",
        <font>"device_id"</font>: "FCM_113002",
        <font>"pH"</font>: "7.700000",
        <font>"pH_status"</font>: "20",
        <font>"temp"</font>: "31.040000",
        <font>"temp_status"</font>: "10",
        <font>"turbidity"</font>: "59.960000",
        <font>"turbidity_status"</font>: "10",
    },
    {
        <font>"COD"</font>: "9.220000",
        <font>"COD_status"</font>: "10",
        <font>"EC"</font>: "841.440000",
        <font>"EC_status"</font>: "10",
        <font>"datetime"</font>: "2024-05-09T06:22:09+04:00",
        <font>"device_id"</font>: "FCM_113002",
        <font>"pH"</font>: "7.710000",
        <font>"pH_status"</font>: "20",
        <font>"temp"</font>: "31.050000",
        <font>"temp_status"</font>: "10",
        <font>"turbidity"</font>: "57.340000",
        <font>"turbidity_status"</font>: "10",
    },...
  ],
  <font>"message"</font>: "success",
  <font>"success"</font>: true
}
</pre>

----

### Status 400
<pre>
{
  <font>"data"</font>: null,
  <font>"message"</font>: "param error",
  <font>"success"</font>: false
}
</pre>

----

### Status 503
<pre>
{
  <font>"data"</font>: null,
  <font>"message"</font>: "query fail",
  <font>"success"</font>: false
}
</pre>