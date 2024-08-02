package main

import (
	"encoding/json"
	"fmt"
	"os"
	"yougo-api/internal"
	"yougo-api/internal/store"
)

func printOutput(output any) {
	jsonOutput, err := json.Marshal(output)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(jsonOutput))

}

func main() {
	backend, err := store.NewLocalBackend()

	if err != nil {
		os.Exit(1)
	}

	service := internal.NewYougoService(backend)

	gos, _ := service.ListGos()
	testGo, _ := service.GetGo("test")
	testMissingGo, _ := service.GetGo("missing")

	printOutput(gos)
	printOutput(testGo)
	printOutput(testMissingGo)
}
