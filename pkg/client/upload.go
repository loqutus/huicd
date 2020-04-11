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
	defer data.Close()
	client := &http.Client{}
    req, err := http.NewRequest("POST", url, data)
    if err != nil {
        return err
	}
	_, err = client.Do(req)
    if err != nil {
        return err
    }
    return nil
}