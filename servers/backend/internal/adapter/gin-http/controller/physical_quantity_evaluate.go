package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type PhysicalQuantityEvaluateController struct {
	response response.IResponse

	physicalQuantityEvaluateUsecase       usecase.PhysicalQuantityEvaluateUsecase
	PhysicalQuantityEvaluateDetailUsecase usecase.PhysicalQuantityEvaluateDetailUsecase
}

func NewPhysicalQuantityEvaluateController(response response.IResponse, usecase *usecase.PhysicalQuantityEvaluateUsecase, usecaseDetail *usecase.PhysicalQuantityEvaluateDetailUsecase) *PhysicalQuantityEvaluateController {
	return &PhysicalQuantityEvaluateController{
		response:                              response,
		physicalQuantityEvaluateUsecase:       *usecase,
		PhysicalQuantityEvaluateDetailUsecase: *usecaseDetail,
	}
}

func (pqec *PhysicalQuantityEvaluateController) GetList(c *gin.Context) {
	physicalQuantityEvaluateList, err := pqec.PhysicalQuantityEvaluateDetailUsecase.ListDetail(domain.PhysicalQuantityEvaluate{})
	if err != nil {
		pqec.response.FailWithError(c, queryFail, err)
		return
	}

	pqec.response.SuccessWithData(c, querySuccess, physicalQuantityEvaluateList)
}

type PhysicalQuantityEvaluateRequest struct {
	UUID        string         `json:"uuid" binding:"required,uuid4"`
	FormulaType string         `json:"formula_type" binding:"required,oneof='Normal' 'SemiCircularWaterFlow' 'RectangularWaterFlow'"`
	Params      map[string]any `json:"params"`
}

func (pqec *PhysicalQuantityEvaluateController) UpdateFormula(c *gin.Context) {
	var req PhysicalQuantityEvaluateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pqec.response.ValidatorFail(c, paramError)
	}

	err := pqec.physicalQuantityEvaluateUsecase.UpdateFormula(domain.PhysicalQuantityEvaluate{UUID: req.UUID}, req.FormulaType, req.Params)
	if err != nil {
		pqec.response.FailWithError(c, updateFail, err)
		return
	}

	pqec.response.Success(c, updateSuccess)
}
