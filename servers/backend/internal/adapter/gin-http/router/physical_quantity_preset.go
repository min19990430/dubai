package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type PhysicalQuantityPresetRouter struct {
	physicalQuantityPresetController controller.PhysicalQuantityPresetController

	jwt jwt.JWT
}

func NewPhysicalQuantityPresetRouter(physicalQuantityPresetController *controller.PhysicalQuantityPresetController, jwt *jwt.JWT) *PhysicalQuantityPresetRouter {
	return &PhysicalQuantityPresetRouter{
		physicalQuantityPresetController: *physicalQuantityPresetController,
		jwt:                              *jwt,
	}
}

func (pqpr *PhysicalQuantityPresetRouter) Setup(router *gin.RouterGroup) {
	physicalQuantityPresetGroup := router.Group("/v1/PhysicalQuantityPreset")
	{
		physicalQuantityPresetGroup.GET("/:uuid", pqpr.physicalQuantityPresetController.GetOne)
		physicalQuantityPresetGroup.GET("", pqpr.physicalQuantityPresetController.GetList)
		physicalQuantityPresetGroup.POST("", pqpr.jwt.Middleware, pqpr.physicalQuantityPresetController.PostCreate)
		physicalQuantityPresetGroup.PUT("", pqpr.jwt.Middleware, pqpr.physicalQuantityPresetController.PutUpdate)
		physicalQuantityPresetGroup.DELETE("/:uuid", pqpr.jwt.Middleware, pqpr.physicalQuantityPresetController.Delete)
	}
}
