package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
)

type AlarmRecordCollectionController struct {
	response response.IResponse

	alarmRecordCollection usecase.AlarmRecordCollectionUsecase
}

func NewAlarmRecordCollectionController(response response.IResponse, usecase *usecase.AlarmRecordCollectionUsecase) *AlarmRecordCollectionController {
	return &AlarmRecordCollectionController{
		response:              response,
		alarmRecordCollection: *usecase,
	}
}

type AlarmRecordCollectionDeviceRequest struct {
	DeviceUUID string    `json:"device_uuid" binding:"required,uuid4"`
	StartTime  time.Time `json:"start_time" binding:"required"`
	EndTime    time.Time `json:"end_time" binding:"required"`
}

func (arc *AlarmRecordCollectionController) PostListByDeviceUUID(c *gin.Context) {
	var request AlarmRecordCollectionDeviceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		arc.response.ValidatorFail(c, paramError)
		return
	}

	if request.StartTime.After(request.EndTime) {
		arc.response.ValidatorFail(c, "The start time should not be greater than the end time")
		return
	}

	result, err := arc.alarmRecordCollection.ListByDeviceUUID(request.StartTime, request.EndTime, request.DeviceUUID, true)
	if err != nil {
		arc.response.FailWithError(c, queryFail, err)
		return
	}

	arc.response.SuccessWithData(c, querySuccess, result)
}

type AlarmRecordCollectionStationRequest struct {
	StationUUID string    `json:"station_uuid" binding:"required,uuid4"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
}

func (arc *AlarmRecordCollectionController) PostListByStationUUID(c *gin.Context) {
	var request AlarmRecordCollectionStationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		arc.response.ValidatorFail(c, paramError)
		return
	}

	if request.StartTime.After(request.EndTime) {
		arc.response.ValidatorFail(c, "The start time should not be greater than the end time")
		return
	}

	result, err := arc.alarmRecordCollection.ListByStationUUID(request.StartTime, request.EndTime, request.StationUUID, true)
	if err != nil {
		arc.response.FailWithError(c, queryFail, err)
		return
	}

	arc.response.SuccessWithData(c, querySuccess, result)
}
