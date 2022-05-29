package controllers

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

type StatusControllerImpl struct {
}

func NewStatusController() StatusController {
	return &StatusControllerImpl{}
}

func (controller *StatusControllerImpl) GetStatus(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusBadRequest)
	}

	t := template.Must(template.ParseFiles("template.html"))

	randomValueWater := rand.Intn(100)
	randomValueWind := rand.Intn(100)

	response := map[string]interface{}{
		"status": map[string]interface{}{
			"water": randomValueWater,
			"wind":  randomValueWind,
		},
	}

	fmt.Println(response)
	t.ExecuteTemplate(w, "template.html", response)
}
