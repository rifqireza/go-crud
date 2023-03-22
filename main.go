package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jeypc/lat1/config"
	"github.com/jeypc/lat1/handler/auth"
	"github.com/jeypc/lat1/handler/customer"
	"github.com/jeypc/lat1/handler/product"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", auth.Login).Methods(http.MethodPost)
	r.HandleFunc("/logout", auth.Logout).Methods(http.MethodPost)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(config.AuthMiddleware)
	customer.Routes(api)
	product.Routes(api)

	handler := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(r)

	http.ListenAndServe(":8090", handler)
}
