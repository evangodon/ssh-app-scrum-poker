package main

import lg "github.com/charmbracelet/lipgloss"

var (
	blue         = lg.Color("#89b4fa")
	mauve        = lg.Color("#8839ef")
	primaryColor = lg.Color("#EBA0AC")
	surface      = lg.Color("#6c7086")
	green        = lg.Color("#a6e3a1")
	red          = lg.Color("#f38ba8")
)

var (
	appSideMargin = 2
)

// Create a new style object
func style() lg.Style {
	return lg.NewStyle()
}

func makeGreen(str string) string {
	return style().Foreground(green).Render(str)
}

func makeRed(str string) string {
	return style().Foreground(red).Render(str)
}

var debugMode = true

func getComponentBorder() lg.Border {
	if debugMode {
		return lg.NormalBorder()
	}

	return lg.HiddenBorder()
}

// Container that wraps the whole app
func (m *model) NewContainer() lg.Style {
	return style().
		Height(m.window.height - 4).
		Width(m.window.width)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pluralize(singular string, plural string, count int) string {
	if count > 1 {
		return plural
	}
	return singular
}
