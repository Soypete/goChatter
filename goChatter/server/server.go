package server

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/davecgh/go-spew/spew"
)

type client chan<- string

var (
	incoming = make(chan client)
	outgoing = make(chan client)
	messages = make(chan string)
	port     string
)

// records list of active clients published messages
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-incoming:
			clients[cli] = true
		case cli := <-outgoing:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + "has arrived"
	messages <- who + " has arrived"
	incoming <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	messages <- who + " has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

//TODO: Add flags to start server and add client
//use script to set up server
//allow only one used
// use tags to add user verse setup server...
///

//add method for closing server
//make setup script do that only client operation run manully

// RunServer starts the chat server that the users will connect to
func RunServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	log.Println("listener running")
	spew.Dump(listener.Addr())

	go broadcaster()
	log.Println("broadcast running")
	for {
		conn, err := listener.Accept()
		spew.Dump(conn)
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
		log.Println("connection made")
	}
}
