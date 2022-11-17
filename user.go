package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/gliderlabs/ssh"
)

type user struct {
	id      string // user remote address + username
	name    string
	program *tea.Program
	vote    int
	isHost  bool
	color   lg.Color
}

func createID(s ssh.Session) string {
	return fmt.Sprintf("%s.%s", s.RemoteAddr().String(), s.User())
}

func newUser(s ssh.Session) user {
	return user{
		id:     createID(s),
		name:   s.User(),
		vote:   -1,
		isHost: false,
	}
}

func (u *user) makeVote(vote int) {
	u.vote = vote
}
