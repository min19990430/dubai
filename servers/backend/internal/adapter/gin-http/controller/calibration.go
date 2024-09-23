package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type CalibrationController struct {
	response    response.IResponse
	calibration usecase.CalibrationUsecase

	station usecase.StationUsecase
	device  usecase.DeviceUsecase
}

func NewCalibrationController(
	response response.IResponse,
	calibration *usecase.CalibrationUsecase,
	station *usecase.StationUsecase,
	device *usecase.DeviceUsecase,
) *CalibrationController {
	return &CalibrationController{
		response:    response,
		calibration: *calibration,
		station:     *station,
		device:      *device,
	}
}

type CalibrationListRequest struct {
	DeviceUUID string `form:"device_uuid" binding:"required,uuid4"`
	Source     string `form:"source" binding:"required,oneof='sensor' 'device' 'debug'"`
}

func (cc *CalibrationController) GetList(c *gin.Context) {
	var request CalibrationListRequest
	if err := c.ShouldBind(&request); err != nil {
		cc.response.ValidatorFail(c, paramError)
		return
	}

	result, err := cc.calibration.TrialCalculator(domain.PhysicalQuantity{IsEnable: true, DeviceUUID: request.DeviceUUID, Source: request.Source})
	if err != nil {
		cc.response.FailWithError(c, queryFail, err)
		return
	}

	device, err := cc.device.FindByUUID(request.DeviceUUID)
	if err != nil {
		cc.response.FailWithError(c, queryFail, err)
		return
	}

	cc.response.SuccessWithData(c, querySuccess,
		gin.H{
			"device":      device,
			"calibration": result,
		})
}

type CalibrationRequest struct {
	UUID                 string `json:"uuid" binding:"required,uuid4"`
	CalibrationEnable    bool   `json:"calibration_enable"`
	CalibrationValue     string `json:"calibration_value" binding:"required,five_point"`
	CalibrationParameter string `json:"calibration_parameter" binding:"required,five_point"`
}

func (cc *CalibrationController) Update(c *gin.Context) {
	var request CalibrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		cc.response.ValidatorFail(c, paramError)
		return
	}

	if err := cc.calibration.Update(
		request.UUID,
		request.CalibrationEnable,
		request.CalibrationValue,
		request.CalibrationParameter); err != nil {
		cc.response.FailWithError(c, updateFail, err)
		return
	}

	cc.response.Success(c, "success")
}
