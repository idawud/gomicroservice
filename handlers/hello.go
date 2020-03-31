package handlers

import (
	"fmt"
	"github.com/idawud/gomicroservice/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request)  {
	// PUT update resource
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 && len(g[0]) != 2  {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(g[0][1])
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		h.l.Println("Got Id, " , id)
		prod := &data.Product{}
		errJson := prod.FromJSON(r.Body)
		if errJson != nil{
			http.Error(rw, "Unable to marshal to json", http.StatusBadRequest)
		}
		product, errUpdate := data.UpdateProduct(id, prod)
		if errUpdate != nil {
			http.Error(rw, "Unable to update product", http.StatusInternalServerError)
		}
		_, _ = fmt.Fprintln(rw, product)
	}
}
