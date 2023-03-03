package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = "8080"

var currentStep = 0

func rootResource(w http.ResponseWriter, r *http.Request) {
	fileToOpen := fmt.Sprintf("steps/step-%v.json", ((currentStep % 29) + 1))
	fmt.Printf("Endpoint called: /data.json, to be served with file %v\n", fileToOpen)

	data, err := os.ReadFile(fileToOpen)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-source", fileToOpen)
	w.Write(data)
	currentStep = currentStep + 1
}

func handleRequests() {
	http.HandleFunc("/data.json", rootResource)

	serverPort := os.Getenv("DUMP1090_SIMU_PORT")

	if serverPort == "" {
		fmt.Printf("NO port defintion found in environment variable. Will use the default: %v\n", PORT)
		serverPort = PORT
	} else {
		fmt.Printf("Port definition found in environment variable. Will use it: %v\n", serverPort)

	}

	fmt.Printf("Starting server on port %v\n", serverPort)
	fmt.Printf("NOTE: only ONE endpoint is valid: /data.json\n")
	fmt.Printf("TIP : curl -s http://localhost:%v/data.json\n", serverPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", serverPort), nil))
}

func main() {
	handleRequests()
}
