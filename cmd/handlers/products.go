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
// @Summary Gets all the products
// @Description Get the list of all the available products
// @Tags products
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /products [get]
func (c *ControllerProducts) ProductsGet(ctx *gin.Context) {
	allProducts, err := c.serviceProduct.GetAllProducts()
	if err != nil {
		web.FailureResponse(ctx, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	web.SuccessfulResponse(ctx, http.StatusOK, allProducts)
}

// ShowAccount godoc
// @Summary Gets a product by id
// @Description Returns a product that matches the id
// @Tags products
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Product id"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /products/{id} [get]
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
	web.SuccessfulResponse(ctx, http.StatusOK, Data{
		Id:       productE.Id,
		Name:     productE.Name,
		Quantity: productE.Quantity,
		Price:    productE.Price,
	})
}

// ShowAccount godoc
// @Summary Gets products with price bigger than a given price
// @Description Returns all the products that have a price equal or greater than the given price
// @Tags products
// @Produce json
// @Param token header string true "token"
// @Param priceGt query float64 true "Product price"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /search [get]
func (c *ControllerProducts) ProductsSearch(ctx *gin.Context) {
	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		web.FailureResponse(ctx, errors.New("invalid price type"), http.StatusBadRequest)
	}
	products, err := c.serviceProduct.SearchProduct(price)
	if err != nil {
		web.FailureResponse(ctx, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	web.SuccessfulResponse(ctx, http.StatusOK, products)
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
// ShowAccount godoc
// @Summary Store a product
// @Description Creates a new product to be saved in the repository
// @Tags products
// @Produce json
// @Param token header string true "token"
// @Param product body RequestBody true "Product to save"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /products [post]
func (c *ControllerProducts) SaveProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req RequestBody

		//caso de error
		if err := ctx.ShouldBindJSON(&req); err != nil {
			code := http.StatusBadRequest
			web.FailureResponse(ctx, errors.New("invalid request body"), code)
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
		web.SuccessfulResponse(ctx, code, rp)
	}
}

/* func (ct *ControllerProducts) MemoryProductsGet(ctx *gin.Context) {
	web.SuccessfulResponse(ctx, http.StatusOK, ct.memoryProducts)
} */

// ShowAccount godoc
// @Summary Updates product of given id
// @Description updates product of given id
// @Tags products
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Product id"
// @Param product body RequestBody true "Product to update"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /products/{id} [put]
func (ct *ControllerProducts) Update(ctx *gin.Context) {
	// Get the ID from the URL.
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		web.FailureResponse(ctx, errors.New("invalid product identifier"), http.StatusBadRequest)
		return
	}
	var req RequestBody
	// Bind the request.
	if err := ctx.ShouldBindJSON(&req); err != nil {
		web.FailureResponse(ctx, errors.New("invalid request body"), http.StatusBadRequest)
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
	web.SuccessfulResponse(ctx, http.StatusOK, gin.H{
		"Message": "product updated correctly",
	})
}

// ShowAccount godoc
// @Summary Deletes product of given id
// @Description Deletes product of given id
// @Tags products
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Product id"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /products/{id} [delete]
func (ct *ControllerProducts) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		web.FailureResponse(ctx, errors.New("invalid product identifier"), http.StatusBadRequest)
		return
	}
	err = ct.serviceProduct.Delete(id)
	if err != nil {
		web.FailureResponse(ctx, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	ct.serviceProduct.Delete(id)
	web.SuccessfulResponse(ctx, http.StatusOK, gin.H{
		"Message": "product updated correctly",
	})
}
