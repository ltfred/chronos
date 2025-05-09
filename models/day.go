package models

import (
	"fmt"
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

type DayModel struct {
	choices     []dayChoice
	month, year int
	cursor      int
}

func NewDayModel(year, month int) DayModel {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	lastDay := time.Date(year, time.Month(month)+1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1)
	now := time.Now()

	week := int(firstDay.Weekday())
	if week == 0 {
		week = 7
	}
	dayModel := DayModel{month: month, year: year}

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

	return dayModel
}

func (model DayModel) Init() tea.Cmd {
	return nil
}

func (model DayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return model, tea.Quit
		case "down", "j":
			if model.cursor+7 >= len(model.choices) {
				return model, nil
			}
			model.cursor += 7
			return model, nil
		case "up", "k":
			if model.cursor-7 < 0 {
				return model, nil
			}
			model.cursor -= 7
			return model, nil
		case "left", "h":
			if model.cursor-1 < 0 {
				return model, nil
			}
			model.cursor -= 1
			return model, nil
		case "right", "l":
			if model.cursor+1 >= len(model.choices) {
				return model, nil
			}
			model.cursor += 1
			return model, nil
		case "m":
			t := model.choices[model.cursor].time
			return NewMonthModel(t.Year(), t.Month()), nil
		case "n":
			if model.month+1 > 12 {
				return model, nil
			}
			return NewDayModel(model.year, model.month+1), nil
		case "p":
			if model.month-1 < 1 {
				return model, nil
			}
			return NewDayModel(model.year, model.month-1), nil
		default:
		}
	}

	return model, nil
}

func (model DayModel) View() string {
	rowCount := len(model.choices) / 7
	rows := make([][]dayChoice, 0, rowCount)
	for i := 0; i < rowCount; i++ {
		rows = append(rows, model.choices[i*7:(i+1)*7])
	}
	var selectDay dayChoice
	joinHorizontal := func(choices []dayChoice) string {
		ss := make([]string, 0, 7)
		for _, v := range choices {
			day := utils.OnlyDayFormat(v.time)
			workState, lunar := model.getWorkAndLunar(v.time)
			importantDayMark := model.getImportantDayMark(v.time)
			if v.pos == model.cursor {
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
			ss = append(ss, modelStyle.Render(importantDayMark, boldTextStyle.Render(day), workState, lunar))
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

	return lipgloss.JoinHorizontal(lipgloss.Top, s, dayDetailStyle.Render(model.getDayDetail(selectDay.time)))
}

func (model DayModel) getDayDetail(day time.Time) string {
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

	if s := model.getBirthDayDay(day); s != "" {
		items = append(items, s)
	}
	if s := model.getMemorialDay(day); s != "" {
		items = append(items, s)
	}

	return strings.Join(items, "\n")
}

func (model DayModel) getWorkAndLunar(t time.Time) (string, string) {
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

func (model DayModel) getBirthDayDay(day time.Time) string {
	birthdays := internal.Cfg.ImportantDay.Birthdays
	items := make([]string, 0, len(birthdays))
	for _, birthday := range birthdays {
		if s := model.genImportantDayStr(day, birthday); s != "" {
			items = append(items, s)
		}
	}
	if len(items) == 0 {
		return ""
	}

	return fmt.Sprintf("\n🎂:\n	%s", strings.Join(items, "\n	"))
}

func (model DayModel) getMemorialDay(day time.Time) string {
	memorialDays := internal.Cfg.ImportantDay.MemorialDays
	items := make([]string, 0, len(memorialDays))
	for _, memorialDay := range memorialDays {
		if s := model.genImportantDayStr(day, memorialDay); s != "" {
			items = append(items, s)
		}
	}
	if len(items) == 0 {
		return ""
	}

	return fmt.Sprintf("\n🎉：\n	%s", strings.Join(items, "\n	"))
}

func (model DayModel) genImportantDayStr(curDay time.Time, day internal.Day) string {
	t, err := time.Parse("01-02", day.Date)
	if err != nil {
		return ""
	}

	var lunarStr string
	m, d := curDay.Month(), curDay.Day()
	if day.IsLunar {
		lunarStr = "阴历"
		l := carbon.Parse(curDay.Format(time.DateOnly)).Lunar()
		m, d = time.Month(l.Month()), l.Day()
	}
	if m == t.Month() && d == t.Day() {
		return fmt.Sprintf("• %s（%s%s）", day.Name, lunarStr, day.Date)
	}
	return ""
}

func (model DayModel) getImportantDayMark(day time.Time) string {
	isImportDay := model.getBirthDayDay(day) != "" || model.getMemorialDay(day) != ""
	if isImportDay {
		return "🌟"
	}
	return ""
}
