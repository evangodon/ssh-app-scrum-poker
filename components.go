package main

import (
	"fmt"
	"strconv"
	"strings"

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

var userStyle = lg.NewStyle().Width(25).Padding(1, 0)

// List of users in room
func (m *model) listUsers() string {
	s := fmt.Sprintf("%d member%s in room\n", len(m.room.users), "s")
	s += "\n"

	leftCol := ""
	rightCol := ""

	for i, user := range m.room.users {
		username := userStyle.Render(user.name)
		card := NewCardForUser(user.vote, m.room.displayVotes)
		order := fmt.Sprintf("%d. ", i)
		userBlock := lg.JoinHorizontal(lg.Center, order, username, card)

		if i < 5 {
			leftCol += userBlock
			leftCol += "\n"
		} else {
			rightCol += userBlock
			rightCol += "\n"
		}
	}

	container := lg.NewStyle().
		Width(m.window.width - 15).
		MarginBottom(2)
	gap := strings.Repeat(" ", 30)
	s += lg.JoinHorizontal(lg.Top, leftCol, gap, rightCol)

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

var cardStyle = lg.NewStyle().
	Padding(0, 1).
	MarginRight(1).
	BorderStyle(lg.RoundedBorder())

// Renders a card to indicate vote selection
func NewCardForSelection(option string, selected bool) string {
	style := cardStyle.Copy()

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

func NewCardForUser(vote int, visible bool) string {
	v := strconv.Itoa(vote)

	if v == "-1" {
		return ""
	}

	if !visible {
		return cardStyle.Render("▒")
	}

	return cardStyle.Render(v)
}

var options = []int{0, 1, 2, 3, 5, 8}

// List all story point options
func (m *model) listOptions() string {
	cards := []string{}
	for _, option := range options {
		selected := option == m.user.vote
		o := strconv.Itoa(option)
		cards = append(cards, NewCardForSelection(o, selected))
	}

	s := lg.JoinHorizontal(lg.Top, cards...)
	container := lg.NewStyle().
		MarginLeft(m.window.width/2 - (lg.Width(s) / 2))

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
