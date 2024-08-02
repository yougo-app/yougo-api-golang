package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"net/http"
	"yougo-api/internal"
	"yougo-api/internal/api"
	"yougo-api/internal/store"
)

func init() {
	functions.HTTP("gos", gosRootFunc)
}

func initialiseDefaultHandler() http.Handler {
	println("Initialising handler...")
	backend, _ := store.NewLocalBackend()
	service := internal.NewYougoService(backend)
	apiImpl := api.NewYougoAPI(service)
	return api.NewRouter(apiImpl)
}

// helloWorld writes "Hello, World!" to the HTTP response.
func gosRootFunc(w http.ResponseWriter, r *http.Request) {
	handler := initialiseDefaultHandler()
	handler.ServeHTTP(w, r)
}
