package handlers

import (
	"github.com/idawud/gomicroservice/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger)  *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP( rw http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	p.l.Println(r.Method + " Not Allowed")
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProduct( rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("Products Get")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "Unable to marshal to json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Products Post")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil{
		http.Error(rw, "Unable to marshal to json", http.StatusBadRequest)
	}
	data.AddNewProduct(prod)
}
