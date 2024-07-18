package cli

import (
	"fmt"
	"regexp"
	"os"
	"strings"
	"path/filepath"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
	"github.com/razzlestorm/kanbassh/internal/entities"
)

	
type CommandHandler struct {
	Commands map[string]func([]string)
	User entities.User
} 

func NewCommandHandler(user entities.User) *CommandHandler {
	c := &CommandHandler{
		Commands: make(map[string]func([]string)),
		User: user,
	}
	c.initCommands()
	return c
}


func (c *CommandHandler) initCommands() {
	c.Commands["create_board"] = c.create_board
	c.Commands["join"] = c.join_board
	c.Commands["list"] = c.list_board
	c.Commands["create_ticket"] = c.create_ticket
	c.Commands["list_tickets"] = c.list_tickets
	c.Commands["edit_ticket"] = c.edit_ticket
	c.Commands["examine_ticket"] = c.examine_ticket
	c.Commands["look_ticket"] = c.examine_ticket
	c.Commands["move_ticket"] = c.move_ticket

}


func (c *CommandHandler) create_board(input []string) {
		
//	knownBoards = append(knownBoards, )
//	r := regexp.MustCompile("[^\\s,]+")
//	cols := r.FindAllString()
	return
}

func (c *CommandHandler) join_board(input []string) {
	return
}

func (c *CommandHandler) list_board(input []string) {
	return
}

func (c *CommandHandler) create_ticket(input []string) {
	return
}

func (c *CommandHandler) list_tickets(input []string) {
	return
}

func (c *CommandHandler) edit_ticket(input []string) {
	return
}

func (c *CommandHandler) examine_ticket(input []string) {
	return
}

func (c *CommandHandler) move_ticket(input []string) {
	return
}


func helpMsgGeneral() string {
	return `
Welcome to your friendly neighborhood kanban board. Please use one of the following commands:
	1. /create_board <name> <column name 1, column name 2, column name 3>: To create a board with columns defined by a comma-separated list.
	2. /join <name>: To join a specific kanban board and view its tickets.
	3. /list: To list the various kanban boards you can join.
`
}


func helpMsgBoard() string {
	return `
	There are several board-related commands. Please use one of the following:
	1. /create_ticket <name> (optional)<due_date> <description> <assignee>: To create a ticket and automatically place it in the first column of the board. 
	2. /edit_ticket <name>: To update the various fields of the ticket.
	3. /list_tickets (optional)<column name>: To list all tickets in all columns of the board, or in one specific column.
	4. /examine_ticket <name>: To view all fields of a specific ticket, not just the name.
	5. /move_ticket <ticket_name> <column_name>: To move the specified ticket to the specified column.
`
}

func connect(s ssh.Session) {
	t := term.NewTerminal(s, fmt.Sprintf("%s > ", s.User()))
	for {
	line, err := t.ReadLine()
	if err != nil {
		break
		}

	if len(line) > 0 {
		if string(line[0]) == "/" {
			switch {
				case createBoard.MatchString(string(line)):
					return
				case joinBoard.MatchString(string(line)):
					return
				case listBoards.MatchString(string(line)):
					return
				case createTicket.MatchString(string(line)):
					return
				case listTickets.MatchString(string(line)):
					return
				case editTicket.MatchString(string(line)):
					return
				case moveTicket.MatchString(string(line)):
					return

				default:
					t.Write([]byte(helpMsgGeneral()))
				}
			}
		}
	}
} 

