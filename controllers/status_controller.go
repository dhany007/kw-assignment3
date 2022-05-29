package controllers

import (
	"net/http"
)

type StatusController interface {
	GetStatus(w http.ResponseWriter, r *http.Request)
}
