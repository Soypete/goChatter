package main

import (
	"flag"
	"fmt"
)

func main() {

	var port string

	server := flag.Bool("server", false, "starts the server")

	user := flag.Bool("user", false, "adds a chat client to the server")

	flag.StringVar(&port, "port", "8081", "the port that a chat user will be user")

	flag.Parse()

	fmt.Println("user:", *user)
	fmt.Println("svar:", port)
	if *user {
		server.Client(port)
		return
	}
	fmt.Println("server:", *server)
	if *server {
		server.RunServer()
	}

}
