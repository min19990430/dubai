package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type AlarmRecordController struct {
	response    response.IResponse
	alarmRecord usecase.AlarmRecordUsecase
}

func NewAlarmRecordController(response response.IResponse, usecase *usecase.AlarmRecordUsecase) *AlarmRecordController {
	return &AlarmRecordController{
		response:    response,
		alarmRecord: *usecase,
	}
}

type AlarmRecordRequest struct {
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

func (arc *AlarmRecordController) PostList(c *gin.Context) {
	var request AlarmRecordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		arc.response.ValidatorFail(c, paramError)
		return
	}

	result, err := arc.alarmRecord.List(request.StartTime, request.EndTime, domain.AlarmRecord{}, true)
	if err != nil {
		arc.response.FailWithError(c, queryFail, err)
		return
	}

	arc.response.SuccessWithData(c, querySuccess, result)
}

type AlarmRecordDeviceRequest struct {
	DeviceUUID string    `json:"device_uuid" binding:"required,uuid4"`
	StartTime  time.Time `json:"start_time" binding:"required"`
	EndTime    time.Time `json:"end_time" binding:"required"`
}

func (arc *AlarmRecordController) PostListByDeviceUUID(c *gin.Context) {
	var request AlarmRecordDeviceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		arc.response.ValidatorFail(c, paramError)
		return
	}

	result, err := arc.alarmRecord.ListByDeviceUUID(request.StartTime, request.EndTime, request.DeviceUUID, true)
	if err != nil {
		arc.response.FailWithError(c, queryFail, err)
		return
	}

	arc.response.SuccessWithData(c, querySuccess, result)
}

type AlarmRecordStationRequest struct {
	StationUUID string    `json:"station_uuid" binding:"required,uuid4"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
}

func (arc *AlarmRecordController) PostListByStationUUID(c *gin.Context) {
	var request AlarmRecordStationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		arc.response.ValidatorFail(c, paramError)
		return
	}

	result, err := arc.alarmRecord.ListByStationUUID(request.StartTime, request.EndTime, request.StationUUID, true)
	if err != nil {
		arc.response.FailWithError(c, queryFail, err)
		return
	}

	arc.response.SuccessWithData(c, querySuccess, result)
}

type AlarmRecordDetailRequest struct {
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

func (arc *AlarmRecordController) PostListDetail(c *gin.Context) {
	var request AlarmRecordDetailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		arc.response.ValidatorFail(c, paramError)
		return
	}

	result, err := arc.alarmRecord.ListDetail(request.StartTime, request.EndTime, domain.AlarmRecord{}, true)
	if err != nil {
		arc.response.FailWithError(c, queryFail, err)
		return
	}

	arc.response.SuccessWithData(c, querySuccess, result)
}
