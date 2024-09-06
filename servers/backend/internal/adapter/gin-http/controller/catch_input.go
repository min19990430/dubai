package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/convert"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type CatchInputController struct {
	response   response.IResponse
	catchInput usecase.CatchInputUsecase
	convert    convert.SignalInputToInput
}

func NewCatchInputController(response response.IResponse, usecase *usecase.CatchInputUsecase, convert *convert.SignalInputToInput) *CatchInputController {
	return &CatchInputController{
		response:   response,
		catchInput: *usecase,
		convert:    *convert,
	}
}

type CatchSignalInputRequest struct {
	DeviceUUID string    `json:"device_uuid" binding:"required,uuid"`
	Datetime   time.Time `json:"datetime"`
	DI         string    `json:"di"`
	AI         []float64 `json:"ai"`
}

func (cic *CatchInputController) CatchSignalInput(c *gin.Context) {
	var request []CatchSignalInputRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		cic.response.ValidatorFail(c, paramError)
		return
	}

	var signalInputs []domain.SignalInput
	for _, r := range request {
		signalInputs = append(signalInputs, domain.SignalInput{
			DeviceUUID: r.DeviceUUID,
			Datetime:   r.Datetime,
			DI:         r.DI,
			AI:         r.AI,
		})
	}

	inputsWithDeviceUUID, convertErr := cic.convert.SignalInputToInput(signalInputs)
	if convertErr != nil {
		cic.response.FailWithError(c, catchFail, convertErr)
		return
	}

	catchInputErr := cic.catchInput.CatchInput(inputsWithDeviceUUID)
	if catchInputErr != nil {
		cic.response.FailWithError(c, catchFail, catchInputErr)
		return
	}

	cic.response.Success(c, "success")
}
