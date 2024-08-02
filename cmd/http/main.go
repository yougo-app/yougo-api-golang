package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func printOutput(output any) {
	jsonOutput, err := json.Marshal(output)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(jsonOutput))

}

func main() {

	gos, _ := service.ListGos()
	testGo, _ := service.GetGo("test")
	testMissingGo, _ := service.GetGo("missing")

	printOutput(gos)
	printOutput(testGo)
	printOutput(testMissingGo)
}
