package svr

import (
	"bufio"
	"fmt"
	"log"
	"net"

	main "github.com/Soypete/goChatter/addUser"
)

type client chan<- string

var (
	incoming = make(chan client)
	outgoing = make(chan client)
	messages = make(chan string)
	port     string
	//Listener is what the clients connect to.
	Listener net.Listener
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

func handleConn(conn net.Conn, ch chan string) {
	go clientWriter(conn, ch)

	who := main.Users[conn]
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
		newmsg := fmt.Sprintln(msg)
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
		ch := make(chan string)
		go handleConn(conn, ch)
		log.Println("connection made")

	}
}
