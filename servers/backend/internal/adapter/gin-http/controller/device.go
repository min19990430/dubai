package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type DeviceController struct {
	response response.IResponse
	device   usecase.DeviceUsecase
}

func NewDeviceController(response response.IResponse, deviceUsecase *usecase.DeviceUsecase) *DeviceController {
	return &DeviceController{
		response: response,
		device:   *deviceUsecase,
	}
}

func (dc *DeviceController) GetList(c *gin.Context) {
	devices, err := dc.device.List(domain.Device{})
	if err != nil {
		dc.response.FailWithError(c, "fail to get devices", err)
		return
	}

	dc.response.SuccessWithData(c, "success", devices)
}

type DeviceStationController struct {
	response response.IResponse
	deviceStation   usecase.DeviceStationUsecase
}

func NewDeviceStationController(response response.IResponse, deviceStationUsecase *usecase.DeviceStationUsecase) *DeviceStationController {
	return &DeviceStationController{
		response: response,
		deviceStation: *deviceStationUsecase,
	}
}

func (dc *DeviceStationController) GetList(c *gin.Context) {
	devices, err := dc.deviceStation.List(domain.DeviceStation{})
	if err != nil {
		dc.response.FailWithError(c, "fail to get device stations", err)
		return
	}

	dc.response.SuccessWithData(c, "success", devices)
}