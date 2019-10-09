package handler

import (
	"fmt"
	"go_boilerplate/app/model"
	"net/http"
)

// func CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	p.DB = model.DBMigrate()
// 	product := model.Product{}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&product); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := p.DB.Save(&product).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusCreated, product)
// }

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var temp = model.DBMigrate()
	products := []model.Product{}
	temp.Find(&products)
	respondJSON(w, http.StatusOK, products)
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST")
}
