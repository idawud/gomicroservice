package data

import (
	"encoding/json"
	"fmt"
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

func (p *Product) Error() string {
	panic("implement me")
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(p)
	return err
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(p)
	return err
}

func GetProducts() Products {
	return productList
}

func AddNewProduct(p *Product)  {
	p.ID = getNextID()
	p.CreatesOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	p.DeletedOn = time.Now().UTC().String()
	productList = append(productList, p)
}

func findProduct(id int)  (int, error){
	for index, p := range productList{
		if id == p.ID {
			return index, nil
		}
	}
	return -1, fmt.Errorf("Product With id %d Not Found", id)
}

func getNextID() int {
	lp := productList[len(productList) -1]
	return lp.ID + 1
}

func GetProduct(id int) (*Product, error)  {
	pos, err := findProduct(id)
	if err != nil{
		return nil, err
	}
	return productList[pos], nil
}

func UpdateProduct(id int, p *Product) (*Product, error)  {
	pos, err := findProduct(id)
	if err != nil{
		return nil, err
	}
	productList[pos].ID = id
	productList[pos].Name = p.Name
	productList[pos].Description = p.Description
	productList[pos].SKU = p.SKU
	productList[pos].Price = p.Price
	productList[pos].UpdatedOn = time.Now().UTC().String()

	return productList[pos], nil
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
