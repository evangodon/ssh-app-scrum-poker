package main

import (
	"fmt"

	"github.com/gliderlabs/ssh"
)

type room struct {
	users  map[string]*user
	status string
}

var anonymousUser = &user{
	id: "admin",
}

var noLog = roomLog{}

// Add user to room
func (r *room) addUser(u *user) {
	r.users[u.id] = u
	r.syncUI(anonymousUser, newRoomLog(fmt.Sprintf("→ %s joined room", u.name)))
}

// Remove user from room
func (r *room) removeUser(u user) {
	delete(r.users, u.id)
	r.syncUI(anonymousUser, newRoomLog(fmt.Sprintf("← %s left room", u.name)))
}

// Get user from room
func (r *room) getUser(s ssh.Session) *user {
	id := createId(s)
	user := r.users[id]

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

// Create a new room
func newRoom() room {
	users := make(map[string]*user)
	return room{
		users: users,
	}
}
