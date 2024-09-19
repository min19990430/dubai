package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
)

type AccountController struct {
	response response.IResponse
}

func NewAccountController(response response.IResponse) *AccountController {
	return &AccountController{
		response: response,
	}
}

func (ac *AccountController) Self(c *gin.Context) {
	uuid, exists := c.Get("uuid")
	if !exists {
		ac.response.AuthFail(c, "no uuid")
		return
	}

	username, exists := c.Get("username")
	if !exists {
		ac.response.AuthFail(c, "no username")
		return
	}

	fullname, exists := c.Get("fullname")
	if !exists {
		ac.response.AuthFail(c, "no fullname")
		return
	}

	nickname, exists := c.Get("nickname")
	if !exists {
		ac.response.AuthFail(c, "no nickname")
		return
	}

	authorityID, exists := c.Get("authority_id")
	if !exists {
		ac.response.AuthFail(c, "no authority_id")
		return
	}

	ac.response.SuccessWithData(c, querySuccess, gin.H{
		"uuid":         uuid,
		"username":     username,
		"fullname":     fullname,
		"nickname":     nickname,
		"authority_id": authorityID,
	})
}
