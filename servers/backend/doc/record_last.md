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

# Record/Last API

## Request
### HTTP request
<pre>POST https://iot.fenri.com.tw/api/v1/Record/Last</pre>


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
      <td>physical_quantity_uuid</td>
      <td>string</td>
      <td>physical quantity UUID</td>
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
      <td>object(JSON)</td>
      <td>record</td>
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
  <font>"physical_quantity_uuid"</font>: "47d28dc4-b888-4880-ab23-3d1c46b24f06",
  <font>"time_zone"</font>: "+04:00"
}
</pre>

#### Response
<pre>
{
  <font>"data"</font>: {
        <font>"device_uuid"</font>: "f52132d2-8a8f-4acc-afbd-6ec556b12182",
        <font>"physical_quantity_uuid"</font>: "47d28dc4-b888-4880-ab23-3d1c46b24f06",
        <font>"datetime"</font>: "2024-05-08T07:46:43+04:00",
        <font>"value"</font>: 7.88,
        <font>"data"</font>: 7.88,
        <font>"status"</font>: "20",
  },
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