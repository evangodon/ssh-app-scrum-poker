package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/gliderlabs/ssh"
)

type room struct {
	users        []*user
	colors       []lg.Color
	displayVotes bool
}

var noOwner = &user{
	id: "admin",
}

var noLog = roomLog{}

// Add user to room
func (r *room) addUser(u *user) {
	u.color = r.GetUserColor()
	r.users = append(r.users, u)

	icon := makeGreen("[→")
	r.syncUI(noOwner, newRoomLog(fmt.Sprintf("%s %s joined room", icon, u.name)))
}

// Remove user from room
func (r *room) removeUser(u user) {
	for i, user := range r.users {
		if user.id == u.id {
			r.users = append(r.users[:i], r.users[i+1:]...)
		}
	}
	r.makeUserColorAvailable(u.color)

	icon := makeRed("←]")
	r.syncUI(noOwner, newRoomLog(fmt.Sprintf("%s %s left room", icon, u.name)))
}

func (r *room) GetUserColor() lg.Color {

	if len(r.colors) == 0 {
		return lg.Color("#bac2de")
	}

	selected, rest := r.colors[len(r.colors)-1], r.colors[:len(r.colors)-1]
	r.colors = rest

	return selected
}

func (r *room) makeUserColorAvailable(color lg.Color) {
	r.colors = append(r.colors, color)
}

// Get user from room
func (r *room) getUser(s ssh.Session) *user {
	id := createId(s)

	var user *user
	for _, u := range r.users {
		if u.id == id {
			user = u
		}
	}

	return user
}

// Calculate how many users votes
func (r *room) getNumberOfVotes() int {
	votes := 0

	for _, user := range r.users {
		if user.vote >= 0 {
			votes++
		}
	}

	return votes
}

// Sync everybody's UI.
func (r *room) syncUI(owner *user, log roomLog) {
	for _, user := range r.users {
		if user.program != nil && owner.id != user.id {
			user.program.Send(log)
		}
	}
}

func (r *room) startCountdownToDisplayVotes() tea.Msg {

	start := 3
	r.syncUI(noOwner, roomLog{log: "Revealing votes in..."})

	for i := start; i > 0; i-- {
		time.Sleep(1 * time.Second)
		r.syncUI(noOwner, roomLog{log: fmt.Sprintf("%d...", i)})
	}

	time.Sleep(1 * time.Second)

	r.displayVotes = true
	count := make(map[int]int)
	for _, user := range r.users {
		count[user.vote]++
	}

	log := "Breakdown of votes: \n"
	log += style().MarginLeft(2).Render(r.showVotesTable())

	r.syncUI(noOwner, roomLog{
		log:         log,
		clearBefore: true,
	})

	return nil
}

func (r *room) showVotesTable() string {
	count := map[int]int{}
	highestAmountOfVotes := 0

	for _, user := range r.users {
		count[user.vote] += 1
		if count[user.vote] > highestAmountOfVotes {
			highestAmountOfVotes = count[user.vote]
		}
	}

	const (
		columnKeyStoryPoints = "story points"
		columnKeyVotes       = "votes"
	)
	rows := []table.Row{}

	for storyPoint, votesCount := range count {
		style := lg.NewStyle()
		if votesCount == highestAmountOfVotes {
			style = style.Foreground(primaryColor)
		}

		voteCircles := strings.Repeat("● ", votesCount)

		rows = append(rows, table.NewRow(table.RowData{
			columnKeyStoryPoints: storyPoint,
			columnKeyVotes:       voteCircles,
		}).WithStyle(style))
	}

	t := table.New([]table.Column{
		table.NewColumn(columnKeyStoryPoints, "Story Points", 15),
		table.NewColumn(columnKeyVotes, "# Votes", 25),
	}).WithRows(rows).
		SortByDesc(columnKeyVotes).
		BorderRounded().
		SelectableRows(false).
		WithHighlightedRow(2)

	return t.View()
}

func (r *room) resetVotes() tea.Msg {
	r.displayVotes = false

	for _, user := range r.users {
		user.vote = -1
	}

	r.syncUI(noOwner, newRoomLog("All votes were reset"))
	return nil
}

// Create a new room
func newRoom() room {
	// users := []*user{}
	users := testusers

	return room{
		users: users,
		colors: []lg.Color{
			lg.Color("#f5c2e7"),
			lg.Color("#cba6f7"),
			lg.Color("#f38ba8"),
			lg.Color("#eba0ac"),
			lg.Color("#fab387"),
			lg.Color("#f9e2af"),
			lg.Color("#a6e3a1"),
			lg.Color("#94e2d5"),
			lg.Color("#74c7ec"),
			lg.Color("#b4befe"),
		},
		displayVotes: false,
	}
}
