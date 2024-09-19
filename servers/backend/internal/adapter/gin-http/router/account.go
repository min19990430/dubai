package router

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller"
	"auto-monitoring/internal/adapter/gin-http/middleware/jwt"
)

type AccountRouter struct {
	accountController controller.AccountController

	jwt jwt.JWT
}

func NewAccountRouter(accountController *controller.AccountController, jwt *jwt.JWT) *AccountRouter {
	return &AccountRouter{
		accountController: *accountController,
		jwt:               *jwt,
	}
}

func (ar *AccountRouter) Setup(router *gin.RouterGroup) {
	accountGroup := router.Group("/v1/Account")
	{
		accountGroup.GET("/Self", ar.jwt.Middleware, ar.accountController.Self)
	}
}
