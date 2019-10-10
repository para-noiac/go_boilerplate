package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB *gorm.DB
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logerr(w.Write([]byte(err.Error())))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	logerr(w.Write([]byte(response)))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func logerr(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}
