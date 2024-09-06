package jwt

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	"auto-monitoring/internal/adapter/gin-http/controller/response"
	"auto-monitoring/internal/application/usecase"
)

type JWT struct {
	response response.IResponse
	token    usecase.TokenUsecase
}

func NewJWT(response response.IResponse, token *usecase.TokenUsecase) *JWT {
	return &JWT{
		response: response,
		token:    *token,
	}
}

func (j *JWT) Middleware(c *gin.Context) {
	token, err := j.extractToken(c)
	if err != nil {
		j.response.AuthFail(c, "no token")
		c.Abort()
		return
	}

	tokenClaims, err := j.token.Verify(token)
	if err != nil {
		j.response.AuthFail(c, "invalid token")
		c.Abort()
		return
	}

	c.Set("uuid", tokenClaims.UserUUID)
	c.Set("username", tokenClaims.Username)
	c.Set("fullname", tokenClaims.FullName)
	c.Set("nickname", tokenClaims.NickName)
	c.Set("authority_id", tokenClaims.AuthorityID)
	c.Next()
}

func (JWT) extractToken(c *gin.Context) (string, error) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		return "", errors.New("no token")
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", errors.New("invalid authorization format")
	}

	token := strings.TrimPrefix(authorization, "Bearer ")
	return token, nil
}
