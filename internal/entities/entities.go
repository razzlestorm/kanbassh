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

func NewBoard(name string, columns []string, users []User) *Board {
	if len(name) < 1 || len(columns) < 1 {
		panic("A new Board must have a name and at least one column")
	}
	b := &Board{
		Name: name,
		Columns: []Column{},
	}
	for _, col := range columns {
		b.Columns = append(b.Columns, Column{Name: col})
	}
	return b

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
