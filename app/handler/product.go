package handler

import (
	"boilerplate/app/model"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Product struct {
	DB *gorm.DB
}

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	p.DB = model.DBMigrate()
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := p.DB.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, product)
}

func GetProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.Product{}
	db.Find(&products)
	respondJSON(w, http.StatusOK, products)
}
