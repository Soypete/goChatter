package main

//TODO:
//allow only one used
//tests
import (
	"flag"

	"github.com/Soypete/goChatter/server"
)

func main() {

	server := flag.Bool("server", false, "starts the server")

	client := flag.Bool("client", false, "adds a chat client to the server")

	name := flag.String("username", "name", "the username of the chat server client")

	flag.Parse()

	if *client {
		svr.Client(*name)
		return
	}
	if *server {
		svr.RunServer()
	}

}
