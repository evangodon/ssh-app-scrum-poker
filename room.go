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

var noOwner = &user{
	id: "admin",
}

var noLog = roomLog{}

// Add user to room
func (r *room) addUser(u *user) {
	r.users = append(r.users, u)
	icon := makeGreen("[→")
	r.syncUI(noOwner, newRoomLog(fmt.Sprintf("%s %s joined room", icon, u.name)))
}

// Remove user from room
func (r *room) removeUser(u user) {
	for i, user := range r.users {
		if user.id == u.id {
			r.users = append(r.users[:i], r.users[i+1:]...)
		}
	}

	icon := makeRed("←]")
	r.syncUI(noOwner, newRoomLog(fmt.Sprintf("%s %s left room", icon, u.name)))
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

// Calculate how many users votes
func (r *room) getNumberOfVotes() int {
	votes := 0

	for _, user := range r.users {
		if user.vote >= 0 {
			votes++
		}
	}

	return votes
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
	r.syncUI(noOwner, roomLog{log: "Revealing votes in..."})

	for i := start; i > 0; i-- {
		time.Sleep(1 * time.Second)
		r.syncUI(noOwner, roomLog{log: fmt.Sprintf("%d...", i)})
	}

	time.Sleep(1 * time.Second)

	r.displayVotes = true
	count := make(map[int]int)
	for _, user := range r.users {
		count[user.vote]++
	}

	log := "Breakdown of votes: \n"
	log += r.showVotesTable()

	r.syncUI(noOwner, roomLog{
		log:         log,
		clearBefore: true,
	})

	return nil
}

func (r *room) resetVotes() tea.Msg {
	r.displayVotes = false

	for _, user := range r.users {
		user.vote = -1
	}

	r.syncUI(noOwner, newRoomLog("All votes were reset"))
	return nil
}

// Create a new room
func newRoom() room {
	users := []*user{}
	// users := testusers

	return room{
		users:        users,
		displayVotes: false,
	}
}
