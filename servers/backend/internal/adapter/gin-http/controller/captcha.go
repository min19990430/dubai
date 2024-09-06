package controller

import (
	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
)

type CaptchaController struct {
	response response.IResponse
	captcha  usecase.CaptchaUsecase
}

func NewCaptchaController(response response.IResponse, captchaUsecase *usecase.CaptchaUsecase) *CaptchaController {
	return &CaptchaController{
		response: response,
		captcha:  *captchaUsecase,
	}
}

func (cc *CaptchaController) NewDigit(c *gin.Context) {
	captcha, err := cc.captcha.NewDigit()
	if err != nil {
		cc.response.FailWithError(c, "fail to generate captcha", err)
		return
	}

	cc.response.SuccessWithData(c, "success", captcha)
}
