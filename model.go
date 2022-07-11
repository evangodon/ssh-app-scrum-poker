package main

type user struct {
	id   string // user remote address + username
	name string
}

type room struct {
	users []user
  status string
}

type model struct {
	term   string
	width  int
	height int
	room room
}

func newModel(u user) model{

  return model{
  }
}
