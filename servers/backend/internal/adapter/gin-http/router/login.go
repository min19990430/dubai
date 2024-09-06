package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/captcha"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type LoginRouter struct {
	loginController controller.LoginController

	captcha captcha.Captcha
	jwt     jwt.JWT
}

func NewLoginRouter(loginController *controller.LoginController, captcha *captcha.Captcha, jwt *jwt.JWT) *LoginRouter {
	return &LoginRouter{
		loginController: *loginController,
		captcha:         *captcha,
		jwt:             *jwt,
	}
}

func (lr *LoginRouter) Setup(router *gin.RouterGroup) {
	loginGroup := router.Group("/v1")
	{
		loginGroup.POST("/Login", lr.captcha.Middleware, lr.loginController.Login)
		loginGroup.POST("/Logout", lr.jwt.Middleware, lr.loginController.Logout)
	}
}
