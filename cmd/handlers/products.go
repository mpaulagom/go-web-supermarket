package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/internal"
)

var supermarket internal.SuperMarket = internal.LoadSuperMarket(
	"/Users/mariapaulgom/Documents/go-web-supermarket/products.json")

func ProductsGet(ctx *gin.Context) {
	allProducts := supermarket.GetAllProducts()
	ctx.JSON(http.StatusOK, allProducts)
}

func ProductsGetById(ctx *gin.Context) {
	product, err := supermarket.GetById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
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
