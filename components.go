package main

import (
	"fmt"
	"strconv"
	"strings"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

var userStyle = lg.NewStyle().Width(20).Padding(1, 0).Inline(true)

// List of users in room
func (m *model) listUsers() string {
	numUsers := len(m.room.users)
	pluralyze := ""
	if numUsers > 1 {
		pluralyze = "s"
	}
	s := fmt.Sprintf("%d member%s in room\n", numUsers, pluralyze)
	s = style().Width(m.window.width).Align(lg.Center).Render(s)
	s += "\n"

	col_1 := ""
	col_2 := ""
	col_3 := ""
	col_4 := ""

	for i, user := range m.room.users {
		isHost := (func() string {
			if user.isHost {
				return "(host) "
			}
			return ""
		})()
		username := userStyle.Render(fmt.Sprintf("%s %s", user.name, isHost))
		card := NewCardForUser(user.vote, m.room.displayVotes)
		order := fmt.Sprintf("%d. ", i+1)
		userBlock := lg.JoinHorizontal(lg.Center, order, username, card)

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

	container := lg.NewStyle().
		Width(m.window.width - 15)
	gap := strings.Repeat(" ", 5)
	s += lg.JoinHorizontal(lg.Top, col_1, gap, col_2, gap, col_3, gap, col_4)

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
			BorderForeground(selectedColor)
	}

	return style.Render(option)
}

func NewCardForUser(vote int, visible bool) string {
	v := strconv.Itoa(vote)

	if v == "-1" {
		return style().
			Inherit(cardStyle).
			BorderStyle(lg.HiddenBorder()).
			Render(" ")
	}

	if !visible {
		return cardStyle.Copy().Faint(true).Render("?")
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

	s := lg.JoinHorizontal(lg.Center, cards...)
	s = lg.JoinVertical(lg.Center, "Available Options", s)

	container := style().
		MarginLeft(m.window.width/2 - (lg.Width(s) / 2)).
		MarginBottom(1)

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

	container := style().Height(10)

	return container.Render(s)
}

func (r *room) showVotesTable() string {
	count := map[int]int{}
	mostVoted := 0

	for _, user := range r.users {
		count[user.vote] += 1
		if count[user.vote] > count[mostVoted] {
			mostVoted = user.vote
		}
	}

	const (
		columnKeyStoryPoints = "story points"
		columnKeyVotes       = "votes"
	)
	rows := []table.Row{}

	for storyPoint, numVotes := range count {
		style := lg.NewStyle()
		if storyPoint == mostVoted {
			style = style.Foreground(primaryColor)
		}
		rows = append(rows, table.NewRow(table.RowData{
			columnKeyStoryPoints: storyPoint,
			columnKeyVotes:       numVotes,
		}).WithStyle(style))
	}

	t := table.New([]table.Column{
		table.NewColumn(columnKeyStoryPoints, "Story Point", 15),
		table.NewColumn(columnKeyVotes, "# Votes", 15),
	}).WithRows(rows).
		SortByDesc(columnKeyVotes).
		BorderRounded().
		SelectableRows(false).
		WithHighlightedRow(2)

	return t.View()
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
	return helpStyle.Render(s)
}
