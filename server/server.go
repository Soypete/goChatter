package svr

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	incoming = make(chan client)
	outgoing = make(chan client)
	messages = make(chan string)
	port     string
	//Listener is what the clients connect to.
	Listener net.Listener

	users = make(map[net.Conn]string)
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

	who := users[conn]
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
		fmt.Println(users[conn], msg)
	}
}

// RunServer starts the chat server that the users will connect to
func RunServer(name string) {
	Listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer Listener.Close()

	log.Println("listener running")

	go broadcaster()
	log.Println("broadcast running")
	for {
		conn, err := Listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		AddUser(name, conn)
		go handleConn(conn)
		log.Println("connection made")

	}
}

// AddUser adds users to this session of the
func AddUser(name string, conn net.Conn) {
	users[conn] = name
}
