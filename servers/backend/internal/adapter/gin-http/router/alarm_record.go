package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type AlarmRecordRouter struct {
	alarmRecordController controller.AlarmRecordController
}

func NewAlarmRecordRouter(alarmRecordController *controller.AlarmRecordController) *AlarmRecordRouter {
	return &AlarmRecordRouter{
		alarmRecordController: *alarmRecordController,
	}
}

func (ar *AlarmRecordRouter) Setup(router *gin.RouterGroup) {
	alarmRecordGroup := router.Group("/v1/Alarm")
	{
		alarmRecordGroup.POST("", ar.alarmRecordController.PostList)
		alarmRecordGroup.POST("/Device", ar.alarmRecordController.PostListByDeviceUUID)
		alarmRecordGroup.POST("/Station", ar.alarmRecordController.PostListByStationUUID)
		alarmRecordGroup.POST("/Detail", ar.alarmRecordController.PostListDetail)
	}
}
