package main

import (
	"fmt"
	"strconv"
	"strings"

	lg "github.com/charmbracelet/lipgloss"
)

var userStyle = lg.NewStyle().Width(14).Padding(1, 0).Inline(true)

// List of users in room
func (m *model) listUsers() string {
	numUsers := len(m.room.users)
	members := pluralize("member", "members", m.room.getNumberOfVotes())
	s := fmt.Sprintf(
		"%d %s in room • %d voted \n",
		numUsers,
		members,
		m.room.getNumberOfVotes(),
	)
	s += "\n"

	col_1 := ""
	col_2 := ""
	col_3 := ""
	col_4 := ""

	for i, user := range m.room.users {
		if user.id == m.user.id {
			continue
		}
		card := NewCardForUser(user, m.room.displayVotes)
		userLine := newUserLine(user)
		userBlock := lg.JoinHorizontal(lg.Center, userLine, card)

		if i < 3 {
			col_1 += userBlock
			col_1 += "\n"
		} else if i < 6 {
			col_2 += userBlock
			col_2 += "\n"
		} else if i < 9 {
			col_3 += userBlock
			col_3 += "\n"
		} else {
			col_4 += userBlock
			col_4 += "\n"
		}
	}

	if len(m.room.users) == 1 {
		col_1 = style().Faint(true).Italic(true).Render("Nobody else is in the room")
	}

	container := lg.NewStyle().
		Border(lg.RoundedBorder()).
		BorderForeground(surface).
		Margin(0, appSideMargin).
		Padding(0, 2).
		Width(m.window.width - 12)
	gap := strings.Repeat(" ", 10)
	s += lg.JoinHorizontal(lg.Top, col_1, gap, col_2, gap, col_3, gap, col_4)

	str := container.Render(s)
	m.sectionHeight.users = lg.Height(str)

	return str
}

func newUserLine(u *user) string {
	isHost := (func() string {
		if u.isHost {
			return "(host) "
		}
		return ""
	})()
	displayedName := func() string {
		if len(u.name) > 11 {
			return u.name[0:11] + "…"
		}
		return u.name
	}()
	username := userStyle.Render(fmt.Sprintf("%s %s", displayedName, isHost))
	userColor := style().Foreground(u.color).Render("●")

	return fmt.Sprintf("%s %s", userColor, username)
}

// Styling for app header
func (m *model) header() string {
	str := lg.NewStyle().
		Bold(true).
		Foreground(blue).
		Margin(0, appSideMargin, 1, appSideMargin).
		Render("SSH Scrum Poker")

	m.sectionHeight.header = lg.Height(str)

	return str
}

var cardStyle = lg.NewStyle().
	Padding(0, 1).
	MarginRight(1).
	BorderStyle(lg.RoundedBorder())

// Renders a card to indicate vote selection
func NewCardForSelection(option string, selected bool, user *user) string {
	style := cardStyle.Copy()

	if selected {
		selectedColor := user.color
		style = style.
			Bold(true).
			Foreground(selectedColor).
			BorderForeground(selectedColor).
			BorderStyle(lg.DoubleBorder())
	}

	return style.Render(option)
}

func NewCardForUser(user *user, visible bool) string {
	v := strconv.Itoa(user.vote)

	if v == "-1" {
		return style().
			Inherit(cardStyle).
			BorderStyle(lg.HiddenBorder()).
			Render(" ")
	}

	userCardStyle := cardStyle.Copy().BorderForeground(user.color).Foreground(user.color)

	if !visible {
		return userCardStyle.Faint(true).Render("?")
	}

	return userCardStyle.Render(v)
}

var options = []int{0, 1, 2, 3, 5, 8}

// List all story point options
func (m *model) listOptions() string {
	cards := []string{}
	for _, option := range options {
		selected := option == m.user.vote
		o := strconv.Itoa(option)
		cards = append(cards, NewCardForSelection(o, selected, m.user))
	}

	cardBlock := lg.JoinHorizontal(lg.Center, cards...)
	cardBlock = lg.JoinVertical(lg.Center, "Your vote", cardBlock)

	userLine := newUserLine(m.user)
	gap := strings.Repeat(" ", 8)
	s := lg.JoinHorizontal(lg.Center, userLine, gap, cardBlock)

	container := style().
		Border(lg.RoundedBorder()).
		BorderForeground(surface).
		Padding(0, 2).
		Margin(0, appSideMargin)

	str := container.Render(s)
	m.sectionHeight.options = lg.Height(str)

	return str
}

func (m *model) roomInfo() string {
	s := fmt.Sprintf("%d member%s in room\n", len(m.room.users), "s")

	return lg.NewStyle().Render(s)
}

// Render a list of logs
func (m *model) showLogs() string {
	displayedLogs := m.logs
	numDisplayedLogs := len(displayedLogs)

	height := m.window.height -
		m.sectionHeight.header -
		m.sectionHeight.users -
		m.sectionHeight.options -
		m.sectionHeight.help
	logSectionHeight := Abs(height - 5)
	container := style().Height(logSectionHeight)

	maxAvailableLines := Abs(logSectionHeight - 2)

	if numDisplayedLogs > maxAvailableLines {
		displayedLogs = displayedLogs[(numDisplayedLogs - maxAvailableLines):numDisplayedLogs]
	}

	left := "── Logs "
	lineRepeated := func() int {
		if m.window.width > len(left) {
			return m.window.width - len(left)
		}
		return m.window.width
	}()
	right := strings.Repeat("─", lineRepeated)
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%s%s\n\n", left, right))

	for _, log := range displayedLogs {
		s.WriteString(fmt.Sprintf("  %s \n", log))
	}

	return container.Render(s.String())
}

func (m *model) showHelp() string {
	helpStyle := style().
		Faint(true).
		Align(lg.Center).
		Width(m.window.width)
	s := "[0-8]: Vote • q: quit"
	if m.user.isHost {
		s += " • V: reveal votes • R: reset votes"
	}
	str := helpStyle.Render(s)

	m.sectionHeight.help = lg.Height(str)

	return str
}
