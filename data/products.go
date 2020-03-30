package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID int 				`json:"id"`
	Name string 		`json:"name"`
	Description string 	`json:"description"`
	Price float32 		`json:"price"`
	SKU string 			`json:"sku"`
	CreatesOn string 	`json:"-"`
	UpdatedOn string 	`json:"-"`
	DeletedOn string 	`json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(p)
	return err
}
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "coffee with milk",
		Price:       2.96,
		SKU:         "abc459",
		CreatesOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espressos",
		Description: "coffee withot milk",
		Price:       1.99,
		SKU:         "fgj768",
		CreatesOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
