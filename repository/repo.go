package repository

import (
	"encoding/json"
	"io"
	"os"
)

type Repository interface {
	LoadSuperMarket() (products []*Product, err error)
}
type RepositoryJson struct {
	FilePath string
}

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code        string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func NewRepositoryJson(fpath string) *RepositoryJson {
	return &RepositoryJson{FilePath: fpath}
}

func (rs *RepositoryJson) LoadSuperMarket() (products []*Product, err error) {
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
