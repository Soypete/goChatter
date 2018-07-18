package adduser

import (
	"errors"
	"net"
)

var (
	//Users list of usernames who have logged in to system
	Users = make(map[net.Conn]string)
)

// AddUser adds users to this session of the
func AddUser(name string, conn net.Conn) error {
	for _, ppl := range Users {
		if name == ppl {
			return errors.New("User already exists, try again")
		}
	}
	Users[conn] = name
	return nil
}
