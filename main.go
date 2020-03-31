package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/idawud/gomicroservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	pr := handlers.NewProducts(l)

	 sm := mux.NewRouter()
	 getRouter := sm.Methods(http.MethodGet).Subrouter()
	 getRouter.HandleFunc("/", pr.GetProduct)

	 putRouter := sm.Methods(http.MethodPut).Subrouter()
	 putRouter.HandleFunc(`/{id:[0-9]+}`, pr.UpdateProduct)
	 putRouter.Use(pr.MiddlewareProductValidation)

	 postRouter := sm.Methods(http.MethodPost).Subrouter()
	 postRouter.HandleFunc(`/`, pr.AddProduct)
	 postRouter.Use(pr.MiddlewareProductValidation)

	 server := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout:120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout:1*time.Second,
	 }

	 fmt.Println("Server running on http://localhost:8080/" )
	 log.Println(" Server started at ", time.Now().String())
	 go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	 }()

	 // Graceful shutdown
	 sigChan := make(chan os.Signal)
	 signal.Notify(sigChan, os.Interrupt)
	 signal.Notify(sigChan, os.Kill)

	 sig := <-sigChan
	 l.Println("Graceful shutdown", sig)

	 ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	 _ = server.Shutdown(ctx)
}
