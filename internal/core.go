package internal

import "yougo-api/api"

type Backend interface {
	ListGos() ([]api.Go, error)
	GetGo(id string) (api.Go, error)
}

type Service interface {
	ListGos() ([]api.Go, error)
	GetGo(id string) (api.Go, error)
}
