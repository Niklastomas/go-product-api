package handlers

import (
	"log"
	"net/http"

	"github.com/niklastomas/go-product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// GET
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// Catch all
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
