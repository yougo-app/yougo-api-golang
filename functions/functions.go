package functions

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"net/http"
	"sync"
	"yougo-api/internal"
	"yougo-api/internal/api"
	"yougo-api/internal/store"
)

// create & access singleton instance of the default http router
var lock = &sync.Mutex{}
var defaultHttpHandler *http.Handler

func getDefaultHttpHandler() *http.Handler {
	if defaultHttpHandler == nil {
		lock.Lock()
		defer lock.Unlock()
		if defaultHttpHandler == nil {
			defaultHttpHandler = initialiseDefaultHandler()
		}
	}
	return defaultHttpHandler
}

func init() {
	// eagerly initialise the default handler singleton
	getDefaultHttpHandler()
	functions.HTTP("gos", gosRootFunc)
}

func initialiseDefaultHandler() *http.Handler {
	println("Initialising handler...")
	backend, _ := store.NewLocalBackend()
	service := internal.NewYougoService(backend)
	apiImpl := api.NewYougoAPI(service)
	router := api.NewRouter(apiImpl)
	return &router
}

// helloWorld writes "Hello, World!" to the HTTP response.
func gosRootFunc(w http.ResponseWriter, r *http.Request) {
	handler := *getDefaultHttpHandler()
	handler.ServeHTTP(w, r)
}
