package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/casbin"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type UserRouter struct {
	userController controller.UserController

	jwt    jwt.JWT
	casbin casbin.Casbin
}

func NewUserRouter(userController *controller.UserController, jwt *jwt.JWT, casbin *casbin.Casbin) *UserRouter {
	return &UserRouter{
		userController: *userController,
		jwt:            *jwt,
		casbin:         *casbin,
	}
}

func (ur *UserRouter) Setup(router *gin.RouterGroup) {
	userGroup := router.Group("/v1/User", ur.jwt.Middleware, ur.casbin.Middleware)
	{
		userGroup.GET("/List", ur.userController.List)
		userGroup.PATCH("/Password", ur.userController.UpdatePassword)
	}
}
