package main

//TODO:
//allow only one used
//add method for closing server
//make setup script do that only client operation run manully
import (
	"flag"

	"github.com/Soypete/goChatter/server"
)

func main() {

	server := flag.Bool("server", false, "starts the server")

	user := flag.Bool("user", false, "adds a chat client to the server")

	flag.Parse()

	if *user {
		svr.Client()
		return
	}
	if *server {
		svr.RunServer()
	}

}
