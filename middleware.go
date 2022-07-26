package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/gliderlabs/ssh"
	"github.com/muesli/termenv"
)

var welcomeMsg = newRoomLog("Welcome to SSH Scrum Poker 👋. Use your num keys to vote.\n")

func newModel(u *user, r *room) model {
	return model{
		user: u,
		room: r,
		logs: []string{welcomeMsg.log},
	}
}

func customBubbleteaMiddleware(room *room) wish.Middleware {
	newProg := func(m tea.Model, opts ...tea.ProgramOption) *tea.Program {
		p := tea.NewProgram(m, opts...)
		return p
	}
	teaHandler := func(s ssh.Session) *tea.Program {
		_, _, active := s.Pty()
		if !active {
			fmt.Println("no active terminal, skipping")
			s.Exit(1)
			return nil
		}
		user := room.getUser(s)
		m := newModel(user, room)

		program := newProg(m, tea.WithInput(s), tea.WithOutput(s), tea.WithAltScreen())
		user.program = program

		return program
	}
	return bm.MiddlewareWithProgramHandler(teaHandler, termenv.ANSI256)
}

// Manage members connecting and leaving room
func manageRoomMembers(room *room) wish.Middleware {
	return func(h ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			user := newUser(s)
			if len(room.users) == 0 {
				user.isHost = true
			}
			room.addUser(&user)
			h(s)
			room.removeUser(user)
		}
	}
}
