package repository

import (
	"io"
	"os"
)

func LoadData(filePath string) (dataR []byte, err error) {
	data, err := os.Open(filePath)
	if err != nil {
		return
	}
	dataR, err = io.ReadAll(data)
	//	fmt.Println("LEI", string(i))

	return
}
