package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mpaulagom/go-web-supermarket/internal/domain"
	"github.com/mpaulagom/go-web-supermarket/internal/product"
	"github.com/stretchr/testify/assert"
)

func CreatTestServerForProductsHandler(prodHandler *ControllerProducts) *gin.Engine {
	r := gin.New()
	r.GET("/products", prodHandler.ProductsGet)
	return r
}
func TestFuncional_ProductsHandler_GetByID(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		//arrange
		// -> expectations
		expectedStatusCode := 200
		expectedResponseBody := `[
				{"id":1,"name":"Oil","quantity":439,"code_value":"S8",
					"is_published":true,"expiration":"15/12/2021","price":71.42
				},
				{"id":2,"name":"Pineapple","quantity":345,"code_value":"M4",
					"is_published":true,"expiration":"09/08/2021","price":352.79
				}
			]`
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		// -> setup generamos el repo
		rp := product.RepositoryProductInMemory{
			Products: []*domain.Product{
				{
					Id:          1,
					Name:        "Oil",
					Quantity:    439,
					Code:        "S8",
					IsPublished: true,
					Expiration:  "15/12/2021",
					Price:       71.42,
				},
				{
					Id:          2,
					Name:        "Pineapple",
					Quantity:    345,
					Code:        "M4",
					IsPublished: true,
					Expiration:  "09/08/2021",
					Price:       352.79,
				},
			}}

		//instancio el servicio
		service := product.ProductsService{
			Repository: rp,
		}
		handler := ControllerProducts{
			serviceProduct: service,
		}
		server := CreatTestServerForProductsHandler(&handler)

		// -> Input
		request := httptest.NewRequest("GET", "/products", nil)
		response := httptest.NewRecorder()

		//act
		server.ServeHTTP(response, request)
		//assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, expectedResponseBody, response.Body.String())
		assert.Equal(t, expectedHeaders, response.Header())

	})
}
