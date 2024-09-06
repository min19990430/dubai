package response

import (
	"github.com/gin-gonic/gin"
)

type IResponse interface {
	// Success returns a HTTP 200 response with the given message
	Success(c *gin.Context, message string)

	// SuccessWithData returns a HTTP 200 response with the given message and data
	SuccessWithData(c *gin.Context, message string, data any)

	// Fail returns a HTTP 500 response with the given message
	Fail(c *gin.Context, message string)

	// FailWithError returns a HTTP 500 response with the given message and error
	FailWithError(c *gin.Context, message string, err error)

	// ValidatorFail returns a HTTP 400 response with the given message
	ValidatorFail(c *gin.Context, message string)

	// AuthFail returns a HTTP 401 response with the given message
	AuthFail(c *gin.Context, message string)
}
