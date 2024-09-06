package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type TimeSeriesRouter struct {
	timeSeriesController controller.TimeSeriesController
}

func NewTimeSeriesRouter(timeSeriesController *controller.TimeSeriesController) *TimeSeriesRouter {
	return &TimeSeriesRouter{
		timeSeriesController: *timeSeriesController,
	}
}

func (tsr *TimeSeriesRouter) Setup(router *gin.RouterGroup) {
	timeSeriesGroup := router.Group("/v1/TimeSeries")
	{
		timeSeriesGroup.POST("/Station", tsr.timeSeriesController.PostListByStation)
		timeSeriesGroup.POST("/Station/JSON", tsr.timeSeriesController.PostListJSONByStation)
	}
}
