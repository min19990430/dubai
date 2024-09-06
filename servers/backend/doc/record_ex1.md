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

# Record API

## Request
### HTTP request
<pre>POST https://iot.fenri.com.tw/api/v1/XXXXXXXXXX</pre>


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
      <td>device_uuid</td>
      <td>string</td>
      <td>device UUID</td>
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
      <td>number</td>
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
      <td>number</td>
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
      <td>number</td>
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
      <td>Turbidity</td>
      <td>number</td>
      <td>Turbidity data value after calibration</td>
    </tr>
    <tr>
      <td>Turbidity_status</td>
      <td>string</td>
      <td>"10":"normal"<p>
      "11":"alarm"<p>
      "20":"maintain"<p>
      </td>
    </tr>
    <tr>
      <td>TEMP.</td>
      <td>number</td>
      <td>TEMP. data value after calibration</td>
    </tr>
    <tr>
      <td>TEMP._status</td>
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
  <font>"end_time"</font>: "2024-05-09T23:00:00+04:00",
  <font>"time_zone"</font>: "+04:00"
}
</pre>

#### Response
<pre>
{
  <font>"data"</font>: [
    {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"datetime"</font>: "2024-05-09T00:00:00+04:00",
        <font>"pH"</font>: 7.7,
        <font>"pH_status"</font>: "20",
        <font>"EC"</font>: 705.1,
        <font>"EC_status"</font>: "10",
        <font>"COD"</font>: 10.9,
        <font>"COD_status"</font>: "10",
        <font>"Turbidity"</font>: 47.5,
        <font>"Turbidity_status"</font>: "10",
        <font>"TEMP."</font>: 32.1,
        <font>"TEMP._status"</font>: "10",
    },
    {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"datetime"</font>: "2024-05-09T00:05:00+04:00",
        <font>"pH"</font>: 7.9,
        <font>"pH_status"</font>: "20",
        <font>"EC"</font>: 707.2,
        <font>"EC_status"</font>: "10",
        <font>"COD"</font>: 10.9,
        <font>"COD_status"</font>: "10",
        <font>"Turbidity"</font>: 48.5,
        <font>"Turbidity_status"</font>: "10",
        <font>"TEMP."</font>: 32.1,
        <font>"TEMP._status"</font>: "10",
    },
    {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"datetime"</font>: "2024-05-09T00:10:00+04:00",
        <font>"pH"</font>: 7.8,
        <font>"pH_status"</font>: "20",
        <font>"EC"</font>: 704.1,
        <font>"EC_status"</font>: "10",
        <font>"COD"</font>: 10.8,
        <font>"COD_status"</font>: "10",
        <font>"Turbidity"</font>: 46.7,
        <font>"Turbidity_status"</font>: "10",
        <font>"TEMP."</font>: 32.1,
        <font>"TEMP._status"</font>: "10",
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