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
      <td>physical_quantity_uuid</td>
      <td>string</td>
      <td>physical quantity UUID</td>
    </tr>
    <tr>
      <td>datetime</td>
      <td>string</td>
      <td>RFC3339<p>
      ex. "2021-09-21T02:19:00+08:00","2023-11-21T02:19:00Z"<p>
      </td>
    </tr>
    <tr>
      <td>value</td>
      <td>number</td>
      <td>Data value after calibration</td>
    </tr>
    <tr>
      <td>data</td>
      <td>number</td>
      <td>Data from device</td>
    </tr>
    <tr>
      <td>status</td>
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
  <font>"start_time"</font>: "2024-05-08T05:00:00+04:00",
  <font>"end_time"</font>: "2024-05-08T07:00:00+04:00",
  <font>"time_zone"</font>: "+04:00"
}
</pre>

#### Response
<pre>
{
  <font>"data"</font>: [
    {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"physical_quantity_uuid"</font>: "9019c4e1-e40d-4e77-b69e-87f70c4de96d",
        <font>"datetime"</font>: "2024-05-08T06:15:37+04:00",
        <font>"value"</font>: 60.71,
        <font>"data"</font>: 60.71,
        <font>"status"</font>: "10"
    },
    {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"physical_quantity_uuid"</font>: "184c9a57-65bc-4271-97d6-b983431e511e",
        <font>"datetime"</font>: "2024-05-08T06:15:37+04:00",
        <font>"value"</font>: 0,
        <font>"data"</font>: 0,
        <font>"status"</font>: "11"
    },
    {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"physical_quantity_uuid"</font>: "8df500a9-41b8-4aa0-b072-2fa2313f9677",
        <font>"datetime"</font>: "2024-05-08T06:15:37+04:00",
        <font>"value"</font>: 821.28,
        <font>"data"</font>: 821.28,
        <font>"status"</font>: "10"
    },
    {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"physical_quantity_uuid"</font>: "a03598ab-8b1a-48ec-8ecf-d1e7838127df",
        <font>"datetime"</font>: "2024-05-08T06:15:37+04:00",
        <font>"value"</font>: 85.75,
        <font>"data"</font>: 23.86,
        <font>"status"</font>: "10"
    },
    {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"physical_quantity_uuid"</font>: "675a68b9-e5f2-4b25-9ddb-d73cfb4c2d6a",
        <font>"datetime"</font>: "2024-05-08T06:15:37+04:00",
        <font>"value"</font>: 31,
        <font>"data"</font>: 31,
        <font>"status"</font>: "10"
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