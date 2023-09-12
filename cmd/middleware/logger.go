package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//De esta forma le digo que se ejecute al final de cada request
		ctx.Next()

		method := ctx.Request.Method
		date := time.Now().Local()
		url := ctx.Request.URL.String()
		size := ctx.Request.ContentLength

		log.Printf("[%s] %s - Date: [%s] Request Size: %d", method, url, date, &size)
	}
}
