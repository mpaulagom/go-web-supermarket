package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

// Successful gives a successful response
func SuccessfulResponse(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data: data,
	})
}

// FailureResponse gives an error response
func FailureResponse(ctx *gin.Context, err error, code int) *errorResponse {
	ctx.JSON(code, errorResponse{
		Message: err.Error(),
		Code:    code,
		Status:  http.StatusText(code),
	})
	return nil
}
