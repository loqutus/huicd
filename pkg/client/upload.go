package client

import (
	"net/http"
	"os"
	"fmt"
)

func Upload(fileName string) error{
	url := "http://" + host + ":" + port + "/upload/" + fileName
	fmt.Println("uploading: " + url)
	data, err := os.Open(fileName)
    if err != nil {
        return err
	}
	client := &http.Client{}
    req, err := http.NewRequest("POST", url, data)
    if err != nil {
        return err
	}
	req.Header.Add("Content-Type", "multipart/form-data")
	_, err = client.Do(req)
    if err != nil {
        return err
    }
    return nil
}