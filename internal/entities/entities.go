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

func NewColumn(name string, tickets []string, users []User) *Column {
	if len(name) < 1 {
		panic("A new Column must have a name")
	}
	c := &Column{
		Name: name,
		Tickets: []Ticket{},
	}
	for _, ticket := range tickets {
		c.Tickets = append(c.Tickets, Ticket{Name: ticket})
	}
	return c

}

type Ticket struct {
	Name        string
	Due_Date    time.Time
	Description string
	Assignee    User
}


func NewTicket(name string, date time.Time, description string, assignee User) *Ticket {
	if len(name) < 1 {
		panic("A new Ticket must have a name")
	}
	t := &Ticket{
		Name: name,
		Due_Date: date,
		Description: description,
		Assignee: assignee,
	}
	return t

}

type User struct {
	Session  ssh.Session
	Terminal *term.Terminal
}
