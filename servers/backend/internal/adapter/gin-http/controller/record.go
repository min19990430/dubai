package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
)

type RecordController struct {
	response response.IResponse

	record usecase.RecordUsecase
}

func NewRecordController(response response.IResponse, record *usecase.RecordUsecase) *RecordController {
	return &RecordController{
		response: response,
		record:   *record,
	}
}

type RecordLastRequest struct {
	PhysicalQuantityUUID string `json:"physical_quantity_uuid" binding:"required,uuid4"`
	TimeZone             string `json:"time_zone" binding:"required,timezone"`
}

func (rc *RecordController) LastRecord(c *gin.Context) {
	var request RecordLastRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		rc.response.ValidatorFail(c, paramError)
		return
	}

	record, err := rc.record.Last(request.PhysicalQuantityUUID, request.TimeZone)
	if err != nil {
		rc.response.FailWithError(c, queryFail, err)
		return
	}

	rc.response.SuccessWithData(c, "success", record)
}

type RecordListRequest struct {
	DeviceUUID string    `json:"device_uuid" binding:"required,uuid4"`
	StartTime  time.Time `json:"start_time" binding:"required"`
	EndTime    time.Time `json:"end_time" binding:"required"`
	TimeZone   string    `json:"time_zone" binding:"required,timezone"`
}

func (rc *RecordController) PostList(c *gin.Context) {
	var request RecordListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		rc.response.ValidatorFail(c, paramError)
		return
	}

	if request.StartTime.After(request.EndTime) {
		rc.response.ValidatorFail(c, "The start time should not be greater than the end time")
		return
	}

	if diff := request.EndTime.Sub(request.StartTime); diff > 90*24*time.Hour {
		rc.response.ValidatorFail(c, "The time interval should not exceed 90 days")
		return
	}

	records, err := rc.record.List(request.StartTime, request.EndTime, request.DeviceUUID, request.TimeZone)
	if err != nil {
		rc.response.FailWithError(c, queryFail, err)
		return
	}

	rc.response.SuccessWithData(c, "success", records)
}

type RecordListDeviceRequest struct {
	DeviceUUID string    `json:"device_uuid" binding:"required,uuid4"`
	StartTime  time.Time `json:"start_time" binding:"required"`
	EndTime    time.Time `json:"end_time" binding:"required"`
	TimeZone   string    `json:"time_zone" binding:"required,timezone"`
}

func (rc *RecordController) PostListDeviceJSON(c *gin.Context) {
	var request RecordListDeviceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		rc.response.ValidatorFail(c, paramError)
		return
	}

	if request.StartTime.After(request.EndTime) {
		rc.response.ValidatorFail(c, "The start time should not be greater than the end time")
		return
	}

	if diff := request.EndTime.Sub(request.StartTime); diff > 90*24*time.Hour {
		rc.response.ValidatorFail(c, "The time interval should not exceed 90 days")
		return
	}

	records, err := rc.record.ListMapByDevice(request.StartTime, request.EndTime, request.DeviceUUID, request.TimeZone)
	if err != nil {
		rc.response.FailWithError(c, queryFail, err)
		return
	}

	rc.response.SuccessWithData(c, "success", records)
}

type RecordListStationRequest struct {
	StationUUID string    `json:"station_uuid" binding:"required,uuid4"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	TimeZone    string    `json:"time_zone" binding:"required,timezone"`
}

func (rc *RecordController) PostListStationJSON(c *gin.Context) {
	var request RecordListStationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		rc.response.ValidatorFail(c, paramError)
		return
	}

	if request.StartTime.After(request.EndTime) {
		rc.response.ValidatorFail(c, "The start time should not be greater than the end time")
		return
	}

	if diff := request.EndTime.Sub(request.StartTime); diff > 90*24*time.Hour {
		rc.response.ValidatorFail(c, "The time interval should not exceed 90 days")
		return
	}

	records, err := rc.record.ListMapByStation(request.StartTime, request.EndTime, request.StationUUID, request.TimeZone)
	if err != nil {
		rc.response.FailWithError(c, queryFail, err)
		return
	}

	rc.response.SuccessWithData(c, "success", records)
}
