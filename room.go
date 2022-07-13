package main

import "github.com/gliderlabs/ssh"

type room struct {
	users  map[string]*user
	status string
}

// Add user to room
func (r *room) addUser(u *user) {
	r.users[u.id] = u
	r.syncRoom()
}

// Remove user from room
func (r *room) removeUser(u user) {
	delete(r.users, u.id)
	r.syncRoom()
}

// Get user from room
func (r *room) getUser(s ssh.Session) *user {
	id := createId(s)
	user := r.users[id]

	return user
}

func (r *room) syncRoom() {
	for _, user := range r.users {
		if user.program != nil {
			user.program.Send("")
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
