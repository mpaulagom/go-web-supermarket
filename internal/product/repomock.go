package product

import "github.com/mpaulagom/go-web-supermarket/internal/domain"

type MockRepository struct {
	Products []*domain.Product
}

// No estoy segura de que tan bien esta esto
func (ms *MockRepository) LoadSuperMarket() (products []*domain.Product, err error) {
	products = ms.Products
	return
}
