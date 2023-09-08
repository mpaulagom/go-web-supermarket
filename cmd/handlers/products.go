package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/internal/domain"
	"github.com/mpaulagom/go-web-supermarket/internal/product"
)

// Los modelos reflan las necesidad de la base de datos

//los dtos=objeto de transferencia de datos reflejan las necesidas del usuario de entrada y salida
//Oculto toda la información que me interesa que sea privada

//Por lo general cuando se crea un constructor, se crea una esrtuctura de conifugracion
/* type Config struct {
	Db []*
	LastId in
}
y al constructor del controller le pasamos el Config
*/
// ControllerProducts is an struct that represents a controller for products
// exposing methods to handle products
type ControllerProducts struct {
	//a efectos del post en memoria
	memoryProducts []*domain.Product
	lastId         int
	sp             *product.SuperMarket
}

func NewControllerProducts(sp *product.SuperMarket, lastId int) *ControllerProducts {
	return &ControllerProducts{
		sp:     sp,
		lastId: lastId,
	}
}

// Es buena practica tener estructuras separadas de cada entidad para cada capa
// una entidad para la request, otra para la response, otra para la entidad interna de la logica
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
	Code       string  `json:"code"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}
type ResponseBody struct {
	Message string `json:"message"`
	Data    *Data  `json:"data"`
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
	productE, err := c.sp.GetById(ctx.Param("id"))

	if err != nil {
		switch err {
		case product.ErrProductNotFound:
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		default:
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
	}
	ctx.JSON(http.StatusOK, productE)
}

func (c *ControllerProducts) ProductsSearch(ctx *gin.Context) {
	products, err := c.sp.SearchProduct(ctx.Query("priceGt"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// ShouldBindJSON < > bindJson es mas forzado si llega a fallar tira 400, shouldbind nos da la posibilidad
// de manejar la respuesta a nuestro gusto, si falla
// Saves the product send from the client in memory
func (c *ControllerProducts) SaveProduct() gin.HandlerFunc {
	//al usar el gin.HandlerFunc en vez de poner directo el contexto
	//yo puedo crear variables aca, y usarlas adentro de la funcion que maneja la request.
	return func(ctx *gin.Context) {
		//request
		var req RequestBody
		//asi recuperso cosas de los headers
		token := ctx.GetHeader("Authorization")
		if token != "123" {
			code := http.StatusUnauthorized
			body := &ResponseBody{
				Message: "invalid token",
				Data:    nil,
			}
			ctx.JSON(code, body)
			return
		}
		//caso de error
		if err := ctx.ShouldBindJSON(&req); err != nil {
			code := http.StatusNotFound
			body := &ResponseBody{
				Message: "invalid request body",
				Data:    nil,
			}
			/* Otra forma de enviar respues
			gin.H{
				"error": err.Error(),*/
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
				Id:         pr.Id,
				Name:       pr.Name,
				Quantity:   pr.Quantity,
				Expiration: pr.Expiration,
				Price:      pr.Price,
			},
		}
		ctx.JSON(code, rp)
	}
}

func (ct *ControllerProducts) MemoryProductsGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ct.memoryProducts)
}
