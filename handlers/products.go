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

    if r.Method == http.MethodPost {
        p.addProducts(rw, r)
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


func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
    p.l.Println("Handle POST products")


    prod := &data.Product{}
    err := prod.FromJSON(r.Body)
    if err != nil {
        http.Error(rw, "Unable to unmarshar json", http.StatusBadRequest)
    }

    p.l.Printf("Product: %#v", prod)

}
