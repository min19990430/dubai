package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type LastRouter struct {
	lastController controller.LastController
}

func NewLastRouter(lastController *controller.LastController) *LastRouter {
	return &LastRouter{
		lastController: *lastController,
	}
}

func (lr *LastRouter) Setup(router *gin.RouterGroup) {
	lastGroup := router.Group("/v1/Last")
	{
		// lastGroup.GET("", lr.lastController.GetLast)
		lastGroup.GET("/Station", lr.lastController.GetStationLast)
		lastGroup.GET("/Device", lr.lastController.GetDeviceLast)
	}
}
