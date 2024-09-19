package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
)

type ControlSignalController struct {
	response      response.IResponse
	controlSignal usecase.ControlSignalUsecase
}

func NewControlSignalController(
	response response.IResponse,
	controlSignal *usecase.ControlSignalUsecase,
) *ControlSignalController {
	return &ControlSignalController{
		response:      response,
		controlSignal: *controlSignal,
	}
}

type ControlSignalListRequest struct {
	UUID []string `json:"uuid" binding:"required"`
}

func (cc *ControlSignalController) PostList(c *gin.Context) {
	var request ControlSignalListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		cc.response.ValidatorFail(c, paramError)
		return
	}

	result, err := cc.controlSignal.ListIn(request.UUID)
	if err != nil {
		cc.response.FailWithError(c, queryFail, err)
		return
	}

	cc.response.SuccessWithData(c, querySuccess, result)
}

type ControlSignalUpdateSignalValueRequest struct {
	UUID        string `json:"uuid" binding:"required"`
	SignalValue string `json:"signal_value" binding:"required"`
}

func (cc *ControlSignalController) UpdateSignalValue(c *gin.Context) {
	var request ControlSignalUpdateSignalValueRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		cc.response.ValidatorFail(c, paramError)
		return
	}

	err := cc.controlSignal.UpdateSignalValue(request.UUID, request.SignalValue)
	if err != nil {
		cc.response.FailWithError(c, updateFail, err)
		return
	}

	cc.response.Success(c, updateSuccess)
}
