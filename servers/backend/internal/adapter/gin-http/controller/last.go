package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type LastController struct {
	response    response.IResponse
	last        usecase.LastUsecase
	alarmRecord usecase.AlarmRecordUsecase
}

func NewLastController(
	response response.IResponse,
	usecase *usecase.LastUsecase,
	alarmRecord *usecase.AlarmRecordUsecase,
) *LastController {
	return &LastController{
		response:    response,
		last:        *usecase,
		alarmRecord: *alarmRecord,
	}
}

func (lc *LastController) GetStationLast(c *gin.Context) {
	lasts, err := lc.last.GetStationLast()
	if err != nil {
		lc.response.FailWithError(c, queryFail, err)
		return
	}

	alarmRecords, err := lc.alarmRecord.ListDetail(
		carbon.Now().StartOfDay().StdTime(),
		carbon.Now().EndOfDay().StdTime(),
		domain.AlarmRecord{}, true)
	if err != nil {
		lc.response.FailWithError(c, queryFail, err)
		return
	}

	lc.response.SuccessWithData(c, querySuccess,
		gin.H{
			"last":  lasts,
			"alarm": alarmRecords,
		})
}

func (lc *LastController) GetDeviceLast(c *gin.Context) {
	lasts, err := lc.last.GetDeviceLast()
	if err != nil {
		lc.response.FailWithError(c, queryFail, err)
		return
	}

	alarmRecords, err := lc.alarmRecord.ListDetail(
		carbon.Now().StartOfDay().StdTime(),
		carbon.Now().EndOfDay().StdTime(),
		domain.AlarmRecord{}, true)
	if err != nil {
		lc.response.FailWithError(c, queryFail, err)
		return
	}

	lc.response.SuccessWithData(c, querySuccess,
		gin.H{
			"last":  lasts,
			"alarm": alarmRecords,
		})
}
