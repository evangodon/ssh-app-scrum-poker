package main

import (
	"fmt"
	"strconv"

	lg "github.com/charmbracelet/lipgloss"
)

var (
	blue         = lg.Color("#89b4fa")
	mauve        = lg.Color("#8839ef")
	primaryColor = lg.Color("#EBA0AC")
	surface      = lg.Color("#6c7086")
)

// Container that wraps the whole app
func (m *model) NewContainer() lg.Style {
	return lg.NewStyle().
		Margin(0, 2).
		Padding(0, 2).
		Height(m.window.height - 2).
		Width(m.window.width - 5).
		BorderStyle(lg.RoundedBorder())
}

// List of users in room
func (m *model) listUsers() string {
	s := ""
	for _, user := range m.room.users {
		s += fmt.Sprintf("%s's vote: %d\n", user.name, user.vote)
	}

	container := lg.NewStyle().Height(10)

	return container.Render(s)
}

// Styling for app header
func (m *model) header() string {
	return lg.NewStyle().
		Bold(true).
		Foreground(blue).
		MarginBottom(1).
		Render("SSH Scrum Poker")
}

// Renders a card
func NewCard(option string, selected bool) string {
	style := lg.NewStyle().
		Padding(0, 1).
		MarginRight(1).
		BorderStyle(lg.RoundedBorder())

	if selected {
		selectedColor := primaryColor
		style = style.
			Bold(true).
			Foreground(selectedColor).
			BorderForeground(selectedColor).
			MarginBottom(1)
	}

	return style.Render(option)
}

var options = []int{0, 1, 2, 3, 5, 8}

// List all story point options
func (m *model) listOptions() string {
	cards := []string{}
	for _, option := range options {
		selected := option == m.user.vote
		o := strconv.Itoa(option)
		cards = append(cards, NewCard(o, selected))
	}

	s := lg.JoinHorizontal(lg.Center, cards...)
	container := lg.NewStyle().MarginLeft(m.window.width/2 - (lg.Width(s) / 2))

	return container.Render(s)
}

func (m *model) roomInfo() string {
	s := fmt.Sprintf("%d member%s in room\n", len(m.room.users), "s")

	return lg.NewStyle().Render(s)
}

const (
	logLimit = 6
)

func titleStyle() lg.Style {
	return lg.NewStyle().Background(surface).Padding(0, 1)
}

// Render a list of logs
func (m *model) showLogs() string {
	viewableLogs := m.logs
	numLogs := len(viewableLogs)

	if numLogs > logLimit {
		viewableLogs = m.logs[(numLogs - logLimit):numLogs]
	}

	s := titleStyle().Render("LOGS")
	s += "\n\n"
	for _, log := range viewableLogs {
		s += log
		s += "\n"
	}

	return s
}
