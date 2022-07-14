package main

import "github.com/gliderlabs/ssh"

type room struct {
	users  map[string]*user
	status string
}

var anonymousUser = &user{
	id: "admin",
}

// Add user to room
func (r *room) addUser(u *user) {
	r.users[u.id] = u
	r.syncUI(anonymousUser)
}

// Remove user from room
func (r *room) removeUser(u user) {
	delete(r.users, u.id)
	r.syncUI(anonymousUser)
}

// Get user from room
func (r *room) getUser(s ssh.Session) *user {
	id := createId(s)
	user := r.users[id]

	return user
}

// Sync everybody's UI. Don't call program.Send on the user who triggered the sync in case it blocks their update method.
func (r *room) syncUI(owner *user) {
	for _, user := range r.users {
		if user.program != nil && owner.id != user.id {
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
