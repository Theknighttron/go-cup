package handler

import (
    "fmt"
    "log"
    "io"
    "net/http"
)


type Hello struct {
    l *log.Logger

}

func NewHello(l *log.Logger) *Hello {
    return &Hello{l}
}


// define the method to implement hello struct
func (h*Hello) ServerHTTP(w http.ResponseWriter, r http.Request) {
		h.l.Println("Hello, world!!")

		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "An error occured", http.StatusBadRequest)
			return
		}
		log.Printf("Data %s\n", data)

		// write back to the user
		fmt.Fprintf(w, "Hello %s\n", data)
}
