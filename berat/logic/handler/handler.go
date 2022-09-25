package handler

import "net/http"

type Handler interface {
	GetBerat(w http.ResponseWriter, r *http.Request)
	InsertBerat(w http.ResponseWriter, r *http.Request)
}
