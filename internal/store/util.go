package store

import "yougo-api/api"

// goverter:converter
// goverter:output:file ./generated.go
// goverter:output:package yougo-api/internal/store
type GoConverter interface {
	GosToDB(source []api.Go) []dbGo
	GoToDB(source api.Go) dbGo
	GosFromDB(source []dbGo) []api.Go
	GoFromDB(source dbGo) api.Go
}
