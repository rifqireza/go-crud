package customer

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/lat1/entity"
	"github.com/jeypc/lat1/models"
)

var customerModel = models.NewCustomerModel()

func Index(resp http.ResponseWriter, req *http.Request) {
	customers, _ := customerModel.FindAll()

	data := map[string]interface{}{
		"customer": customers,
	}

	json.NewEncoder(resp).Encode(data)
}

func GetById(resp http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	customer, err := customerModel.FindById(id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(resp).Encode(customer)
}

func Add(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var customer entity.Customer

		body, err := io.ReadAll(req.Body)
		err = json.Unmarshal(body, &customer)

		if err != nil {
			json.NewEncoder(resp).Encode(err)
		}

		customerModel.Create(customer)
		json.NewEncoder(resp).Encode(body)
	}
}

func DeleteByID(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		id := mux.Vars(req)["id"]
		message, err := customerModel.DeleteByID(id)
		if err != nil {
			json.NewEncoder(resp).Encode(err)
		}

		json.NewEncoder(resp).Encode(message)
	}
}
