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
	users    = 1
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

	var who string
	ch := make(chan string)
	go clientWriter(conn, ch)
	who = fmt.Sprintf("USER%d", users)
	ch <- "You are " + who
	messages <- who + " has arrived"
	users++
	incoming <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	messages <- who + " has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch <-chan string, name string) {
	for msg := range ch {
		newmsg := fmt.Sprintf("%s: %s\n", name, msg)
		fmt.Fprintln(conn, newmsg)
	}
}

// RunServer starts the chat server that the users will connect to
func RunServer() {
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

		go handleConn(conn)
		log.Println("connection made")

	}
}
