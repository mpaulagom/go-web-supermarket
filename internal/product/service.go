package product

import (
	"errors"
	"strconv"

	"github.com/mpaulagom/go-web-supermarket/internal/domain"
)

var (
	ErrEmptySupermarket = errors.New("no domain.Products in this supermarket")
	ErrProductNotFound  = errors.New("product not found")
)

func NewSuperMarket(rj Repository) *SuperMarket {
	return &SuperMarket{rj: rj}
}

type SuperMarket struct {
	rj Repository
}

/*
	 func LoadSuperMarket(filePath string) (sp SuperMarket) {
		sp = SuperMarket{}
		jsonProducts, err := LoadData(filePath)
		if err != nil {
			return
		}
		json.Unmarshal(jsonProducts, &sp.Products)
		// = jsonProducts.([]Product)
		return
	}
*/
func (s SuperMarket) GetAllProducts() (products []*domain.Product, err error) {
	products, err = s.rj.LoadSuperMarket()
	return
}

func (s SuperMarket) GetById(id string) (product domain.Product, err error) {
	products, err := s.rj.LoadSuperMarket()
	if err != nil {
		return
	}
	if len(products) == 0 {
		err = ErrEmptySupermarket
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	//products := jsonProducts.([]Product)
	for _, pr := range products {
		if pr.Id == idInt {
			product = *pr
			return
		}
	}
	err = ErrProductNotFound
	return
}

func (s SuperMarket) SearchProduct(priceS string) (prs []domain.Product, err error) {
	products, err := s.rj.LoadSuperMarket()
	if err != nil {
		return
	}
	if len(products) == 0 {
		err = ErrEmptySupermarket
		return
	}

	price, err := strconv.ParseFloat(priceS, 64)
	if err != nil {
		return
	}
	for _, pr := range products {
		if pr.Price > price {
			prs = append(prs, *pr)
		}
	}
	return
}
