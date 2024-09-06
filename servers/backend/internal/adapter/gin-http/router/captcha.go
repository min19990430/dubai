package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
)

type CaptchaRouter struct {
	captchaController controller.CaptchaController
}

func NewCaptchaRouter(captchaController *controller.CaptchaController) *CaptchaRouter {
	return &CaptchaRouter{
		captchaController: *captchaController,
	}
}

func (cr *CaptchaRouter) Setup(router *gin.RouterGroup) {
	captchaGroup := router.Group("/v1/Captcha")
	{
		captchaGroup.GET("", cr.captchaController.NewDigit)
	}
}
