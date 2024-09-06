package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"
)

type UserController struct {
	response response.IResponse
	user     usecase.UserUsecase

	loginUsecase usecase.LoginUsecase
}

func NewUserController(response response.IResponse, userUsecase *usecase.UserUsecase, loginUsecase *usecase.LoginUsecase) *UserController {
	return &UserController{
		response:     response,
		user:         *userUsecase,
		loginUsecase: *loginUsecase,
	}
}

func (uc *UserController) List(c *gin.Context) {
	authority := c.GetString("authority_id")
	if authority == "" {
		uc.response.AuthFail(c, accessDenied)
		return
	}

	var isRoot bool
	if authority == "root" {
		isRoot = true
	}

	users, err := uc.user.ListWithRoot(domain.User{}, isRoot)
	if err != nil {
		uc.response.FailWithError(c, queryFail, err)
		return
	}

	uc.response.SuccessWithData(c, querySuccess, users)
}

type UpdateUserPasswordRequest struct {
	UserUUID string `json:"user_uuid" binding:"required,uuid4"`
	Password string `json:"password" binding:"required,password"`
}

func (uc *UserController) UpdatePassword(c *gin.Context) {
	var request UpdateUserPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		uc.response.ValidatorFail(c, paramError)
		return
	}

	err := uc.user.UpdatePassword(domain.User{UUID: request.UserUUID}, request.Password)
	if err != nil {
		uc.response.FailWithError(c, updateFail, err)
		return
	}

	// 更新後，登出被更新密碼的使用者
	err = uc.loginUsecase.Logout(request.UserUUID)
	if err != nil {
		uc.response.FailWithError(c, logoutFail, err)
		return
	}

	uc.response.Success(c, updateSuccess)
}

type UserAuthController struct {
	response response.IResponse
	userAuth usecase.UserAuthUsecase
}

func NewUserAuthController(response response.IResponse, userAuthUsecase *usecase.UserAuthUsecase) *UserAuthController {
	return &UserAuthController{
		response: response,
		userAuth: *userAuthUsecase,
	}
}

func (uac *UserAuthController) GetOne(c *gin.Context) {
	username, ok := c.GetQuery("Username")
	if !ok {
		uac.response.ValidatorFail(c, paramError)
		return
	}

	userAuth, err := uac.userAuth.FindByUser(domain.User{Username: username})
	if err != nil {
		uac.response.FailWithError(c, queryFail, err)
		return
	}

	uac.response.SuccessWithData(c, querySuccess, userAuth)
}
