package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
	// "golang.org/x/term"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hi %s\n", s.User()))
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
