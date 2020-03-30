package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Main endpoint, Hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil{
			/* //same as below
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Ooops Error"))
			 */
			http.Error(rw, "Oops Error", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye Corona")
	})

	fmt.Println("Server running on http://localhost:8080/" )
	log.Println(" Server started at ", time.Now().String())
	http.ListenAndServe(":8080", nil)
}
