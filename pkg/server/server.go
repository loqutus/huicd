package server

import (
	"net/http"
	"os"
	"log"
)

func Serve() error{
	var port = "6666"
	portEnv := os.Getenv("PORT")
	if portEnv != ""{
		port = portEnv
	}
	var dataDir = "/data"
	dataDirEnv := os.Getenv("DATA_DIR")
	if dataDirEnv != ""{
		dataDir = dataDirEnv
	}
	log.Println("DataDir:", dataDir)
	http.HandleFunc("/version", version)
	log.Println("Starting HTTP server on port", port)
	http.ListenAndServe(":"+port, nil)
	return nil
}