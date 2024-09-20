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
		dc.response.FailWithError(c, queryFail, err)
		return
	}

	dc.response.SuccessWithData(c, querySuccess, devices)
}

type DeviceRequest struct {
	UUID        string  `json:"uuid" binding:"required,uuid"`
	ID          string  `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	IsEnable    bool    `json:"is_enable" `
	Priority    int     `json:"priority"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Description string  `json:"description"`
}

func (dc *DeviceController) PostCreate(c *gin.Context) {
	var req DeviceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dc.response.ValidatorFail(c, paramError)
		return
	}

	device := domain.Device{
		UUID:        req.UUID,
		ID:          req.ID,
		Name:        req.Name,
		IsEnable:    req.IsEnable,
		Priority:    req.Priority,
		Lat:         req.Lat,
		Lon:         req.Lon,
		Description: req.Description,
	}

	if err := dc.device.Create(device); err != nil {
		dc.response.FailWithError(c, createFail, err)
		return
	}

	dc.response.Success(c, createSuccess)
}

func (dc *DeviceController) PutUpdate(c *gin.Context) {
	var req DeviceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		dc.response.ValidatorFail(c, paramError)
		return
	}

	device := domain.Device{
		UUID:        req.UUID,
		ID:          req.ID,
		Name:        req.Name,
		IsEnable:    req.IsEnable,
		Priority:    req.Priority,
		Lat:         req.Lat,
		Lon:         req.Lon,
		Description: req.Description,
	}

	if err := dc.device.Update(device); err != nil {
		dc.response.FailWithError(c, updateFail, err)
		return
	}

	dc.response.Success(c, updateSuccess)
}

func (dc *DeviceController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	if err := dc.device.Delete(uuid); err != nil {
		dc.response.FailWithError(c, deleteFail, err)
		return
	}

	dc.response.Success(c, deleteSuccess)
}
