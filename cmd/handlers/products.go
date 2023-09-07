package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/internal"
	"github.com/mpaulagom/go-web-supermarket/repository"
)

var (
	filePath = "/Users/mariapaulgom/Documents/go-web-supermarket/products.json"
)
var repoJson = repository.NewRepositoryJson(filePath)
var supermarket = internal.NewSuperMarket(repoJson)

func ProductsGet(ctx *gin.Context) {
	allProducts, err := supermarket.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, allProducts)
}

func ProductsGetById(ctx *gin.Context) {
	product, err := supermarket.GetById(ctx.Param("id"))

	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		default:
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
	}
	ctx.JSON(http.StatusOK, product)
}

func ProductsSearch(ctx *gin.Context) {
	products, err := supermarket.SearchProduct(ctx.Query("priceGt"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)
}
