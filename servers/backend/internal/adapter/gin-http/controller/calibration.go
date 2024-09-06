package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
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

func (cc *CalibrationController) GetList(c *gin.Context) {
	deviceUUID, hasDeviceID := c.GetQuery("DeviceUUID")
	if !hasDeviceID {
		cc.response.ValidatorFail(c, "DeviceUUID is required")
		return
	}

	validate := validator.New()
	if validateErr := validate.Var(deviceUUID, "required,uuid"); validateErr != nil {
		cc.response.ValidatorFail(c, validatorFail)
		return
	}

	result, err := cc.calibration.TrialCalculator(deviceUUID)
	if err != nil {
		cc.response.FailWithError(c, queryFail, err)
		return
	}

	device, err := cc.device.FindByUUID(deviceUUID)
	if err != nil {
		cc.response.FailWithError(c, queryFail, err)
		return
	}

	station, err := cc.station.FindByUUID(device.StationUUID)
	if err != nil {
		cc.response.FailWithError(c, queryFail, err)
		return
	}

	cc.response.SuccessWithData(c, "success",
		gin.H{
			"station":     station,
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
