package controller

import (
	"log"

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

func (pqc *PhysicalQuantityController) GetOne(c *gin.Context) {
	uuid := c.Param("uuid")

	physicalQuantity, err := pqc.physicalQuantityUsecase.FindByUUID(uuid)
	if err != nil {
		pqc.response.FailWithError(c, queryFail, err)
		return
	}

	pqc.response.SuccessWithData(c, querySuccess, physicalQuantity)
}

func (pqc *PhysicalQuantityController) GetList(c *gin.Context) {
	stationUUID := c.Query("station_uuid")
	deviceUUID := c.Query("device_uuid")

	physicalQuantities, err := pqc.physicalQuantityUsecase.List(domain.PhysicalQuantity{
		StationUUID: stationUUID,
		DeviceUUID:  deviceUUID,
	})
	if err != nil {
		pqc.response.FailWithError(c, queryFail, err)
		return
	}

	pqc.response.SuccessWithData(c, querySuccess, physicalQuantities)
}

type PhysicalQuantityRequest struct {
	UUID     string `json:"uuid" binding:"required,uuid4"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	SiUnit   string `json:"si_unit"`

	DeviceUUID  string `json:"device_uuid" binding:"required,uuid4"`
	StationUUID string `json:"station_uuid" binding:"required,uuid4"`

	StatusCode string `json:"status_code" binding:"required,oneof='10' '11' '20' '30' '31' '32' '40' '93'"`
	IsEnable   bool   `json:"is_enable"`
	Priority   int    `json:"priority"`

	Source string `json:"source" binding:"required,oneof='sensor' 'device' 'debug'"`

	PhysicalQuantityDataType   string `json:"physical_quantity_data_type"`
	AggregateCalculationMethod string `json:"aggregate_calculation_method"`

	CalibrationEnable    bool   `json:"calibration_enable"`
	CalibrationValue     string `json:"calibration_value" `
	CalibrationParameter string `json:"calibration_parameter" `

	Description string `json:"description"`
}

func (pqc *PhysicalQuantityController) PostCreate(c *gin.Context) {
	var request PhysicalQuantityRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		pqc.response.ValidatorFail(c, paramError)
		return
	}

	err = pqc.physicalQuantityUsecase.Create(domain.PhysicalQuantity{
		UUID:                       request.UUID,
		Name:                       request.Name,
		FullName:                   request.FullName,
		SiUnit:                     request.SiUnit,
		DeviceUUID:                 request.DeviceUUID,
		StationUUID:                request.StationUUID,
		StatusCode:                 request.StatusCode,
		IsEnable:                   request.IsEnable,
		Priority:                   request.Priority,
		Source:                     request.Source,
		PhysicalQuantityDataType:   request.PhysicalQuantityDataType,
		AggregateCalculationMethod: request.AggregateCalculationMethod,
		CalibrationEnable:          request.CalibrationEnable,
		CalibrationValue:           request.CalibrationValue,
		CalibrationParameter:       request.CalibrationParameter,
		Description:                request.Description,
	})
	if err != nil {
		pqc.response.FailWithError(c, createFail, err)
		return
	}

	pqc.response.Success(c, createSuccess)
}

func (pqc *PhysicalQuantityController) PutUpdate(c *gin.Context) {
	var request PhysicalQuantityRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		// FIXME: 刪除
		log.Println(err)
		pqc.response.ValidatorFail(c, paramError)
		return
	}

	err = pqc.physicalQuantityUsecase.Update(domain.PhysicalQuantity{
		UUID:                       request.UUID,
		Name:                       request.Name,
		FullName:                   request.FullName,
		SiUnit:                     request.SiUnit,
		DeviceUUID:                 request.DeviceUUID,
		StationUUID:                request.StationUUID,
		StatusCode:                 request.StatusCode,
		IsEnable:                   request.IsEnable,
		Priority:                   request.Priority,
		Source:                     request.Source,
		PhysicalQuantityDataType:   request.PhysicalQuantityDataType,
		AggregateCalculationMethod: request.AggregateCalculationMethod,
		CalibrationEnable:          request.CalibrationEnable,
		CalibrationValue:           request.CalibrationValue,
		CalibrationParameter:       request.CalibrationParameter,
		Description:                request.Description,
	})
	if err != nil {
		pqc.response.FailWithError(c, updateFail, err)
		return
	}

	pqc.response.Success(c, updateSuccess)
}

func (pqc *PhysicalQuantityController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	err := pqc.physicalQuantityUsecase.Delete(uuid)
	if err != nil {
		pqc.response.FailWithError(c, deleteFail, err)
		return
	}

	pqc.response.Success(c, deleteSuccess)
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
		pqc.response.FailWithError(c, updateFail, err)
		return
	}

	pqc.response.Success(c, updateSuccess)
}
