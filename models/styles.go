package models

import "github.com/charmbracelet/lipgloss"

// day
var (
	grayTextStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#808080", Dark: "#808080"})
	redTextStyle  = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#FF0000", Dark: "#FF0000"}).Bold(true)
	boldTextStyle = lipgloss.NewStyle().Bold(true)

	modelStyle = lipgloss.NewStyle().
			Width(9).
			Height(2).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().
				Width(9).
				Height(2).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("#DE684F"))
	focusedModelInvalidStyle = lipgloss.NewStyle().
					Width(9).
					Height(2).
					Align(lipgloss.Center, lipgloss.Center).
					BorderStyle(lipgloss.NormalBorder()).
					BorderForeground(lipgloss.Color("#808080"))
	todayStyle = lipgloss.NewStyle().
			Width(9).
			Height(2).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#DE684F"))
	headerStyle = lipgloss.NewStyle().
			Width(9).
			Height(1).
			Bold(true).
			BorderStyle(lipgloss.HiddenBorder()).
			Align(lipgloss.Center, lipgloss.Center)
	dayDetailStyle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Left, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
)

// month
var (
	monthModelStyle = lipgloss.NewStyle().
			Width(9).
			Height(3).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	monthFocusedModelStyle = lipgloss.NewStyle().
				Width(9).
				Height(3).
				Bold(true).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("#DE684F"))
)

var (
	cursorLineStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("57")).
			Foreground(lipgloss.Color("230"))
	textAreaFocusStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("238"))
)
