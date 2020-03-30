package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger)  *GoodBye{
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP( rw http.ResponseWriter, r *http.Request)  {
	g.l.Println("goodbye ")
	fmt.Fprintf(rw, "goodbye corona")
}