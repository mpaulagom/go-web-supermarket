package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/cmd/handlers"
)

func main() {

	server := gin.Default()

	superPaths := server.Group("products")
	superPaths.GET("/", handlers.ProductsGet)
	superPaths.GET("/:id", handlers.ProductsGetById)
	superPaths.GET("/search", handlers.ProductsSearch)

	server.Run(":8080")
}
