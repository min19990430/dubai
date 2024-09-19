package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type ControlSignalRouter struct {
	controlSignalController controller.ControlSignalController
}

func NewControlSignalRouter(controlSignalController *controller.ControlSignalController) *ControlSignalRouter {
	return &ControlSignalRouter{
		controlSignalController: *controlSignalController,
	}
}

func (csr *ControlSignalRouter) Setup(router *gin.RouterGroup) {
	controlSignalGroup := router.Group("/v1/ControlSignal")
	{
		controlSignalGroup.POST("", csr.controlSignalController.PostList)
		controlSignalGroup.PATCH("", csr.controlSignalController.UpdateSignalValue)
	}
}
