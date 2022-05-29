package controllers

import (
	"net/http"
)

type StatusController interface {
	UpdateStatus(w http.ResponseWriter, r *http.Request)
}
