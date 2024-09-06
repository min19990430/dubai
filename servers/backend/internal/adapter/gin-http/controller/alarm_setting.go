package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
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
