package captcha

import (
	"bytes"
	"encoding/json"
	"io"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
	"auto-monitoring/internal/domain"

	"github.com/gin-gonic/gin"
)

type Captcha struct {
	response response.IResponse

	captcha usecase.CaptchaUsecase
}

func NewCaptcha(response response.IResponse, captcha *usecase.CaptchaUsecase) *Captcha {
	return &Captcha{
		response: response,
		captcha:  *captcha,
	}
}

type Request struct {
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

func (cc *Captcha) Middleware(c *gin.Context) {
	body, getErr := c.GetRawData()
	if getErr != nil {
		cc.response.AuthFail(c, "body is empty")
		c.Abort()
		return
	}

	var request Request
	unmarshalErr := json.Unmarshal(body, &request)
	if unmarshalErr != nil {
		cc.response.AuthFail(c, "captcha not found")
		c.Abort()
		return
	}

	captchaPO := domain.Captcha{
		CaptchaID:     request.CaptchaID,
		CaptchaAnswer: request.Captcha,
	}

	if !cc.captcha.Verify(captchaPO) {
		cc.response.AuthFail(c, "captcha is wrong")
		c.Abort()
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	c.Next()
}
