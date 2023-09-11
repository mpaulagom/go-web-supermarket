package store

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/mpaulagom/go-web-supermarket/internal/domain"
	"github.com/mpaulagom/go-web-supermarket/internal/product"
)

type ProductStorage struct {
	FilePath string
}

/* 	APUNTE:
	Cuando usar un puntero o no en un slice' si voy a tener una estructura unica
 	y no necesito tener otros repositorios que tengan que referenciar directamente a la misma
	slice ahi si tendria que usar un puntero, pero si no no hace falta
*/

// NewProductStorage creates a new storage for products
func NewProductStorage(fpath string) *ProductStorage {
	return &ProductStorage{
		FilePath: fpath,
	}
}

// ReadAllData reads data from the json file and loads it in memory
func (rs *ProductStorage) ReadAllData() (products []*domain.Product, err error) {
	data, err := os.Open(rs.FilePath)
	if err != nil {
		err = errors.New("no encontre el fucking path" + rs.FilePath)
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

func (rs *ProductStorage) GetById(id int) (prod *domain.Product, err error) {
	products, err := rs.ReadAllData()
	if len(products) == 0 {
		err = product.ErrEmptySupermarket
		return
	}
	if err != nil {
		return
	}
	for _, pr := range products {
		if pr.Id == id {
			prod = pr
			return
		}
	}
	err = product.ErrProductNotFound
	return
}

// Update updates the product by its id
func (rp ProductStorage) Update(id int, product *domain.Product) (err error) {
	products, err := rp.ReadAllData()
	if err != nil {
		return
	}
	for index, pr := range products {
		if pr.Id == id {
			products[index] = product
		}
	}
	productsJson, err := json.Marshal(products)
	os.WriteFile("products.json", productsJson, 0644)
	return
}

// Delete deletes the product from the json file by its id
func (rp ProductStorage) Delete(id int) (err error) {
	products, err := rp.ReadAllData()
	if err != nil {
		return
	}
	var found bool
	for index, pr := range products {
		if pr.Id == id {
			products = removeElement(index, products)
			found = true
		}
	}
	if !found {
		err = product.ErrProductNotFound
		return
	}
	productsJson, err := json.Marshal(products)
	if err != nil {
		return
	}
	err = os.WriteFile(rp.FilePath, productsJson, os.ModePerm)
	return
}

// removeElement removes element by replacing it with the last element of the slice, asuming order isn't relevant
func removeElement(position int, slice []*domain.Product) (pr []*domain.Product) {
	slice[position] = slice[len(slice)-1]
	return slice
}

// SearchProduct searches for the products that have price equal or greater than param price
func (rp ProductStorage) SearchProduct(price float64) (prs []domain.Product, err error) {
	products, err := rp.ReadAllData()
	if err != nil {
		return
	}
	if len(products) == 0 {
		err = product.ErrEmptySupermarket
		return
	}

	for _, pr := range products {
		if pr.Price > price {
			prs = append(prs, *pr)
		}
	}
	return
}
