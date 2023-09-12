package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mpaulagom/go-web-supermarket/cmd/handlers"
	"github.com/mpaulagom/go-web-supermarket/cmd/middleware"
	"github.com/mpaulagom/go-web-supermarket/internal/product"
	"github.com/mpaulagom/go-web-supermarket/pkg/store"
)

var (
	filePath = "/Users/mariapaulgom/Documents/go-web-supermarket/products.json"
)

// swaggo docs
// @tittle Products API
func main() {
	//Agarra el archivo .env y hace el os.setEnv con mis clave valor
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	fpath := os.Getenv("JSON_FILEPATH")
	token := os.Getenv("TOKEN")
	var repoJson = store.NewProductStorage(fpath)
	var productService = product.NewProductsService(repoJson)
	productHandler := handlers.NewControllerProducts(productService, 0, token)

	/* APUNTE:
	rt := gin.New()
	-> middlewares
	el logger puede tener un timer por ejeplo y por detras darme el tiempo que demora la request
		rt.Use(gin.Logger())
		rt.Use(gin.Recovery())
	Default ya crea los middlewares, pero de esta forma lo puedo manejar mas manual
	*/

	server := gin.Default()

	productPaths := server.Group("products")
	// insert the middleware definition before any routes
	productPaths.Use(middleware.Logger())
	productPaths.Use(middleware.Authentication(token))
	productPaths.GET("/", productHandler.ProductsGet)
	productPaths.GET("/inmemory", productHandler.MemoryProductsGet)
	productPaths.GET("/:id", productHandler.ProductsGetById)
	productPaths.PUT("/:id", productHandler.Update)
	productPaths.DELETE("/:id", productHandler.Delete)
	productPaths.GET("/search", productHandler.ProductsSearch)
	productPaths.POST("/", productHandler.SaveProduct())

	server.Run(":8080")
}
