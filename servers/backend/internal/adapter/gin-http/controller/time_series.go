package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
)

type TimeSeriesController struct {
	response   response.IResponse
	timeSeries usecase.TimeSeriesUsecase
}

func NewTimeSeriesController(response response.IResponse, timeSeriesUsecase *usecase.TimeSeriesUsecase) *TimeSeriesController {
	return &TimeSeriesController{
		response:   response,
		timeSeries: *timeSeriesUsecase,
	}
}

type TimeSeriesStationRequest struct {
	StationUUID string    `json:"station_uuid" binding:"required,uuid4"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	Interval    string    `json:"interval" binding:"required"`
	Reverse     bool      `json:"reverse" binding:"omitempty"`
}

func (tsc *TimeSeriesController) PostListByStation(c *gin.Context) {
	var request TimeSeriesStationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		tsc.response.ValidatorFail(c, paramError)
		return
	}

	aggregateData, err := tsc.timeSeries.AggregateDataByStation(request.StartTime, request.EndTime, request.StationUUID, request.Interval, request.Reverse)
	if err != nil {
		tsc.response.FailWithError(c, "fail to get time series", err)
		return
	}

	tsc.response.SuccessWithData(c, "success", aggregateData)
}

func (tsc *TimeSeriesController) PostListJSONByStation(c *gin.Context) {
	var request TimeSeriesStationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		tsc.response.ValidatorFail(c, paramError)
		return
	}

	aggregateData, err := tsc.timeSeries.AggregateMapDataByStation(request.StartTime, request.EndTime, request.StationUUID, request.Interval)
	if err != nil {
		tsc.response.FailWithError(c, "fail to get time series", err)
		return
	}

	tsc.response.SuccessWithData(c, "success", aggregateData)
}
