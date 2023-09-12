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
	return &ProductsService{Repository: repository}
}

type ProductsService struct {
	Repository Repository
}

// GetAllProducts returns all the products in the repository
func (s ProductsService) GetAllProducts() (products []*domain.Product, err error) {
	products, err = s.Repository.ReadAllData()
	return
}

// GetById returns a product by its id
func (s ProductsService) GetById(id int) (producto *domain.Product, err error) {
	producto, err = s.Repository.GetById(id)
	if err != nil {
		return
	}
	return
}

// SearchProducts searches for products whith price >= priceS
func (s ProductsService) SearchProduct(priceS float64) (prs []domain.Product, err error) {
	prs, err = s.Repository.SearchProduct(priceS)
	return
}

// Update updates a product by its id
func (s ProductsService) Update(id int, product *domain.Product) error {
	return s.Repository.Update(id, product)
}

// Delete deletes a product by its id
func (s ProductsService) Delete(id int) (err error) {
	err = s.Repository.Delete(id)
	if err != nil {
		if err != ErrProductNotFound {
			return
		}
		err = nil
	}
	return
}
