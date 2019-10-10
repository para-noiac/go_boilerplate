package handler

import (
	"encoding/json"
	"fmt"
	"go_boilerplate/app/model"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := h.DB.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, product)
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []model.Product{}
	h.DB.Find(&products)
	respondJSON(w, http.StatusOK, products)
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST")
}
