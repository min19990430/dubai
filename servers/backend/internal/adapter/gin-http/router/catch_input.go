package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type CatchInputRouter struct {
	catchInputController controller.CatchInputController
}

func NewCatchInputRouter(catchInputController *controller.CatchInputController) *CatchInputRouter {
	return &CatchInputRouter{
		catchInputController: *catchInputController,
	}
}

func (cir CatchInputRouter) Setup(router *gin.RouterGroup) {
	catchInputGroup := router.Group("/v1/Catch")
	{
		catchInputGroup.POST("/Signal", cir.catchInputController.CatchSignalInput)
	}
}
