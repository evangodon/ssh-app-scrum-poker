package main

type user struct {
	id   string
	name string
}

type shared struct {
	users []user
}

type model struct {
	term   string
	width  int
	height int
	shared shared
}

func newModel() model{

  return model{}
}
