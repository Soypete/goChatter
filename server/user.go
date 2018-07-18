package svr

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	main "github.com/Soypete/goChatter/addUser"
)

// Client creates a client that can chat on server
func Client(name string) {
	url := fmt.Sprintf("localhost:8080")

	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	err = main.AddUser(name, conn)
	if err != nil {
		panic(err)
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
