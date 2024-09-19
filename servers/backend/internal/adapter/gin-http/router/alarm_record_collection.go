package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type AlarmRecordCollectionRouter struct {
	alarmRecordCollectionController controller.AlarmRecordCollectionController
}

func NewAlarmRecordCollectionRouter(alarmRecordCollectionController *controller.AlarmRecordCollectionController) *AlarmRecordCollectionRouter {
	return &AlarmRecordCollectionRouter{
		alarmRecordCollectionController: *alarmRecordCollectionController,
	}
}

func (arc *AlarmRecordCollectionRouter) Setup(router *gin.RouterGroup) {
	alarmRecordCollectionGroup := router.Group("/v1/Alarm/Collection")
	{
		alarmRecordCollectionGroup.POST("/Device", arc.alarmRecordCollectionController.PostListByDeviceUUID)
		alarmRecordCollectionGroup.POST("/Station", arc.alarmRecordCollectionController.PostListByStationUUID)
	}
}
