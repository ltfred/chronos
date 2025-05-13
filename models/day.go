package models

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dromara/carbon/v2"
	"github.com/ltfred/chronos/constants"
	"github.com/ltfred/chronos/internal"
	"github.com/ltfred/chronos/utils"
	"strings"
	"time"
)

type dayChoice struct {
	time      time.Time
	pos       int
	isInvalid bool
}

type dayKeymap struct {
	preMonth, nextMonth, month key.Binding
	moveKeymap
	quitKeymap
	todayKeymap
}

type DayModel struct {
	choices     []dayChoice
	month, year int
	cursor      int
	help        help.Model
	keymap      dayKeymap
}

func NewDayModel(year, month int) DayModel {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastDay := time.Date(year, time.Month(month)+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
	now := time.Now()

	week := int(firstDay.Weekday())
	if week == 0 {
		week = 7
	}
	dayModel := DayModel{month: month, year: year, help: help.New()}

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
		if cur.Year() == now.Year() && cur.Month() == now.Month() {
			if cur.Day() == now.Day() {
				dayModel.cursor = c.pos
				isSetMothFirstDaySelected = false
			}
		}
		curMonthDays = append(curMonthDays, c)
		pos += 1
	}
	if isSetMothFirstDaySelected {
		dayModel.cursor = curMonthDays[0].pos
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

	k := dayKeymap{
		moveKeymap: newMoveKeymap(),
		quitKeymap: newQuitKeymap(),
		preMonth: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "pre month"),
		),
		nextMonth: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "next month"),
		),
		month: key.NewBinding(
			key.WithKeys("m"),
			key.WithHelp("m", "month"),
		),
		todayKeymap: newTodayKeymap(),
	}
	dayModel.keymap = k

	return dayModel
}

func (m DayModel) Init() tea.Cmd {
	return nil
}

func (m DayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.up):
			if m.cursor-7 < 0 {
				return m, nil
			}
			m.cursor -= 7
			return m, nil
		case key.Matches(msg, m.keymap.down):
			if m.cursor+7 >= len(m.choices) {
				return m, nil
			}
			m.cursor += 7
			return m, nil
		case key.Matches(msg, m.keymap.left):
			if m.cursor-1 < 0 {
				return m, nil
			}
			m.cursor -= 1
			return m, nil
		case key.Matches(msg, m.keymap.right):
			if m.cursor+1 >= len(m.choices) {
				return m, nil
			}
			m.cursor += 1
			return m, nil
		case key.Matches(msg, m.keymap.today):
			return NewDayModel(time.Now().Year(), int(time.Now().Month())), nil
		}
		switch msg.String() {
		case "m":
			t := m.choices[m.cursor].time
			return NewMonthModel(t.Year(), t.Month()), nil
		case "n":
			if m.month+1 > 12 {
				return m, nil
			}
			return NewDayModel(m.year, m.month+1), nil
		case "p":
			if m.month-1 < 1 {
				return m, nil
			}
			return NewDayModel(m.year, m.month-1), nil
		default:
		}
	}

	return m, nil
}

func (m DayModel) View() string {
	helpMsg := m.help.ShortHelpView([]key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.left,
		m.keymap.right,
		m.keymap.preMonth,
		m.keymap.nextMonth,
		m.keymap.month,
		m.keymap.today,
		m.keymap.quit,
	})

	rowCount := len(m.choices) / 7
	rows := make([][]dayChoice, 0, rowCount)
	for i := 0; i < rowCount; i++ {
		rows = append(rows, m.choices[i*7:(i+1)*7])
	}
	var selectDay dayChoice
	joinHorizontal := func(choices []dayChoice) string {
		ss := make([]string, 0, 7)
		for _, v := range choices {
			day := utils.OnlyDayFormat(v.time)
			workState, lunar := m.getWorkAndLunar(v.time)
			importantDayMark := m.getImportantDayMark(v.time)
			if v.pos == m.cursor {
				selectDay = v
				s := focusedModelStyle.Render(importantDayMark, boldTextStyle.Render(day), workState, lunar)
				if v.isInvalid {
					s = focusedModelInvalidStyle.Render(importantDayMark, grayTextStyle.Render(day), workState, grayTextStyle.Render(lunar))
				}
				ss = append(ss, s)
				continue
			}
			if v.isInvalid {
				ss = append(ss, modelStyle.Render(importantDayMark, grayTextStyle.Render(day), workState, grayTextStyle.Render(lunar)))
				continue
			}
			ss = append(ss, modelStyle.Render(importantDayMark, day, workState, lunar))
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
	s := borderStyle.Render(lipgloss.JoinVertical(lipgloss.Top, horizontalJoins...))

	return lipgloss.JoinHorizontal(lipgloss.Top, s, dayDetailStyle.Render(m.getDayDetail(selectDay.time))) + "\n\n" + helpMsg
}

func (m DayModel) getDayDetail(day time.Time) string {
	items := make([]string, 0, 5)
	date := day.Format("2006-01-02")
	items = append(items, fmt.Sprintf("阳历：%s（%s）", date, constants.Weekdays[day.Weekday().String()]))
	lunar := carbon.Parse(date).Lunar()
	items = append(items, fmt.Sprintf("阴历：【%s年】%s", lunar.Animal(), lunar.ToDateString()))

	holidays, ok := constants.Holidays[day.Year()]
	if ok {
		if holiday, ok := holidays[day.Format("01-02")]; ok {
			items = append(items, fmt.Sprintf("节日：%s", holiday))
		}
	}

	if s := m.getDays(day); s != "" {
		items = append(items, s)
	}

	return strings.Join(items, "\n")
}

func (m DayModel) getWorkAndLunar(t time.Time) (string, string) {
	workState, lunar := "\n", utils.GetLunarCalendar(t)
	if utils.IsWeekend(t) {
		workState = redTextStyle.Render("末")
	}
	_, ok := constants.WorkDays[t.Format("2006-01-02")]
	if ok {
		workState = redTextStyle.Render("班")
	}
	holidays, ok := constants.Holidays[t.Year()]
	if ok {
		if holiday, ok := holidays[t.Format("01-02")]; ok {
			workState = redTextStyle.Render("休")
			lunar = string(holiday)
		}
	}
	return workState, lunar
}

func (m DayModel) getDays(day time.Time) string {
	days := internal.Cfg.Days
	items := make([]string, 0, len(days))
	for _, birthday := range days {
		if s := m.genDayStr(day, birthday); s != "" {
			items = append(items, s)
		}
	}
	if len(items) == 0 {
		return ""
	}

	return fmt.Sprintf("\n🎉:\n	%s", strings.Join(items, "\n	"))
}

func (m DayModel) genDayStr(curDay time.Time, day internal.Day) string {
	t, err := time.Parse("01-02", day.Date)
	if err != nil {
		return ""
	}

	var lunarStr string
	month, d := curDay.Month(), curDay.Day()
	if day.IsLunar {
		lunarStr = "阴历"
		l := carbon.Parse(curDay.Format(time.DateOnly)).Lunar()
		month, d = time.Month(l.Month()), l.Day()
	}
	if month == t.Month() && d == t.Day() {
		return fmt.Sprintf("• %s（%s%s）", day.Name, lunarStr, day.Date)
	}
	return ""
}

func (m DayModel) getImportantDayMark(day time.Time) string {
	isImportDay := m.getDays(day) != ""
	if isImportDay {
		return "🌟"
	}
	return ""
}
