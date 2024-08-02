package internal

import (
	"net/http"
	"yougo-api/api"
)

type YougoBackend interface {
	ListGos() ([]api.Go, error)
	GetGo(id string) (api.Go, error)
}

type YougoService interface {
	ListGos() ([]api.Go, error)
	GetGo(id string) (api.Go, error)
}

type YougoAPI interface {
	ListGos(w http.ResponseWriter, r *http.Request)
	GetGo(w http.ResponseWriter, r *http.Request)
}
