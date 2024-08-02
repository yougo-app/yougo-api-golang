package internal

import (
	"yougo-api/api"
)

type yougoService struct {
	db YougoBackend
}

func NewYougoService(backend YougoBackend) YougoService {
	return &yougoService{db: backend}
}

func (svc yougoService) ListGos() ([]api.Go, error) {
	return svc.db.ListGos()
}

func (svc yougoService) GetGo(id string) (api.Go, error) {
	return svc.db.GetGo(id)
}
