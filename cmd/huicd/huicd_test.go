package main

import (
	"testing"
	"os"
	"github.com/loqutus/huicd/pkg/server"
	"github.com/loqutus/huicd/pkg/client"
	"time"
	"fmt"
)

func TestMain(t *testing.T){
	err := os.Chdir("../../build")
	if err != nil{
		t.Errorf(err.Error())
	}
	fmt.Println(path)
	if err := os.Setenv("DATA_DIR", "/tmp"); err != nil{
		t.Errorf(err.Error())
	}
	go server.Serve()
	time.Sleep(2*time.Second)
	if err := os.Remove("/tmp/test"); err != nil{
		t.Errorf(err.Error())
	}
	if err := client.Upload("test"); err != nil{
		t.Errorf(err.Error())
	}
	if err := os.Remove("./test"); err != nil{
		t.Errorf(err.Error())
	}
	time.Sleep(1*time.Second)
	if err := client.Download("test"); err != nil{
		t.Errorf(err.Error())
	}
}