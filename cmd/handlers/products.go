package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/internal"
)

// ControllerProducts is an struct that represents a controller for products
// exposing methods to handle products
type ControllerProducts struct {
	//db     []*repository.Product
	sp     *internal.SuperMarket
	lastId int
}

func NewControllerProducts(sp *internal.SuperMarket, lastId int) *ControllerProducts {
	return &ControllerProducts{
		sp:     sp,
		lastId: lastId,
	}
}

func (c *ControllerProducts) ProductsGet(ctx *gin.Context) {
	allProducts, err := c.sp.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, allProducts)
}

func (c *ControllerProducts) ProductsGetById(ctx *gin.Context) {
	product, err := c.sp.GetById(ctx.Param("id"))

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

func (c *ControllerProducts) ProductsSearch(ctx *gin.Context) {
	products, err := c.sp.SearchProduct(ctx.Query("priceGt"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)
}
