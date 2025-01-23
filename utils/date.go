package utils

import (
	"github.com/dromara/carbon/v2"
	"time"
)

// GetLunarCalendar 获取农历日期 如果是农历节日则返回节日
func GetLunarCalendar(t time.Time) string {
	l := carbon.Parse(t.Format("2006-01-02")).Lunar()
	lunar := l.ToDayString()
	if l.Festival() != "" {
		lunar = l.Festival()
	}
	return lunar
}
