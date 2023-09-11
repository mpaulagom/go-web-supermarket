package web

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/internal/product"
)

type errorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

func ManageErrorResponse(ctx *gin.Context, err error, msg string) *errorResponse {

	if errors.Is(err, strconv.ErrSyntax) {
		errRep := &errorResponse{
			Status:  "Bad Request!!",
			Code:    http.StatusBadRequest,
			Message: msg,
		}
		ctx.JSON(http.StatusBadRequest, errRep)
		return errRep
	}
	switch err {
	case product.ErrProductNotFound:
		errRep := &errorResponse{
			Status:  "Not found",
			Code:    http.StatusNotFound,
			Message: msg,
		}
		ctx.JSON(http.StatusNotFound, errRep)
		return errRep
	default:
		errRep := &errorResponse{
			Status:  "Internal Server Error",
			Code:    http.StatusInternalServerError,
			Message: "",
		}
		ctx.JSON(http.StatusInternalServerError, errRep)
		return errRep
	}
	return nil
}
