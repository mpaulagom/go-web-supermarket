package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/internal/domain"
	"github.com/mpaulagom/go-web-supermarket/internal/product"
	"github.com/mpaulagom/go-web-supermarket/pkg/web"
)

/* APUNTE:
	-> Los modelos reflejan las necesidad de la base de datos
    -> Los dtos=objeto de transferencia de datos reflejan las necesidas del usuario de entrada y salida
		oculto toda la información que me interesa que sea privada
	Por lo general cuando se crea un constructor, se crea una estructura de configuracion de la sig forma
 		type Config struct {
			Db []*
			LastId in
		}
	y al constructor del controller le pasamos el Config
*/

// ControllerProducts is an struct that represents a controller for products
type ControllerProducts struct {
	//a efectos del post en memoria
	memoryProducts []*domain.Product
	lastId         int
	serviceProduct product.Service
	token          string
}

// NewControllerPorducts instantiates a ControllerProducts with the dependency of the Service
func NewControllerProducts(sp product.Service, lastId int, token string) *ControllerProducts {
	return &ControllerProducts{
		serviceProduct: sp,
		lastId:         lastId,
		token:          token,
	}
}

type RequestBody struct {
	Name       string  `json:"name" binding:"required"`
	Quantity   int     `json:"quantity"`
	Code       string  `json:"code"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price" binding:"required"`
}
type Data struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}
type ResponseBody struct {
	Message string `json:"message"`
	Data    *Data  `json:"data"`
}

// ProductsGet returns all the products in the repository
// swaggo docs
// @Summary
// @Description
// @Tags products
// @Produce json
// @Success
// @Failure 500 {object}
// @Router /products
func (c *ControllerProducts) ProductsGet(ctx *gin.Context) {
	allProducts, err := c.serviceProduct.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, allProducts)
}

// @Param title path string true "Product identifier"
// ProductsGetById returns a product by its id
func (c *ControllerProducts) ProductsGetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		web.FailureResponse(ctx, errors.New("invalid product identifier"), http.StatusBadRequest)
		return
	}

	// Get the product from the service.
	productE, err := c.serviceProduct.GetById(id)
	if err != nil {
		web.FailureResponse(ctx, errors.New("product not found"), http.StatusNotFound)
		return
	}
	// Return the product.
	ctx.JSON(http.StatusOK, Data{
		Id:       productE.Id,
		Name:     productE.Name,
		Quantity: productE.Quantity,
		Price:    productE.Price,
	})
}

func (c *ControllerProducts) ProductsSearch(ctx *gin.Context) {
	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid price type",
		})
	}
	products, err := c.serviceProduct.SearchProduct(price)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)
}

/*
	APUNTE:
		ShouldBindJSON < > bindJson es mas forzado si llega a fallar tira 400, shouldbind nos da la posibilidad
		de manejar la respuesta a nuestro gusto, si falla
		- al usar el gin.HandlerFunc en vez de poner directo el contexto
			yo puedo crear variables aca, y usarlas adentro de la funcion que maneja la request.
		-  Otra forma de enviar respues
			gin.H {
				"error": err.Error()
			}
*/

// SaveProduct saves the product send from the client in memory
func (c *ControllerProducts) SaveProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req RequestBody

		//caso de error
		if err := ctx.ShouldBindJSON(&req); err != nil {
			code := http.StatusNotFound
			body := &ResponseBody{
				Message: "invalid request body",
				Data:    nil,
			}

			ctx.JSON(code, body)
			return
		}
		//process
		pr := &domain.Product{
			Name:       req.Name,
			Quantity:   req.Quantity,
			Code:       req.Code,
			Expiration: req.Expiration,
			Price:      req.Price,
		}
		pr.Id = c.lastId + 1
		// -> save in storage
		c.memoryProducts = append(c.memoryProducts, pr)
		c.lastId++
		//response
		code := http.StatusCreated
		rp := ResponseBody{
			Message: "Product created",
			Data: &Data{
				Id:       pr.Id,
				Name:     pr.Name,
				Quantity: pr.Quantity,
				Price:    pr.Price,
			},
		}
		ctx.JSON(code, rp)
	}
}

func (ct *ControllerProducts) MemoryProductsGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ct.memoryProducts)
}

func (ct *ControllerProducts) Update(ctx *gin.Context) {
	// Get the ID from the URL.
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product identifier",
		})
		return
	}
	var req RequestBody
	// Bind the request.
	if err := ctx.ShouldBindJSON(&req); err != nil {
		code := http.StatusNotFound
		body := &ResponseBody{
			Message: "invalid request body",
			Data:    nil,
		}
		ctx.JSON(code, body)
		return
	}
	// Prepare valid dto to service layer
	pr := &domain.Product{
		Id:         id,
		Name:       req.Name,
		Quantity:   req.Quantity,
		Code:       req.Code,
		Expiration: req.Expiration,
		Price:      req.Price,
	}
	// Update the product
	ct.serviceProduct.Update(id, pr)
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "product updated correctly",
	})
}

func (ct *ControllerProducts) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product identifier",
		})
		return
	}
	err = ct.serviceProduct.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			// No se debe retornar el error que dio, lo puse provisioriamente para identificar que paso
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "product deleted correctly",
	})
}
