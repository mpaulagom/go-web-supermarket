package product

import (
	"errors"

	"github.com/mpaulagom/go-web-supermarket/internal/domain"
)

var (
	ErrEmptySupermarket = errors.New("no domain.Products in this supermarket")
	ErrProductNotFound  = errors.New("product not found")
)

type Repository interface {
	// ReadAllData returns all the producrs read from the .json
	ReadAllData() (products []*domain.Product, err error)
	// GetById returns a product by id
	GetById(id int) (*domain.Product, error)
	// SearchProduct search for the products with price greater than the param price
	SearchProduct(price float64) (prs []domain.Product, err error)
	// Update updates a product by its id
	Update(id int, product *domain.Product) error
	// Delete delets a product by its id
	Delete(id int) error
}
