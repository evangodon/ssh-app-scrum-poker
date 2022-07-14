package main

import (
	"fmt"
	"strconv"

	lg "github.com/charmbracelet/lipgloss"
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

	return s
}

// Styling for app header
func (m *model) header() string {
	return lg.NewStyle().Bold(true).MarginBottom(1).Render("Scrum Poker")
}

// Renders a card
func NewCard(option string, selected bool) string {
	style := lg.NewStyle().
		Padding(0, 1).
		Margin(0, 1).
		BorderStyle(lg.RoundedBorder())

	if selected {
		selectedColor := lg.Color("#EBA0AC")
		style = style.
			Bold(true).
			Foreground(selectedColor).
			BorderForeground(selectedColor).
			MarginBottom(1)
	}

	return style.Render(option)
}

func (m *model) roomInfo() string {
	s := fmt.Sprintf("%d member%s in room\n", len(m.room.users), "s")

	return lg.NewStyle().Render(s)
}

// List all story point options
func (m *model) listOptions(options []int) string {
	cards := []string{
		"Select an option:",
	}
	for _, option := range options {
		selected := option == m.user.vote
		o := strconv.Itoa(option)
		cards = append(cards, NewCard(o, selected))
	}

	return lg.JoinHorizontal(lg.Center, cards...)
}

// Render a list of logs
func (m *model) showLogs() string {

	return ""
}
