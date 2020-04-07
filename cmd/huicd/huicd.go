package main
import (
	"os"
	"github.com/loqutus/huicd/pkg/server"
	"github.com/loqutus/huicd/pkg/client"
)

	
func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0{
		panic("specify mode: server or client")
	}
	mode := argsWithoutProg[0]
	if mode == "server"{
		err := server.Serve()
		if err != nil{
			panic(err.Error())
		}
	} else if mode == "client"{
		err := client.Client()
		if err != nil{
			panic(err.Error())
		}
	} else {
		panic("unknown mode: " + mode)
	}
}