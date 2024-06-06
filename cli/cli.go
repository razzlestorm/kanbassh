package cli

import (
	"fmt"
	"regexp"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
	"github.com/razzlestorm/kanbassh/internal/entities"
)


var (
	boards map[ssh.Session]*entities.Board
	knownBoards []*entities.Board
	createBoard = regexp.MustCompile(`^/create_board.*`)
	joinBoard = regexp.MustCompile(`^/join.*`)
	listBoards = regexp.MustCompile(`^/list.*`)
	createTicket = regexp.MustCompile(`^/create_ticket.*`)
	listTickets = regexp.MustCompile(`^/list_tickets.*`)
	editTicket = regexp.MustCompile(`^/edit_ticket.*`)
	examineTicket = regexp.MustCompile(`^/[examine_ticket]|[look_ticket].*`)
	moveTicket = regexp.MustCompile(`^/move_ticket.*`)
)

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

