package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type PhysicalQuantityRouter struct {
	physicalQuantityController *controller.PhysicalQuantityController

	jwt jwt.JWT
}

func NewPhysicalQuantityRouter(physicalQuantityController *controller.PhysicalQuantityController, jwt *jwt.JWT) *PhysicalQuantityRouter {
	return &PhysicalQuantityRouter{
		physicalQuantityController: physicalQuantityController,
		jwt:                        *jwt,
	}
}

func (pqr *PhysicalQuantityRouter) Setup(router *gin.RouterGroup) {
	physicalQuantityGroup := router.Group("/v1/PhysicalQuantity")
	{
		physicalQuantityGroup.GET("/:uuid", pqr.physicalQuantityController.GetOne)
		physicalQuantityGroup.GET("", pqr.physicalQuantityController.GetList)
		physicalQuantityGroup.POST("", pqr.jwt.Middleware, pqr.physicalQuantityController.PostCreate)
		physicalQuantityGroup.PUT("", pqr.jwt.Middleware, pqr.physicalQuantityController.PutUpdate)
		physicalQuantityGroup.DELETE("/:uuid", pqr.jwt.Middleware, pqr.physicalQuantityController.Delete)

		physicalQuantityGroup.PATCH("/Status", pqr.physicalQuantityController.UpdateStatus)
	}
}
