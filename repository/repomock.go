package repository

type MockRepository struct {
	Products []*Product
}

// No estoy segura de que tan bien esta esto
func (ms *MockRepository) LoadSuperMarket() (products []*Product, err error) {
	products = ms.Products
	return
}
