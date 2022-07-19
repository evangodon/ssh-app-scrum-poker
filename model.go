package main

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type window struct {
	height int
	width  int
}

type model struct {
	user   *user
	room   *room
	window window
	logs   []string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.window.height = msg.Height
		m.window.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "0", "1", "2", "3", "5", "8":
			{
				v, _ := strconv.Atoi(msg.String())
				if m.user.vote == v {
					m.user.vote = -1
				} else {
					m.user.makeVote(v)
				}

				m.room.syncUI(m.user, noLog)
			}

		case "V":
			return m, m.room.startCountdownToDisplayVotes

		case "R":
			return m, m.room.resetVotes

		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case roomLog:
		m.logs = append(m.logs, msg.log)
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	container := m.NewContainer()

	sb := strings.Builder{}
	sb.WriteString(m.header())
	sb.WriteString(m.listUsers())
	sb.WriteString("\n")
	sb.WriteString(m.listOptions())
	sb.WriteString("\n")

	sb.WriteString(m.showLogs())

	return container.Render(sb.String())
}
