package client

import (
	"errors"
	"os"
)

var host = "127.0.0.1"
var port = "6666"

func Client(args []string) error{
	hostEnv := os.Getenv("HOST")
	if hostEnv != ""{
		host = hostEnv
	}
	portEnv := os.Getenv("PORT")
	if portEnv != ""{
		port = portEnv
	}
	if len(args) == 0{
		return errors.New("missing client action: upload or download")
	}
	action := args[0]
	if action == "upload"{
		Upload(args[1])
	} else {
		return errors.New("unknown client action: " + args[0])
	}
	return nil
}