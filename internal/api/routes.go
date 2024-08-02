package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"yougo-api/internal"
)

func NewRouter(api internal.YougoAPI) http.Handler {
	router := chi.NewRouter()
	router.Get("/", api.ListGos)
	router.Get("/{alias}", api.GetGo)
	return router
}
