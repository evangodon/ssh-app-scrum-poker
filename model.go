package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type window struct {
	height int
	width  int
}

type sectionHeight struct {
	app     int
	header  int
	users   int
	options int
	help    int
}

type model struct {
	user          *user
	room          *room
	window        window
	sectionHeight sectionHeight
	logs          []string
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
		case "h", "j", "k", "l", ";", "'":
			{
				v := map[string]int{
					"h": 0,
					"j": 1,
					"k": 2,
					"l": 3,
					";": 5,
					"'": 8,
				}[msg.String()]

				if m.user.vote == v {
					m.user.vote = -1
				} else {
					m.user.makeVote(v)
				}
			}
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
			if m.user.isHost {
				notVoted := 0
				for _, user := range m.room.users {
					if user.vote < 0 {
						notVoted++
					}
				}

				if notVoted > 0 {
					members := pluralize("member has", "members have", notVoted)
					str := fmt.Sprintf("%d %s not voted yet", notVoted, members)
					log := newRoomLog(str)

					m.logs = append(m.logs, log.log)
					return m, nil
				}
				return m, m.room.startCountdownToDisplayVotes
			}

		case "R":
			if m.user.isHost {
				return m, m.room.resetVotes
			}

		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case roomLog:
		if msg.log != "" {
			if msg.clearBefore {
				println(" clearing")
				m.logs = make([]string, 0)
				m.logs = append(m.logs, msg.log)
			} else {
				println("not clearing")
				m.logs = append(m.logs, msg.log)
			}
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	container := m.NewContainer()

	sections := strings.Builder{}
	sections.WriteString(m.header())
	sections.WriteString("\n")
	sections.WriteString(m.listOptions())
	sections.WriteString("\n")
	sections.WriteString(m.listUsers())
	sections.WriteString("\n\n")

	sections.WriteString(m.showLogs())
	sections.WriteString("\n")

	app := container.Render(sections.String())

	ui := strings.Builder{}
	ui.WriteString(app)

	ui.WriteString("\n")
	ui.WriteString(m.showHelp())

	return ui.String()
}
