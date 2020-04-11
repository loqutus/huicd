package server

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"os"
	"io/ioutil"
)

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST"{
		vars := mux.Vars(req)
		fileName := vars["filename"]
		log.Println("POST /upload/"+fileName)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil{
			log.Println("ERROR: ioutil.ReadAll error")
			log.Println(err.Error())
			fmt.Fprintf(w, "upload error: " + err.Error())
			return
		}
		out, err := os.Create(dataDir + "/" + fileName)
        if err != nil {
			log.Println("ERROR: failed to open the file for writing", err.Error())
            fmt.Fprintf(w, "Failed to open the file for writing: " + err.Error())
            return
        }
		defer out.Close()
		_, err = out.Write(body)
		if err != nil{
			log.Println("ERROR: io.Write has failed", err.Error())
			fmt.Println(w, "io.Write has failed: " + err.Error())
			return
		}
	} else {
		log.Println("ERROR: unknown method for /upload: " + req.Method)
		fmt.Fprintf(w, "unknown method: " + req.Method)
	}
}