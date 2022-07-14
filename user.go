package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
)

type user struct {
	id      string // user remote address + username
	name    string
	program *tea.Program
	vote    int
}

func createId(s ssh.Session) string {
	return fmt.Sprintf("%s.%s", s.RemoteAddr().String(), s.User())
}

func newUser(s ssh.Session) user {
	return user{
		id:   createId(s),
		name: s.User(),
		vote: -1,
	}
}

func (u *user) makeVote(vote int) {
	u.vote = vote
}
