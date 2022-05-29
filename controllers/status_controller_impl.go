package controllers

import (
	"assignment3/models"
	"encoding/json"
	"math/rand"
	"net/http"
)

type StatusControllerImpl struct {
}

func NewStatusController() StatusController {
	return &StatusControllerImpl{}
}

func (controller *StatusControllerImpl) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	randomValueWater := rand.Intn(100)
	randomValueWind := rand.Intn(100)

	tempStatus := models.Status{
		Water: randomValueWater,
		Wind:  randomValueWind,
	}

	response := map[string]interface{}{
		"status": tempStatus,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
