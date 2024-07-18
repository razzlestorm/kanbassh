package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gliderlabs/ssh"
	"github.com/razzlestorm/kanbassh/cli"
	"github.com/razzlestorm/kanbassh/internal/entities"
	"golang.org/x/term"
)

func evaluate(input string, comms *cli.CommandHandler){
	args := strings.Split(input, " ")
	
	command, optional := args[0], args[1:]

	output, ok := comms.Commands[command]
	
	if ok {
		output(optional)		
	} else if input == "" {


	} else {
		fmt.Printf("%v: command not found\n", command)
	}
}

func main() {

	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hi %s\n", s.User()))
		term := term.NewTerminal(s, fmt.Sprintf("%s > ", s.User()))
		new_user := entities.User{Session: s, Terminal: term}
		commandlist := cli.NewCommandHandler(new_user)
		for {
			line, err := term.ReadLine()
			if err != nil {
				break
			}
			evaluate(line, commandlist)
		}
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))

}
