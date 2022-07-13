package main

type model struct {
	user *user
	room *room
}

func newModel(u *user, r *room) model {

	return model{
		user: u,
		room: r,
	}
}
