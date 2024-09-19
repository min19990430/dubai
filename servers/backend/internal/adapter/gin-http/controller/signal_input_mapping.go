package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type SignalInputMappingController struct {
	response response.IResponse

	signalInputMapping usecase.SignalInputMappingUsecase
}

func NewSignalInputMappingController(response response.IResponse, usecase *usecase.SignalInputMappingUsecase) *SignalInputMappingController {
	return &SignalInputMappingController{
		response:           response,
		signalInputMapping: *usecase,
	}
}

type SignalInputMappingRequest struct {
	DeviceUUID                 string `json:"device_uuid" binding:"required,uuid"`
	Type                       string `json:"type" binding:"required"`
	Pointer                    int    `json:"pointer" `
	TargetPhysicalQuantityUUID string `json:"target_physical_quantity_uuid" binding:"required,uuid"`
}

func (simc *SignalInputMappingController) PostCreate(c *gin.Context) {
	var req SignalInputMappingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		simc.response.ValidatorFail(c, paramError)
		return
	}

	signalInputMapping := domain.SignalInputMapping{
		DeviceUUID:                 req.DeviceUUID,
		Type:                       req.Type,
		Pointer:                    req.Pointer,
		TargetPhysicalQuantityUUID: req.TargetPhysicalQuantityUUID,
	}

	if err := simc.signalInputMapping.Create(signalInputMapping); err != nil {
		simc.response.FailWithError(c, createFail, err)
		return
	}

	simc.response.Success(c, createSuccess)
}

type PutSignalInputMappingRequest struct {
	Old SignalInputMappingRequest `json:"old"`
	New SignalInputMappingRequest `json:"new"`
}

func (simc *SignalInputMappingController) PutUpdate(c *gin.Context) {
	var req PutSignalInputMappingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		simc.response.ValidatorFail(c, paramError)
		return
	}

	oldMapping := domain.SignalInputMapping{
		DeviceUUID:                 req.Old.DeviceUUID,
		Type:                       req.Old.Type,
		Pointer:                    req.Old.Pointer,
		TargetPhysicalQuantityUUID: req.Old.TargetPhysicalQuantityUUID,
	}

	newMapping := domain.SignalInputMapping{
		DeviceUUID:                 req.New.DeviceUUID,
		Type:                       req.New.Type,
		Pointer:                    req.New.Pointer,
		TargetPhysicalQuantityUUID: req.New.TargetPhysicalQuantityUUID,
	}

	if err := simc.signalInputMapping.Update(oldMapping, newMapping); err != nil {
		simc.response.FailWithError(c, updateFail, err)
		return
	}

	simc.response.Success(c, updateSuccess)
}

func (simc *SignalInputMappingController) Delete(c *gin.Context) {
	var req SignalInputMappingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		simc.response.ValidatorFail(c, paramError)
		return
	}

	signalInputMapping := domain.SignalInputMapping{
		DeviceUUID:                 req.DeviceUUID,
		Type:                       req.Type,
		Pointer:                    req.Pointer,
		TargetPhysicalQuantityUUID: req.TargetPhysicalQuantityUUID,
	}

	if err := simc.signalInputMapping.Delete(signalInputMapping); err != nil {
		simc.response.FailWithError(c, deleteFail, err)
		return
	}

	simc.response.Success(c, deleteSuccess)
}

type SignalInputMappingDetailController struct {
	response response.IResponse

	signalInputMappingDetail usecase.SignalInputMappingDetailUsecase
}

func NewSignalInputMappingDetailController(response response.IResponse, usecase *usecase.SignalInputMappingDetailUsecase) *SignalInputMappingDetailController {
	return &SignalInputMappingDetailController{
		response:                 response,
		signalInputMappingDetail: *usecase,
	}
}

func (simdc *SignalInputMappingDetailController) GetList(c *gin.Context) {
	var signalInputMapping domain.SignalInputMapping

	deviceUUID, isDeviceUUIDExist := c.GetQuery("device_uuid")
	if isDeviceUUIDExist {
		signalInputMapping.DeviceUUID = deviceUUID
	}

	signalInputMappingDetail, err := simdc.signalInputMappingDetail.List(signalInputMapping)
	if err != nil {
		simdc.response.FailWithError(c, queryFail, err)
		return
	}

	simdc.response.SuccessWithData(c, querySuccess, signalInputMappingDetail)
}
