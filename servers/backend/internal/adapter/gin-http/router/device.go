package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type DeviceRouter struct {
	deviceController controller.DeviceController
}

func NewDeviceRouter(deviceController *controller.DeviceController) *DeviceRouter {
	return &DeviceRouter{
		deviceController: *deviceController,
	}
}

func (dr *DeviceRouter) Setup(router *gin.RouterGroup) {
	deviceGroup := router.Group("/v1/Device")
	{
		deviceGroup.GET("", dr.deviceController.GetList)
	}
}

type DeviceStationRouter struct {
	deviceStationController controller.DeviceStationController
}

func NewDeviceStationRouter(deviceStationController *controller.DeviceStationController) *DeviceStationRouter {
	return &DeviceStationRouter{
		deviceStationController: *deviceStationController,
	}
}

func (dr *DeviceStationRouter) Setup(router *gin.RouterGroup) {
	deviceStationGroup := router.Group("/v1/Device/Station")
	{
		deviceStationGroup.GET("", dr.deviceStationController.GetList)
	}
}
