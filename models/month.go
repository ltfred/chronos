package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

type monthChoice struct {
	time       time.Time
	pos        int
	isSelected bool
}

type MonthModel struct {
	choices []monthChoice
	cursor  int
}

func NewMonthModel(year int) MonthModel {
	choices := make([]monthChoice, 12)
	for i := 0; i < 12; i++ {
		choice := monthChoice{
			time: time.Date(year, time.Month(i+1), 1, 0, 0, 0, 0, time.Local),
			pos:  i,
		}
		now := time.Now()
		if now.Year() == year && now.Month() == time.Month(i+1) {
			choice.isSelected = true
		}
		choices[i] = choice
	}
	return MonthModel{choices: choices}
}

func (m MonthModel) Init() tea.Cmd {
	return nil
}

func (m MonthModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down", "j":
			if m.cursor+4 > 12 {
				return m, nil
			}
			m.cursor += 4
			return m, nil
		case "up", "k":
			if m.cursor-4 < 0 {
				return m, nil
			}
			m.cursor -= 4
			return m, nil
		case "left", "h":
			if m.cursor-1 < 0 {
				return m, nil
			}
			m.cursor -= 1
			return m, nil
		case "right", "l":
			if m.cursor+1 > 12 {
				return m, nil
			}
			m.cursor += 1
			return m, nil
		case "enter":
			dayModel := NewDayModel(m.choices[m.cursor].time.Year(), int(m.choices[m.cursor].time.Month()))
			return dayModel, nil
		default:
		}
	}

	return m, nil
}

func (m MonthModel) View() string {
	h1, h2, h3 := m.choices[:4], m.choices[4:8], m.choices[8:]
	joinHorizontal := func(choices []monthChoice) string {
		ss := make([]string, 0, 7)
		for _, v := range choices {
			if v.pos == m.cursor {
				ss = append(ss, monthFocusedModelStyle.Render(v.time.Format("2006-01")))
				continue
			}
			ss = append(ss, monthModelStyle.Render(v.time.Format("2006-01")))
		}
		return lipgloss.JoinHorizontal(lipgloss.Top, ss...)
	}
	s := lipgloss.JoinVertical(lipgloss.Top, joinHorizontal(h1), joinHorizontal(h2), joinHorizontal(h3))
	return s
}
