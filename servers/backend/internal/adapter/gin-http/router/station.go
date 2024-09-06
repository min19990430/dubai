package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type StationRouter struct {
	stationController controller.StationController
}

func NewStationRouter(stationController *controller.StationController) *StationRouter {
	return &StationRouter{
		stationController: *stationController,
	}
}

func (sr *StationRouter) Setup(router *gin.RouterGroup) {
	stationGroup := router.Group("/v1/Station")
	{
		stationGroup.GET("", sr.stationController.GetList)
	}
}
