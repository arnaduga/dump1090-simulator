package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = "8080"

var s = 0

func dataJson(w http.ResponseWriter, r *http.Request) {
	f := fmt.Sprintf("steps/step-%v.json", ((s % 29) + 1))
	fmt.Printf("Endpoint called: /data.json, to be served with file %v\n", f)

	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-source", f)
	w.Write(data)
	s += 1
}

func handleRequests() {
	http.HandleFunc("/data.json", dataJson)

	sp := os.Getenv("DUMP1090_SIMU_PORT")

	if sp == "" {
		fmt.Printf("NO port defintion found in environment variable. Will use the default: %v\n", PORT)
		sp = PORT
	} else {
		fmt.Printf("Port definition found in environment variable. Will use it: %v\n", sp)

	}

	fmt.Printf("Starting server on port %v\n", sp)
	fmt.Printf("NOTE: only ONE endpoint is valid: /data.json\n")
	fmt.Printf("TIP : curl -s http://localhost:%v/data.json\n", sp)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", sp), nil))
}

func main() {
	handleRequests()
}
