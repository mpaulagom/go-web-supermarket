package product

import (
	"encoding/json"
	"io"
	"os"

	"github.com/mpaulagom/go-web-supermarket/internal/domain"
)

type RepositoryProductInMemory struct {
	Products []*domain.Product
	FilePath string
}

// Cuando usar un puntero o no en un slice' si voy a tener una estructura unica
// y no necesito tener otros repositorios que tengan que referenciar directamente a la misma
// slice ahi si tendria que usar un puntero, pero si no no hace falta
// NewRepositoryProductInMemory creates a new storage for products
func NewRepositoryProductInMemory(prods []*domain.Product, fpath string) *RepositoryProductInMemory {
	return &RepositoryProductInMemory{
		Products: prods,
		FilePath: fpath,
	}
}

func (rs RepositoryProductInMemory) ReadAllData() (products []*domain.Product, err error) {

	/* if rs.Products ==  {
		err =  errors.New("there are no products")
	} */
	products = rs.Products
	return
}

func (rs RepositoryProductInMemory) ReadAllData2() (products []*domain.Product, err error) {
	data, err := os.Open(rs.FilePath)
	if err != nil {
		return
	}
	jsonProducts, err := io.ReadAll(data)
	//jsonProducts, err := rs.LoadData()
	if err != nil {
		return
	}
	json.Unmarshal(jsonProducts, &products)
	return
}

func (rs RepositoryProductInMemory) GetById(id int) (product *domain.Product, err error) {
	products, err := rs.ReadAllData()
	if len(products) == 0 {
		err = ErrEmptySupermarket
		return
	}
	if err != nil {
		return
	}
	for _, pr := range products {
		if pr.Id == id {
			product = pr
			return
		}
	}
	err = ErrProductNotFound
	return
}

func (rs RepositoryProductInMemory) Delete(id int) error {
	return nil
}

func (rs RepositoryProductInMemory) SearchProduct(price float64) (prs []domain.Product, err error) {
	return nil, nil
}

func (rs RepositoryProductInMemory) Update(id int, product *domain.Product) error {
	return nil
}
