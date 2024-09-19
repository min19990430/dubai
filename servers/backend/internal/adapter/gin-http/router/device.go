package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type DeviceRouter struct {
	deviceController controller.DeviceController

	jwt jwt.JWT
}

func NewDeviceRouter(deviceController *controller.DeviceController, jwt *jwt.JWT) *DeviceRouter {
	return &DeviceRouter{
		deviceController: *deviceController,
		jwt:              *jwt,
	}
}

func (dr *DeviceRouter) Setup(router *gin.RouterGroup) {
	deviceGroup := router.Group("/v1/Device")
	{
		deviceGroup.GET("", dr.deviceController.GetList)
		deviceGroup.POST("", dr.jwt.Middleware, dr.deviceController.PostCreate)
		deviceGroup.PUT("", dr.jwt.Middleware, dr.deviceController.PutUpdate)
		deviceGroup.DELETE("/:uuid", dr.jwt.Middleware, dr.deviceController.Delete)
	}
}
