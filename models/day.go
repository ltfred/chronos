package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ltfred/chronos/utils"
	"time"
)

type dayChoice struct {
	time       time.Time
	pos        int
	isSelected bool
	isInvalid  bool
}

type DayModel struct {
	choices []dayChoice
	cursor  int
}

func NewDayModel(year, month int) DayModel {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastDay := time.Date(year, time.Month(month)+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)

	week := int(firstDay.Weekday())
	if week == 0 {
		week = 7
	}
	dayModel := DayModel{}

	var pos int
	preMonthDays := make([]dayChoice, 0, week-1)
	if week <= 7 {
		for i := week - 1; i >= 1; i-- {
			preMonthDays = append(preMonthDays, dayChoice{
				time:      firstDay.AddDate(0, 0, -i),
				isInvalid: true,
				pos:       pos,
			})
			pos += 1
		}
	}
	curMonthDays := make([]dayChoice, 0, lastDay.Day())
	isSetMothFirstDaySelected := true
	for i := 1; i <= lastDay.Day(); i++ {
		cur := time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.Local)
		c := dayChoice{time: cur, pos: pos}
		y, m, d := time.Now().Date()
		if cur.Year() == y && cur.Month() == m {
			if cur.Day() == d {
				dayModel.cursor = c.pos
				isSetMothFirstDaySelected = false
			}
		}
		curMonthDays = append(curMonthDays, c)
		pos += 1
	}
	if isSetMothFirstDaySelected {
		curMonthDays[0].isSelected = true
	}
	nextMonthDays := make([]dayChoice, 0, 7-week)
	week = int(lastDay.Weekday())
	if week < 7 {
		for i := 1; i <= 7-week; i++ {
			nextMonthDays = append(nextMonthDays, dayChoice{
				time:      lastDay.AddDate(0, 0, i),
				isInvalid: true,
				pos:       pos,
			})
			pos += 1
		}
	}

	dayModel.choices = append(dayModel.choices, preMonthDays...)
	dayModel.choices = append(dayModel.choices, curMonthDays...)
	dayModel.choices = append(dayModel.choices, nextMonthDays...)

	return dayModel
}

func (m DayModel) Init() tea.Cmd {
	return nil
}

func (m DayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down", "j":
			if m.cursor+7 > len(m.choices) {
				return m, nil
			}
			m.cursor += 7
			return m, nil
		case "up", "k":
			if m.cursor-7 < 0 {
				return m, nil
			}
			m.cursor -= 7
			return m, nil
		case "left", "h":
			if m.cursor-1 < 0 {
				return m, nil
			}
			m.cursor -= 1
			return m, nil
		case "right", "l":
			if m.cursor+1 > len(m.choices) {
				return m, nil
			}
			m.cursor += 1
			return m, nil
		default:
		}
	}

	return m, nil
}

func (m DayModel) View() string {
	rowCount := len(m.choices) / 7
	rows := make([][]dayChoice, 0, rowCount)
	for i := 0; i < rowCount; i++ {
		rows = append(rows, m.choices[i*7:(i+1)*7])
	}
	onlyDayFormat := func(t time.Time) string {
		return t.Format("02")
	}
	joinHorizontal := func(choices []dayChoice) string {
		ss := make([]string, 0, 7)
		for _, v := range choices {
			calendar := utils.GetLunarCalendar(v.time)
			symbol := "\n"
			if utils.IsWeekend(v.time) {
				symbol = redTextStyle.Render("末")
			}
			if v.pos == m.cursor {
				s := focusedModelStyle.Render(boldTextStyle.Render(onlyDayFormat(v.time)), symbol, calendar)
				if v.isInvalid {
					s = focusedModelInvalidStyle.Render(grayTextStyle.Render(onlyDayFormat(v.time)), symbol, calendar)
				}
				ss = append(ss, s)
				continue
			}
			if v.isInvalid {
				ss = append(ss, modelStyle.Render(grayTextStyle.Render(onlyDayFormat(v.time)), symbol, calendar))
				continue
			}
			ss = append(ss, modelStyle.Render(boldTextStyle.Render(onlyDayFormat(v.time)), symbol, calendar))
		}
		return lipgloss.JoinHorizontal(lipgloss.Top, ss...)
	}
	horizontalJoins := make([]string, 0, rowCount+1)
	header := lipgloss.JoinHorizontal(
		lipgloss.Top, headerStyle.Render("星期一"), headerStyle.Render("星期二"),
		headerStyle.Render("星期三"), headerStyle.Render("星期四"), headerStyle.Render("星期五"),
		headerStyle.Render("星期六"), headerStyle.Render("星期日"),
	)
	horizontalJoins = append(horizontalJoins, header)
	for _, row := range rows {
		horizontalJoins = append(horizontalJoins, joinHorizontal(row))
	}

	s := lipgloss.JoinVertical(lipgloss.Top, horizontalJoins...)

	return lipgloss.JoinHorizontal(lipgloss.Top, s, focusedModelInvalidStyle.Render("日期详情"))
}
