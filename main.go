package main

//TODO:
//use script to set up server
//allow only one used
//add method for closing server
//make setup script do that only client operation run manully
import (
	"flag"

	"github.com/SoyPete/goChatter/server"
)

func main() {

	server := flag.Bool("server", false, "starts the server")

	user := flag.Bool("user", false, "adds a chat client to the server")

	close := flag.Bool("close", false, "closes the server")

	flag.Parse()

	if *user {
		svr.Client()
		return
	}
	if *server {
		svr.RunServer()
	}
	if *close {
		svr.Listener.Close()
	}

}
