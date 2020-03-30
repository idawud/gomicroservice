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
		p.GetProducts(rw, r)
		return
	}
	p.l.Println(r.Method + " Not Allowed")
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts( rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("Products Get")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "Unable to marshal to json", http.StatusInternalServerError)
	}
}
