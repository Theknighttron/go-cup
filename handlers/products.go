package handlers

import (
	"log"
	"net/http"

	"github.com/polyhistor2050/microservices/data"
)

type Products struct {
    l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
    return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        p.getProducts(rw, r)
        return
    }

    rw.WriteHeader(http.StatusMethodNotAllowed)
}


func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
    // Return a list of all products
    lp := data.GetProducts()
    err := lp.ToJSON(rw)

    if err != nil {
        http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
    }
}
