package product

import (
	"github.com/mpaulagom/go-web-supermarket/internal/domain"
)

type Service interface {
	GetById(id int) (*domain.Product, error)
	Update(id int, product *domain.Product) error
	Delete(id int) error
	GetAllProducts() (products []*domain.Product, err error)
	SearchProduct(priceS float64) (prs []domain.Product, err error)
}

// NewProductsService returns a new service with the dependency to repository
func NewProductsService(repository Repository) *ProductsService {
	return &ProductsService{repository: repository}
}

type ProductsService struct {
	repository Repository
}

// GetAllProducts returns all the products in the repository
func (s ProductsService) GetAllProducts() (products []*domain.Product, err error) {
	products, err = s.repository.ReadAllData()
	return
}

// GetById returns a product by its id
func (s ProductsService) GetById(id int) (producto *domain.Product, err error) {
	producto, err = s.repository.GetById(id)
	if err != nil {
		return
	}
	return
}

// SearchProducts searches for products whith price >= priceS
func (s ProductsService) SearchProduct(priceS float64) (prs []domain.Product, err error) {
	prs, err = s.repository.SearchProduct(priceS)
	return
}

// Update updates a product by its id
func (s ProductsService) Update(id int, product *domain.Product) error {
	return s.repository.Update(id, product)
}

// Delete deletes a product by its id
func (s ProductsService) Delete(id int) (err error) {
	err = s.repository.Delete(id)
	if err != nil {
		if err != ErrProductNotFound {
			return
		}
		err = nil
	}
	return
}
