package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type AlarmSettingRouter struct {
	alarmSettingController controller.AlarmSettingController
}

func NewAlarmSettingRouter(alarmSettingController *controller.AlarmSettingController) *AlarmSettingRouter {
	return &AlarmSettingRouter{
		alarmSettingController: *alarmSettingController,
	}
}

func (as *AlarmSettingRouter) Setup(router *gin.RouterGroup) {
	alarmSettingGroup := router.Group("/v1/Alarm/Setting")
	{
		alarmSettingGroup.PATCH("/Expression", as.alarmSettingController.UpdateExpression)
		alarmSettingGroup.GET("", as.alarmSettingController.ListByDeviceUUID)
		alarmSettingGroup.GET("/Station", as.alarmSettingController.ListByStationUUID)
	}
}
