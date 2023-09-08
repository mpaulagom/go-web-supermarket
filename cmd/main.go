package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/cmd/handlers"
	"github.com/mpaulagom/go-web-supermarket/internal/product"
)

var (
	filePath = "/Users/mariapaulgom/Documents/go-web-supermarket/products.json"
)

func main() {

	var repoJson = product.NewRepositoryJson(filePath)
	var supermarket = product.NewSuperMarket(repoJson)
	ct := handlers.NewControllerProducts(supermarket, 0)

	/* rt := gin.New()
	// -> middlewares
	//el logger puede tener un timer por ejeplo y por detras darme el tiempo que demora la request
	rt.Use(gin.Logger())
	rt.Use(gin.Recovery()) */
	//Default ya crea los middlewares
	server := gin.Default()

	marketPaths := server.Group("products")
	marketPaths.GET("/", ct.ProductsGet)
	marketPaths.GET("/inmemory", ct.MemoryProductsGet)
	marketPaths.GET("/:id", ct.ProductsGetById)
	marketPaths.GET("/search", ct.ProductsSearch)
	marketPaths.POST("/", ct.SaveProduct())

	server.Run(":8080")
}
