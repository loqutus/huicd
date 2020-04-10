package server

import (
	"github.com/gorilla/mux"
	"os"
	"log"
	"net/http"
)

var dataDir = "/data"

func Serve() error{
	var port = "6666"
	portEnv := os.Getenv("PORT")
	if portEnv != ""{
		port = portEnv
	}
	dataDirEnv := os.Getenv("DATA_DIR")
	if dataDirEnv != ""{
		dataDir = dataDirEnv
	}
	log.Println("DataDir:", dataDir)
	r := mux.NewRouter()
	r.HandleFunc("/version", versionHandler)
	r.HandleFunc("/upload/{filename}", uploadHandler)
	log.Println("Starting HTTP server on port", port)
	http.ListenAndServe(":"+port, r)
	return nil
}