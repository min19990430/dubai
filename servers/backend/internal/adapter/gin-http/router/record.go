package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/logger"
)

type RecordRouter struct {
	recordController controller.RecordController

	logger logger.Logger
}

func NewRecordRouter(recordController *controller.RecordController, logger logger.Logger) *RecordRouter {
	return &RecordRouter{
		recordController: *recordController,
		logger:           logger,
	}
}

func (rr *RecordRouter) Setup(router *gin.RouterGroup) {
	recordGroup := router.Group("/v1/Record")
	{
		recordGroup.POST("/Last", rr.recordController.LastRecord)
		recordGroup.POST("", rr.recordController.PostList)
		recordGroup.POST("/Device", rr.recordController.PostListDevice)
		recordGroup.POST("/Device/JSON", rr.recordController.PostListDeviceJSON)
		recordGroup.POST("/Station", rr.recordController.PostListStation)
		recordGroup.POST("/Station/JSON", rr.logger.Middleware, rr.recordController.PostListStationJSON)
	}
}
