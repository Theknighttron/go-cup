package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, world!!")

		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "An error occured", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", data)

		// write back to the user
		fmt.Fprintf(w, "Hello %s\n", data)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye world!!")
	})

	http.ListenAndServe(":8080", nil)
}
