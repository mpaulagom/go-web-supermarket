package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/cmd/handlers"
	"github.com/mpaulagom/go-web-supermarket/internal"
	"github.com/mpaulagom/go-web-supermarket/repository"
)

var (
	filePath = "/Users/mariapaulgom/Documents/go-web-supermarket/products.json"
)

func main() {

	var repoJson = repository.NewRepositoryJson(filePath)
	var supermarket = internal.NewSuperMarket(repoJson)
	ct := handlers.NewControllerProducts(supermarket, 0)
	server := gin.Default()

	superPaths := server.Group("products")
	superPaths.GET("/", ct.ProductsGet)
	superPaths.GET("/:id", ct.ProductsGetById)
	superPaths.GET("/search", ct.ProductsSearch)

	server.Run(":8080")
}
