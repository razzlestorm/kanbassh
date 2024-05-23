package entities

import (
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type Board struct {
	Name    string
	Columns []Column
	Users   []User
}

type Column struct {
	Name    string
	Tickets []Ticket
}

type Ticket struct {
	Name        string
	Due_Date    time.Time
	Description string
	Assignee    User
}

type User struct {
	Session  ssh.Session
	Terminal *term.Terminal
}
