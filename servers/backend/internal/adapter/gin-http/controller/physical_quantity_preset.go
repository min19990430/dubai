package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type PhysicalQuantityPresetController struct {
	response response.IResponse

	physicalQuantityPresetUsecase usecase.PhysicalQuantityPresetUsecase
}

func NewPhysicalQuantityPresetController(response response.IResponse, usecase *usecase.PhysicalQuantityPresetUsecase) *PhysicalQuantityPresetController {
	return &PhysicalQuantityPresetController{
		response:                      response,
		physicalQuantityPresetUsecase: *usecase,
	}
}

func (pqpc *PhysicalQuantityPresetController) GetOne(c *gin.Context) {
	uuid := c.Param("uuid")

	physicalQuantityPreset, err := pqpc.physicalQuantityPresetUsecase.Get(domain.PhysicalQuantityPreset{UUID: uuid})
	if err != nil {
		pqpc.response.FailWithError(c, queryFail, err)
		return
	}

	pqpc.response.SuccessWithData(c, querySuccess, physicalQuantityPreset)
}

func (pqpc *PhysicalQuantityPresetController) GetList(c *gin.Context) {
	physicalQuantityPresets, err := pqpc.physicalQuantityPresetUsecase.List(domain.PhysicalQuantityPreset{})
	if err != nil {
		pqpc.response.FailWithError(c, queryFail, err)
		return
	}

	pqpc.response.SuccessWithData(c, querySuccess, physicalQuantityPresets)
}

type PhysicalQuantityPresetRequest struct {
	UUID                       string `json:"uuid" binding:"required,uuid4"`
	Priority                   int    `json:"priority"`
	Name                       string `json:"name"`
	FullName                   string `json:"full_name"`
	SiUnit                     string `json:"si_unit"`
	StatusCode                 string `json:"status_code"`
	PhysicalQuantityDataType   string `json:"physical_quantity_data_type"`
	AggregateCalculationMethod string `json:"aggregate_calculation_method"`
	CalibrationEnable          bool   `json:"calibration_enable"`
	CalibrationValue           string `json:"calibration_value"`
	CalibrationParameter       string `json:"calibration_parameter"`
	Description                string `json:"description"`
}

func (pqpc *PhysicalQuantityPresetController) PostCreate(c *gin.Context) {
	var request PhysicalQuantityPresetRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		pqpc.response.ValidatorFail(c, paramError)
		return
	}

	physicalQuantityPreset := domain.PhysicalQuantityPreset{
		UUID:                       request.UUID,
		Priority:                   request.Priority,
		Name:                       request.Name,
		FullName:                   request.FullName,
		SiUnit:                     request.SiUnit,
		StatusCode:                 request.StatusCode,
		PhysicalQuantityDataType:   request.PhysicalQuantityDataType,
		AggregateCalculationMethod: request.AggregateCalculationMethod,
		CalibrationEnable:          request.CalibrationEnable,
		CalibrationValue:           request.CalibrationValue,
		CalibrationParameter:       request.CalibrationParameter,
		Description:                request.Description,
	}

	err = pqpc.physicalQuantityPresetUsecase.Create(physicalQuantityPreset)
	if err != nil {
		pqpc.response.FailWithError(c, createFail, err)
		return
	}

	pqpc.response.Success(c, createSuccess)
}

func (pqpc *PhysicalQuantityPresetController) PutUpdate(c *gin.Context) {
	var request PhysicalQuantityPresetRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		pqpc.response.ValidatorFail(c, paramError)
		return
	}

	physicalQuantityPreset := domain.PhysicalQuantityPreset{
		UUID:                       request.UUID,
		Priority:                   request.Priority,
		Name:                       request.Name,
		FullName:                   request.FullName,
		SiUnit:                     request.SiUnit,
		StatusCode:                 request.StatusCode,
		PhysicalQuantityDataType:   request.PhysicalQuantityDataType,
		AggregateCalculationMethod: request.AggregateCalculationMethod,
		CalibrationEnable:          request.CalibrationEnable,
		CalibrationValue:           request.CalibrationValue,
		CalibrationParameter:       request.CalibrationParameter,
		Description:                request.Description,
	}

	err = pqpc.physicalQuantityPresetUsecase.Update(physicalQuantityPreset)
	if err != nil {
		pqpc.response.FailWithError(c, updateFail, err)
		return
	}

	pqpc.response.Success(c, updateSuccess)
}

func (pqpc *PhysicalQuantityPresetController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	err := pqpc.physicalQuantityPresetUsecase.Delete(uuid)
	if err != nil {
		pqpc.response.FailWithError(c, deleteFail, err)
		return
	}

	pqpc.response.Success(c, deleteSuccess)
}
