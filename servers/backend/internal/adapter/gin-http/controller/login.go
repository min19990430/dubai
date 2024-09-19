package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
)

type LoginController struct {
	response     response.IResponse
	loginUsecase usecase.LoginUsecase
}

func NewLoginController(response response.IResponse, usecase *usecase.LoginUsecase) *LoginController {
	return &LoginController{
		response:     response,
		loginUsecase: *usecase,
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,password"`
}

func (lc *LoginController) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		lc.response.ValidatorFail(c, paramError)
		return
	}

	token, err := lc.loginUsecase.Login(request.Username, request.Password)
	if err != nil {
		lc.response.FailWithError(c, loginFail, err)
		return
	}

	lc.response.SuccessWithData(c, loginSuccess, token)
}

func (lc *LoginController) Logout(c *gin.Context) {
	uuid := c.GetString("uuid")
	if uuid == "" {
		lc.response.AuthFail(c, accessDenied)
		return
	}

	err := lc.loginUsecase.Logout(uuid)
	if err != nil {
		lc.response.FailWithError(c, logoutFail, err)
		return
	}
	lc.response.Success(c, logoutSuccess)
}
