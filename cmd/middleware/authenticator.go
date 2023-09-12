package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/pkg/web"
)

// El token lo saco de las variables de entorno
// Authentication a middleware function to handle the auth token check
func Authentication(token string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tkn := ctx.GetHeader("Authorization")
		if tkn != token {
			//ctx.JSON(http.StatusUnauthorized, )
			web.FailureResponse(ctx, errors.New("not the right token"), http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		//le digo que siga con el siguiente handler
		ctx.Next()
	}
}
