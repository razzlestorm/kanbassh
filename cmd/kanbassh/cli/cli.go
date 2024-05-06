package cli

import (
	"regexp"
	"github.com/razzlestorm/kanbassh/internal/entities"
)


var (
	createBoard = regexp.MustCompile(`^/create_board.*`)
	joinBoard = regexp.MustCompile(`^/join.*`)
	editBoard = regexp.MustCompile(`^/edit_board.*`)
	createTicket = regexp.MustCompile(`^/create_ticket.*`)
	editTicket = regexp.MustCompile(`^/edit_ticket.*`)
	moveTicket = regexp.MustCompile(`^/move_ticket.*`)
)
