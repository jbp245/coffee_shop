package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jbp245/coffee_shop/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	json.Marshal(lp)
}
