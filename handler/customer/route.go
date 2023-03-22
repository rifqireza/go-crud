package customer

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) *mux.Router {

	r.HandleFunc("/customers", Index).Methods(http.MethodGet)
	r.HandleFunc("/customer", Add).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id}", GetById).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", DeleteByID).Methods(http.MethodDelete)

	return r
}
