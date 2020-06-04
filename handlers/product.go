package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/idawud/gomicroservice/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger)  *Products {
	return &Products{l}
}


func (p *Products) GetProducts( rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("Products Get")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "Unable to marshal to json", http.StatusInternalServerError)
	}
}

func (p *Products) GetProduct( rw http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	p.l.Println("Product with id ", vars["id"])
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	product, err := data.GetProduct(id)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}else {
		rw.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(rw).Encode(product)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Products Post")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddNewProduct(&prod)
}

func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p.l.Println("Products with id ", vars["id"])
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	_, err = data.UpdateProduct(id, &prod)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
}

type KeyProduct struct {}

func (p Products) MiddlewareProductValidation( next http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to process json ::" + err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}