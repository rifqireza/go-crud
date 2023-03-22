package product

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/products", GetAllProduct).Methods(http.MethodGet)

	return r
}
