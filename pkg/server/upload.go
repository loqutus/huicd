package server

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"bytes"
	"fmt"
	"os"
	"io"
)

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST"{
		vars := mux.Vars(req)
		fileName := vars["filename"]
		log.Println("POST /upload/"+fileName)
		var buf bytes.Buffer
		reader, err := req.MultipartReader()
		if err != nil{
			log.Println("ERROR: MultipartReader error")
			log.Println(err.Error())
			fmt.Fprintf(w, "upload error: " + err.Error())
		}
		out, err := os.Create(dataDir + "/" + fileName)
        if err != nil {
			log.Println("ERROR: failed to open the file for writing")
            fmt.Fprintf(w, "Failed to open the file for writing")
            return
        }
		defer out.Close()
		_, err = io.Copy(out, reader)
        if err != nil {
            fmt.Fprintln(w, "copy error: " + err.Error())
        }
	} else {
		log.Println("ERROR: unknown method for /upload: " + req.Method)
		fmt.Fprintf(w, "unknown method: " + req.Method)
	}
}