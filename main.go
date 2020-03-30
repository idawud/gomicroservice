package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/idawud/gomicroservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	hh := handlers.NewHello(l)
	gb := handlers.NewGoodBye(l)

	 sm := http.NewServeMux()
	 sm.Handle("/", hh)
	 sm.Handle("/bye", gb)

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
