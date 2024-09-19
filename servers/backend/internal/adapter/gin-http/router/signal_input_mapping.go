package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type SignalInputMappingRouter struct {
	signalInputMappingController controller.SignalInputMappingController

	jwt jwt.JWT
}

func NewSignalInputMappingRouter(signalInputMappingController *controller.SignalInputMappingController, jwt *jwt.JWT) *SignalInputMappingRouter {
	return &SignalInputMappingRouter{
		signalInputMappingController: *signalInputMappingController,
		jwt:                          *jwt,
	}
}

func (sir *SignalInputMappingRouter) Setup(router *gin.RouterGroup) {
	signalInputMappingGroup := router.Group("/v1/SignalInputMapping")
	{
		signalInputMappingGroup.POST("", sir.jwt.Middleware, sir.signalInputMappingController.PostCreate)
		signalInputMappingGroup.PUT("", sir.jwt.Middleware, sir.signalInputMappingController.PutUpdate)
		signalInputMappingGroup.DELETE("", sir.jwt.Middleware, sir.signalInputMappingController.Delete)
	}
}

type SignalInputMappingDetailRouter struct {
	signalInputMappingDetailController controller.SignalInputMappingDetailController

	jwt jwt.JWT
}

func NewSignalInputMappingDetailRouter(signalInputMappingDetailController *controller.SignalInputMappingDetailController, jwt *jwt.JWT) *SignalInputMappingDetailRouter {
	return &SignalInputMappingDetailRouter{
		signalInputMappingDetailController: *signalInputMappingDetailController,
		jwt:                                *jwt,
	}
}

func (sir *SignalInputMappingDetailRouter) Setup(router *gin.RouterGroup) {
	signalInputMappingGroup := router.Group("/v1/SignalInputMapping/Detail")
	{
		signalInputMappingGroup.GET("", sir.signalInputMappingDetailController.GetList)
	}
}
