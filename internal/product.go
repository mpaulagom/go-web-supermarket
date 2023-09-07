package internal

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/mpaulagom/go-web-supermarket/repository"
)

var (
	ErrEmptySupermarket = errors.New("no products in this supermarket")
	ErrProductNotFound  = errors.New("product not found")
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type SuperMarket struct {
	Products []Product
}

func LoadSuperMarket(filePath string) (sp SuperMarket) {
	sp = SuperMarket{}
	jsonProducts, err := repository.LoadData(filePath)
	if err != nil {
		return
	}
	json.Unmarshal(jsonProducts, &sp.Products)
	// = jsonProducts.([]Product)
	return
}
func (s SuperMarket) GetAllProducts() (p []Product) {
	p = s.Products
	return
}
func (s SuperMarket) GetById(id string) (product Product, err error) {
	//jsonProducts, err := repository.LoadData()
	if len(s.Products) == 0 {
		err = ErrEmptySupermarket
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	//products := jsonProducts.([]Product)
	for _, pr := range s.Products {
		if pr.Id == idInt {
			product = pr
			return
		}
	}
	err = ErrProductNotFound
	return
}

func (s SuperMarket) SearchProduct(priceS string) (prs []Product, err error) {
	if len(s.Products) == 0 {
		err = ErrEmptySupermarket
		return
	}

	price, err := strconv.ParseFloat(priceS, 64)
	if err != nil {
		return
	}
	for _, pr := range s.Products {
		if pr.Price > price {
			prs = append(prs, pr)
		}
	}
	return
}
