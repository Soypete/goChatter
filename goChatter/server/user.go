package svr

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Client creates a client that can chat on server
func Client() {
	url := fmt.Sprintf("localhost:8080")

	conn, err := net.Dial("tcp", url)
	spew.Dump(conn)
	spew.Dump(err)
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
