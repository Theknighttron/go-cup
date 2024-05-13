package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
    l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
    return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
   h.l.Println("Hello world!!")

   data, err := io.ReadAll(r.Body)
   if err != nil {
       http.Error(rw, "An error occured", http.StatusBadRequest)
       return
   }

   fmt.Fprintf(rw, "Hello %s", data)
}
