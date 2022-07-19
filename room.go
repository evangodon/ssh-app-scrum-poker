package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
)

type room struct {
	users        []*user
	displayVotes bool
}

var anonymousUser = &user{
	id: "admin",
}

var noLog = roomLog{}

// Add user to room
func (r *room) addUser(u *user) {
	r.users = append(r.users, u)
	r.syncUI(anonymousUser, newRoomLog(fmt.Sprintf("→ %s joined room", u.name)))
}

// Remove user from room
func (r *room) removeUser(u user) {
	for i, user := range r.users {
		if user.id == u.id {
			r.users = append(r.users[:i], r.users[i+1:]...)
		}
	}

	r.syncUI(anonymousUser, newRoomLog(fmt.Sprintf("← %s left room", u.name)))
}

// Get user from room
func (r *room) getUser(s ssh.Session) *user {
	id := createId(s)

	var user *user
	for _, u := range r.users {
		if u.id == id {
			user = u
		}
	}

	return user
}

// Sync everybody's UI. Don't call program.Send on the user who triggered
// the sync since it will block their update method.
func (r *room) syncUI(owner *user, log roomLog) {
	for _, user := range r.users {
		if user.program != nil && owner.id != user.id {
			user.program.Send(log)
		}
	}
}

func (r *room) startCountdownToDisplayVotes() tea.Msg {
	start := 3
	r.syncUI(anonymousUser, roomLog{log: "Revealing votes in..."})

	for i := start; i > 0; i-- {
		time.Sleep(1 * time.Second)
		r.syncUI(anonymousUser, roomLog{log: fmt.Sprintf("%d...", i)})
	}

	r.displayVotes = true
	count := make(map[int]int)
	for _, user := range r.users {
		count[user.vote]++
	}

	t := "Breakdown of votes: \n"
	for vote, numOfVotes := range count {
		t += fmt.Sprintf("%d: %d votes", vote, numOfVotes)
		t += "\n"
	}

	r.syncUI(anonymousUser, roomLog{log: t})

	return nil
}

func (r *room) resetVotes() tea.Msg {
	r.displayVotes = false

	for _, user := range r.users {
		user.vote = -1
	}

	return roomLog{
		log: "All votes were reset",
	}
}

// Create a new room
func newRoom() room {
	// users := make(map[string]*user)
	users := []*user{
		{
			id:      "id123",
			name:    "Joe",
			program: nil,
			vote:    8,
		},
		{
			id:      "id124",
			name:    "Tracy",
			program: nil,
			vote:    3,
		},
		{
			id:      "id125",
			name:    "Kim",
			program: nil,
			vote:    1,
		},
		{
			id:      "id126",
			name:    "Steve",
			program: nil,
			vote:    2,
		},
		{
			id:      "id126",
			name:    "Dustan",
			program: nil,
			vote:    3,
		},
		{
			id:      "id193",
			name:    "Will",
			program: nil,
			vote:    3,
		},
		{
			id:      "id126",
			name:    "Micheal",
			program: nil,
			vote:    3,
		},
		{
			id:      "id1234",
			name:    "John",
			program: nil,
			vote:    3,
		},
		{
			id:      "id231",
			name:    "Sara",
			program: nil,
			vote:    3,
		},
	}

	return room{
		users:        users,
		displayVotes: false,
	}
}
