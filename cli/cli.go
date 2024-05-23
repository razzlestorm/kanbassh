package cli

import (
	"regexp"
	"github.com/razzlestorm/kanbassh/internal/entities"
)


var (
	createBoard = regexp.MustCompile(`^/create_board.*`)
	joinBoard = regexp.MustCompile(`^/join.*`)
	listBoards = regexp.MustCompile(`^/list.*`)
	editBoard = regexp.MustCompile(`^/edit_board.*`)
	createTicket = regexp.MustCompile(`^/create_ticket.*`)
	listTickets = regexp.MustCompile(`^/list_tickets.*`)
	editTicket = regexp.MustCompile(`^/edit_ticket.*`)
	moveTicket = regexp.MustCompile(`^/move_ticket.*`)
)

