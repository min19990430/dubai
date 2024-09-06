package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type PhysicalQuantityController struct {
	response response.IResponse

	physicalQuantityUsecase usecase.PhysicalQuantityUsecase
}

func NewPhysicalQuantityController(response response.IResponse, usecase *usecase.PhysicalQuantityUsecase) *PhysicalQuantityController {
	return &PhysicalQuantityController{
		response:                response,
		physicalQuantityUsecase: *usecase,
	}
}

type UpdateStatusRequest struct {
	UUID       string `json:"uuid" binding:"required,uuid4"`
	StatusCode string `json:"status_code" binding:"required,oneof='10' '11' '20' '30' '31' '32' '40' '93'"`
}

func (pqc *PhysicalQuantityController) UpdateStatus(c *gin.Context) {
	var request UpdateStatusRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		pqc.response.ValidatorFail(c, paramError)
		return
	}

	err = pqc.physicalQuantityUsecase.UpdateStatus(domain.PhysicalQuantity{UUID: request.UUID}, request.StatusCode)
	if err != nil {
		pqc.response.FailWithError(c, "failed to update status", err)
		return
	}

	pqc.response.Success(c, "success")
}
