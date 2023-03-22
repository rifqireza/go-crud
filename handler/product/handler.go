package product

import (
	"encoding/json"
	"net/http"

	"github.com/jeypc/lat1/models"
)

var productModel = models.NewProductModel()

func GetAllProduct(resp http.ResponseWriter, req *http.Request) {
	products, err := productModel.FindAll()
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"product": products,
	}

	json.NewEncoder(resp).Encode(data)
}
