package client

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func Download(fileName string) error{
	url := "http://" + host + ":" + port + "/download/" + fileName
	fmt.Println("downloading: " + url)
	resp, err := http.Get(url)
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(fileName)
	if err != nil{
		return err
	}
	defer out.Close()
	io.Copy(out, resp.Body)
	return nil
}