package store

import (
	"golang.org/x/exp/maps"
	"yougo-api/api"
	"yougo-api/internal"
)

type localBackend struct {
	converter GoConverter
	gos       map[string]dbGo
}

func NewLocalBackend() (internal.Backend, error) {
	gos, err := loadData()

	if err != nil {
		return nil, err
	}

	return localBackend{gos: gos, converter: &GoConverterImpl{}}, nil
}

func (b localBackend) ListGos() ([]api.Go, error) {
	return b.converter.GosFromDB(maps.Values(b.gos)), nil
}

func (b localBackend) GetGo(id string) (api.Go, error) {
	return b.converter.GoFromDB(b.gos[id]), nil
}

func loadData() (map[string]dbGo, error) {
	gos := make(map[string]dbGo)
	gos["test"] = dbGo{
		Id:    "1",
		Alias: "test",
		Href:  "https://test.example.com",
	}
	gos["google"] = dbGo{
		Id:    "1",
		Alias: "test",
		Href:  "https://google.com",
	}
	return gos, nil
}
