package models

import "github.com/charmbracelet/lipgloss"

var (
	grayTextStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#808080",
		Dark:  "#808080",
	})

	modelStyle = lipgloss.NewStyle().
			Width(15).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(5).
				Bold(true).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
	focusedModelInvalidStyle = lipgloss.NewStyle().
					Width(15).
					Height(5).
					Align(lipgloss.Center, lipgloss.Center).
					BorderStyle(lipgloss.NormalBorder()).
					BorderForeground(lipgloss.Color("#808080"))
	headerStyle = lipgloss.NewStyle().
			Width(15).
			Height(1).
			Bold(true).
			BorderStyle(lipgloss.NormalBorder()).
			Align(lipgloss.Center, lipgloss.Center)
)
