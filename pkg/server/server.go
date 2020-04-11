package server

import (
	"github.com/gorilla/mux"
	"os"
	"log"
	"net/http"
)

var dataDir = "/data"
var port = "6969"

func Serve() error{
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
	r.HandleFunc("/download/{filename}", downloadHandler)
	log.Println("Starting HTTP server on port", port)
	http.ListenAndServe(":"+port, r)
	return nil
}