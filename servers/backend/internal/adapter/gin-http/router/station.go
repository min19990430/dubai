package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type StationRouter struct {
	stationController controller.StationController

	jwt jwt.JWT
}

func NewStationRouter(stationController *controller.StationController, jwt *jwt.JWT) *StationRouter {
	return &StationRouter{
		stationController: *stationController,
		jwt:               *jwt,
	}
}

func (sr *StationRouter) Setup(router *gin.RouterGroup) {
	stationGroup := router.Group("/v1/Station")
	{
		stationGroup.GET("", sr.stationController.GetList)
		stationGroup.POST("", sr.jwt.Middleware, sr.stationController.PostCreate)
		stationGroup.PUT("", sr.jwt.Middleware, sr.stationController.PutUpdate)
		stationGroup.DELETE("/:uuid", sr.jwt.Middleware, sr.stationController.Delete)
	}
}
