package server

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
)

func downloadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET"{
		vars := mux.Vars(req)
		fileName := vars["filename"]
		log.Println("GET /download/"+fileName)
		data, err := ioutil.ReadFile(dataDir + "/" + fileName)
    	if err != nil {
        	log.Println("ERROR: ioutil.ReadFile error")
			log.Println(err.Error())
			fmt.Fprintf(w, "download error: " + err.Error())
		}
		fmt.Fprintf(w, string(data))
	} else {
		log.Println("ERROR: unknown method for /download: " + req.Method)
		fmt.Fprintf(w, "unknown method: " + req.Method)
	}
}