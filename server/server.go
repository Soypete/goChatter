package svr

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
	spew.Fdump(conn, conn.LocalAddr().String())
	who := Users[conn.LocalAddr().String()]
	spew.Fdump(conn, who)

	go clientWriter(conn, ch, who)

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
		ch := make(chan string)
		go handleConn(conn, ch)
		log.Println("connection made")

	}
}
