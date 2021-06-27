package handler

import (
	"net/http"

	"github.com/MAAARKIN/unico/api/helper"
	"github.com/go-chi/chi/v5"
)

type Health struct{}

func (h Health) Route(r chi.Router) {
	r.Get("/", h.check)
}

func (m Health) check(w http.ResponseWriter, r *http.Request) {
	status := "UP"

	resp := ManagerHealth{
		Status: status,
	}
	statusCode := http.StatusOK
	helper.JsonResponse(w, resp, statusCode)
}

type ManagerHealth struct {
	Status string `json:"status"`
}
