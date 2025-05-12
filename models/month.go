package models

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dromara/carbon/v2"
	"github.com/ltfred/chronos/internal"
	"time"
)

type monthChoice struct {
	time time.Time
	pos  int
}

type monthKeymap struct {
	up, down, left, right, enter, quit key.Binding
}

type MonthModel struct {
	choices []monthChoice
	cursor  int
	help    help.Model
	keymap  monthKeymap
}

func NewMonthModel(year int, month time.Month) MonthModel {
	choices := make([]monthChoice, 12)
	monthModel := MonthModel{help: help.New()}
	for i := 0; i < 12; i++ {
		t := time.Date(year, time.Month(i+1), 1, 0, 0, 0, 0, time.Local)
		choice := monthChoice{
			time: t,
			pos:  i,
		}
		if t.Year() == year && t.Month() == month {
			monthModel.cursor = i
		}
		choices[i] = choice
	}
	monthModel.choices = choices

	k := monthKeymap{
		up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "left"),
		),
		right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "right"),
		),
		enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "days"),
		),
		quit: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quit"),
		),
	}
	monthModel.keymap = k
	return monthModel
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
	helpMsg := m.help.ShortHelpView([]key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.left,
		m.keymap.right,
		m.keymap.enter,
		m.keymap.quit,
	})
	h1, h2, h3 := m.choices[:4], m.choices[4:8], m.choices[8:]
	joinHorizontal := func(choices []monthChoice) string {
		ss := make([]string, 0, 7)
		for _, v := range choices {
			if v.pos == m.cursor {
				ss = append(ss, monthFocusedModelStyle.Render(m.genMonthStr(v.time.Year(), v.time.Month())))
				continue
			}
			ss = append(ss, monthModelStyle.Render(m.genMonthStr(v.time.Year(), v.time.Month())))
		}
		return lipgloss.JoinHorizontal(lipgloss.Top, ss...)
	}
	s := lipgloss.JoinVertical(lipgloss.Top, joinHorizontal(h1), joinHorizontal(h2), joinHorizontal(h3)) + "\n\n" + helpMsg
	return s
}

func (m MonthModel) genMonthStr(year int, month time.Month) string {
	f := func(days []internal.Day) bool {
		for _, day := range days {
			parse, _ := time.Parse("01-02", day.Date)
			if day.IsLunar {
				mm := carbon.CreateFromLunar(year, int(parse.Month()), parse.Day(), 0, 0, 0, false).Month()
				if mm == int(month) {
					return true
				}
			} else {
				if parse.Month() == month {
					return true
				}
			}
		}
		return false
	}
	if f(internal.Cfg.ImportantDay.Birthdays) {
		return fmt.Sprintf("🌟%d-%d", year, month)
	}
	if f(internal.Cfg.ImportantDay.MemorialDays) {
		return fmt.Sprintf("🌟%d-%02d", year, month)
	}
	return fmt.Sprintf("%d-%02d", year, month)
}
