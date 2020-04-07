package server

import (
	"net/http"
	"fmt"
	"log"
)

func version(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET"{
		log.Println("GET /version")
		fmt.Fprintf(w, "0.0.1")
	} else {
		log.Println("ERROR: unknown method for /version: " + req.Method)
		fmt.Fprintf(w, "unknown method: " + req.Method)
	}
}