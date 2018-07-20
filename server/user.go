package svr

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/davecgh/go-spew/spew"
)

var (
	//Users list of usernames who have logged in to system
	Users = make(map[string]string)
)

// Client creates a client that can chat on server
func Client(name string) {
	url := fmt.Sprintf("localhost:8080")
	fmt.Println(name)
	fmt.Println(Users)

	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	if len(Users) != 0 {
		for this, ppl := range Users {
			fmt.Println(conn, name)
			spew.Dump(conn, ppl)
			spew.Dump(conn, this)
			if name == ppl {
				panic(errors.New("User already exists, try again"))
			}
		}
	}
	Users[conn.LocalAddr().String()] = name
	fmt.Println(Users)

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
