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

// Create a new style object
func style() lg.Style {
	return lg.NewStyle()
}

var debugMode = false

func getComponentBorder() lg.Border {
	if debugMode {
		return lg.NormalBorder()
	}

	return lg.HiddenBorder()
}

// Container that wraps the whole app
func (m *model) NewContainer() lg.Style {
	return style().
		Margin(0, 2).
		Padding(0, 2).
		Height(m.window.height - 2).
		Width(m.window.width - 5).
		BorderStyle(lg.RoundedBorder()).
		BorderForeground(surface)
}

var userStyle = lg.NewStyle().Width(25).Padding(1, 0)

// List of users in room
func (m *model) listUsers() string {
	numUsers := len(m.room.users)
	pluralyze := ""
	if numUsers > 1 {
		pluralyze = "s"
	}
	s := fmt.Sprintf("%d member%s in room\n", numUsers, pluralyze)
	s += "\n"

	leftCol := ""
	middleCol := ""
	rightCol := ""

	for i, user := range m.room.users {
		username := userStyle.Render(user.name)
		card := NewCardForUser(user.vote, m.room.displayVotes)
		order := fmt.Sprintf("%d. ", i+1)
		userBlock := lg.JoinHorizontal(lg.Center, order, username, card)

		if i < 4 {
			leftCol += userBlock
			leftCol += "\n"
		} else if i < 8 {
			middleCol += userBlock
			middleCol += "\n"
		} else {
			rightCol += userBlock
			rightCol += "\n"
		}
	}

	container := lg.NewStyle().BorderStyle(getComponentBorder()).
		Width(m.window.width - 15)
	gap := strings.Repeat(" ", 10)
	s += lg.JoinHorizontal(lg.Top, leftCol, gap, middleCol, gap, rightCol)

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
	s = lg.JoinVertical(lg.Center, "Available Options", s)

	container := style().
		BorderStyle(getComponentBorder()).
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

	container := style().Height(10).BorderStyle(getComponentBorder())

	return container.Render(s)
}
