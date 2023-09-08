package product

import (
	"encoding/json"
	"io"
	"os"

	"github.com/mpaulagom/go-web-supermarket/internal/domain"
)

type Repository interface {
	LoadSuperMarket() (products []*domain.Product, err error)
}
type RepositoryJson struct {
	FilePath string
}

// Cuando usar un puntero o no en un slice' si voy a tener una estructura unica
// y no necesito tener otros repositorios que tengan que referenciar directamente a la misma
// slice ahi si tendria que usar un puntero, pero si no no hace falta
func NewRepositoryJson(fpath string) *RepositoryJson {
	return &RepositoryJson{FilePath: fpath}
}

func (rs *RepositoryJson) LoadSuperMarket() (products []*domain.Product, err error) {
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
