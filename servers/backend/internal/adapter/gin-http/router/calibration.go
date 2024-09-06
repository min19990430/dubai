package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type CalibrationRouter struct {
	calibrationController controller.CalibrationController
}

func NewCalibrationRouter(calibrationController *controller.CalibrationController) *CalibrationRouter {
	return &CalibrationRouter{
		calibrationController: *calibrationController,
	}
}

func (cr *CalibrationRouter) Setup(router *gin.RouterGroup) {
	calibrationGroup := router.Group("/v1/Calibration")
	{
		calibrationGroup.GET("", cr.calibrationController.GetList)
		calibrationGroup.PATCH("", cr.calibrationController.Update)
	}
}
