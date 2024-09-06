package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type PhysicalQuantityRouter struct {
	physicalQuantityController controller.PhysicalQuantityController
}

func NewPhysicalQuantityRouter(physicalQuantityController *controller.PhysicalQuantityController) *PhysicalQuantityRouter {
	return &PhysicalQuantityRouter{
		physicalQuantityController: *physicalQuantityController,
	}
}

func (pqr *PhysicalQuantityRouter) Setup(router *gin.RouterGroup) {
	physicalQuantityGroup := router.Group("/v1/PhysicalQuantity")
	{
		physicalQuantityGroup.PATCH("/Status", pqr.physicalQuantityController.UpdateStatus)
	}
}
