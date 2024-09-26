package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type PhysicalQuantityEvaluateRouter struct {
	PhysicalQuantityEvaluateController *controller.PhysicalQuantityEvaluateController

	jwt jwt.JWT
}

func NewPhysicalQuantityEvaluateRouter(physicalQuantityEvaluateController *controller.PhysicalQuantityEvaluateController, jwt *jwt.JWT) *PhysicalQuantityEvaluateRouter {
	return &PhysicalQuantityEvaluateRouter{
		PhysicalQuantityEvaluateController: physicalQuantityEvaluateController,
		jwt:                                *jwt,
	}
}

func (pqer *PhysicalQuantityEvaluateRouter) Setup(router *gin.RouterGroup) {
	physicalQuantityEvaluateGroup := router.Group("/v1/PhysicalQuantityEvaluate")
	{
		physicalQuantityEvaluateGroup.GET("", pqer.PhysicalQuantityEvaluateController.GetList)
		physicalQuantityEvaluateGroup.PUT("/Formula", pqer.jwt.Middleware, pqer.PhysicalQuantityEvaluateController.UpdateFormula)
	}
}
