# goChatter
chat server written in go


# Usage
To start the server use either

`go build main.go`
`./main -server`

or in a single command

`go run main.go -server`

This will allow you to see user's logging in about out and records of activity. To suppress to the background use

`go build main.go`
`./main -server`

To add user use the command

`./main -user`

if your use `go build` or if you built it in a single command use

`go run main.go -client`

The server will notify you that a connection has been made and allow you to send and receive messages.

To close the server cancel the run.

## Explanation
This chat server what build following an exercise in _The Go Programming Language_ by: Alan A.a. Donovan and Brian W. Kernighan 
[github repo](https://github.com/adonovan/gopl.io/tree/master/ch8)

## Other Resource Used
http://choly.ca/post/go-experiments-with-handler/
https://gobyexample.com/command-line-flags
https://gist.github.com/enricofoltran/10b4a980cd07cb02836f70a4ab3e72d7
