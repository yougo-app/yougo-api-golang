package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"yougo-api/internal"
)

type yougoAPIImpl struct {
	service internal.YougoService
}

func NewYougoAPI(service internal.YougoService) internal.YougoAPI {
	return &yougoAPIImpl{service: service}
}

func (h yougoAPIImpl) ListGos(w http.ResponseWriter, r *http.Request) {
	gos, err := h.service.ListGos()
	if err != nil {
		handleError(w, err)
		return
	}
	jsonResponse(gos, w)
}

func (h yougoAPIImpl) GetGo(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")
	g, err := h.service.GetGo(alias)
	if err != nil {
		handleError(w, err)
		return
	}
	jsonResponse(g, w)
}

func jsonResponse(v any, w http.ResponseWriter) {
	jsonOutput, err := json.Marshal(v)
	if err != nil {
		handleError(w, err)
		return
	}
	_, err = w.Write(jsonOutput)
	if err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
