package controller

import (
	"log"

	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type AlarmSettingController struct {
	response     response.IResponse
	alarmSetting usecase.AlarmSettingUsecase
}

func NewAlarmSettingController(response response.IResponse, usecase *usecase.AlarmSettingUsecase) *AlarmSettingController {
	return &AlarmSettingController{
		response:     response,
		alarmSetting: *usecase,
	}
}

type AlarmSettingTestExpressionRequest struct {
	Expression string  `json:"expression" binding:"required"`
	Value      float64 `json:"value"`
}

func (asc *AlarmSettingController) TestExpression(c *gin.Context) {
	var request AlarmSettingTestExpressionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		asc.response.ValidatorFail(c, paramError)
		return
	}

	result, err := asc.alarmSetting.TestExpression(request.Expression, request.Value)
	if err != nil {
		asc.response.FailWithError(c, testFail, err)
		return
	}

	asc.response.SuccessWithData(c, testSuccess, result)
}

type AlarmSettingUpdateExpressionRequest struct {
	UUID       string `json:"uuid" binding:"required,uuid4"`
	Expression string `json:"expression" binding:"required"`
}

func (asc *AlarmSettingController) UpdateExpression(c *gin.Context) {
	var request AlarmSettingUpdateExpressionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		asc.response.ValidatorFail(c, paramError)
		return
	}

	err := asc.alarmSetting.UpdateExpression(request.UUID, request.Expression)
	if err != nil {
		asc.response.FailWithError(c, updateFail, err)
		return
	}

	asc.response.Success(c, updateSuccess)
}

func (asc *AlarmSettingController) ListByDeviceUUID(c *gin.Context) {
	deviceUUID, ok := c.GetQuery("DeviceUUID")
	if !ok {
		asc.response.ValidatorFail(c, paramError)
		return
	}

	alarmSettings, err := asc.alarmSetting.ListByDeviceUUID(deviceUUID)
	if err != nil {
		asc.response.FailWithError(c, queryFail, err)
		return
	}

	asc.response.SuccessWithData(c, querySuccess, alarmSettings)
}

func (asc *AlarmSettingController) ListByStationUUID(c *gin.Context) {
	stationUUID, ok := c.GetQuery("StationUUID")
	if !ok {
		asc.response.ValidatorFail(c, paramError)
		return
	}

	alarmSettings, err := asc.alarmSetting.ListByStationUUID(stationUUID)
	if err != nil {
		asc.response.FailWithError(c, queryFail, err)
		return
	}

	asc.response.SuccessWithData(c, querySuccess, alarmSettings)
}

type AlarmSettingRequest struct {
	UUID                 string `json:"uuid" binding:"required,uuid4"`
	PhysicalQuantityUUID string `json:"physical_quantity_uuid" binding:"required,uuid4"`
	Name                 string `json:"name" binding:"required"`
	FullName             string `json:"full_name" binding:"required"`
	Priority             int    `json:"priority" `
	IsEnable             bool   `json:"is_enable"`
	IsNotify             bool   `json:"is_notify"`
	IsActivated          bool   `json:"is_activated"`
	BooleanExpression    string `json:"boolean_expression"`
	AlarmContentSetting  string `json:"alarm_content_setting"`
}

func (asc *AlarmSettingController) PostCreate(c *gin.Context) {
	var request AlarmSettingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		asc.response.ValidatorFail(c, paramError)
		return
	}

	err := asc.alarmSetting.Create(domain.AlarmSetting{
		UUID:                 request.UUID,
		PhysicalQuantityUUID: request.PhysicalQuantityUUID,
		Name:                 request.Name,
		FullName:             request.FullName,
		Priority:             request.Priority,
		IsEnable:             request.IsEnable,
		IsNotify:             request.IsNotify,
		IsActivated:          request.IsActivated,
		BooleanExpression:    request.BooleanExpression,
		AlarmContentSetting:  &request.AlarmContentSetting,
	})
	if err != nil {
		asc.response.FailWithError(c, createFail, err)
		return
	}

	asc.response.Success(c, createSuccess)
}

func (asc *AlarmSettingController) PutUpdate(c *gin.Context) {
	var request AlarmSettingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		asc.response.ValidatorFail(c, paramError)
		return
	}

	err := asc.alarmSetting.Update(domain.AlarmSetting{
		UUID:                 request.UUID,
		PhysicalQuantityUUID: request.PhysicalQuantityUUID,
		Name:                 request.Name,
		FullName:             request.FullName,
		Priority:             request.Priority,
		IsEnable:             request.IsEnable,
		IsNotify:             request.IsNotify,
		IsActivated:          request.IsActivated,
		BooleanExpression:    request.BooleanExpression,
		AlarmContentSetting:  &request.AlarmContentSetting,
	})
	if err != nil {
		asc.response.FailWithError(c, updateFail, err)
		return
	}

	asc.response.Success(c, updateSuccess)
}

func (asc *AlarmSettingController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	err := asc.alarmSetting.Delete(uuid)
	if err != nil {
		asc.response.FailWithError(c, deleteFail, err)
		return
	}

	asc.response.Success(c, deleteSuccess)
}
