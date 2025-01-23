package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLunarCalendar(t *testing.T) {
	calendar := GetLunarCalendar(time.Date(2025, time.Month(1), 29, 0, 0, 0, 0, time.Local))
	assert.Equal(t, "春节", calendar)

	calendar = GetLunarCalendar(time.Date(2025, time.Month(1), 8, 0, 0, 0, 0, time.Local))
	assert.Equal(t, "初九", calendar)
}
