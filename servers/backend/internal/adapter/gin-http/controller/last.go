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
	alarmRecord *usecase.AlarmRecordUsecase) *LastController {
	return &LastController{
		response:    response,
		last:        *usecase,
		alarmRecord: *alarmRecord,
	}
}

func (lc *LastController) GetLast(c *gin.Context) {
	lasts, err := lc.last.GetLast()
	if err != nil {
		lc.response.FailWithError(c, "fail to get last", err)
		return
	}

	alarmRecords, err := lc.alarmRecord.ListDetail(
		carbon.Now().StartOfDay().ToStdTime(),
		carbon.Now().EndOfDay().ToStdTime(),
		domain.AlarmRecord{}, true)
	if err != nil {
		lc.response.FailWithError(c, "fail to get alarm records", err)
		return
	}

	lc.response.SuccessWithData(c, "success",
		gin.H{
			"last":  lasts,
			"alarm": alarmRecords,
		})
}
