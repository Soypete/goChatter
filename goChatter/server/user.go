package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Client creates a client that can chat on server
func Client(port string) {
	url := fmt.Sprintf("localhost:%s", port)

	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	if _, err := io.Copy(conn, os.Stdin); err != nil {
		log.Fatal(err)
	}
	conn.Close()
	<-done
}
