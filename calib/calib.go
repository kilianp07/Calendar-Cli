package calib

import (
	"encoding/json"
	"strconv"
	"time"
)

// Calendar represents a calendar for a specific month and year.
type Calendar struct {
	Month    time.Month   `json:"month"`
	Year     int          `json:"year"`
	Days     [][]int      `json:"days"`
	Weekdays []string     `json:"weekdays"`
	NumWeeks int          `json:"num_weeks"`
	FirstDay time.Weekday `json:"first_day"`
	LastDay  int          `json:"last_day"`
}

// NewCalendar generates a calendar for the specified month and year.
func NewCalendar(month int, year int) *Calendar {
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	firstDayOfWeek := (int(firstDay.Weekday()) + 6) % 7
	lastDay := firstDay.AddDate(0, 1, -1)

	numWeeks := (int(lastDay.Day()) + firstDayOfWeek + 6) / 7
	weeks := make([][]int, numWeeks)
	currentDay := 1

	for i := 0; i < numWeeks; i++ {
		weeks[i] = make([]int, 7)
		for j := 0; j < 7; j++ {
			if i == 0 && j < firstDayOfWeek {
				weeks[i][j] = 0 // Padding for the first week
			} else if currentDay <= int(lastDay.Day()) {
				weeks[i][j] = currentDay
				currentDay++
			}
		}
	}

	return &Calendar{
		Month:    firstDay.Month(),
		Year:     year,
		Days:     weeks,
		Weekdays: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		NumWeeks: numWeeks,
		FirstDay: firstDay.Weekday(),
		LastDay:  int(lastDay.Day()),
	}
}

// ToJSON returns the calendar data as JSON.
func (c *Calendar) ToJSON() (string, error) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (c *Calendar) ToHtml() (string, error) {
	html := "<table><thead><tr>"
	for _, weekday := range c.Weekdays {
		html += "<th>" + weekday + "</th>"
	}
	html += "</tr></thead><tbody>"
	for _, week := range c.Days {
		html += "<tr>"
		for _, day := range week {
			if day == 0 {
				html += "<td></td>"
			} else {
				html += "<td>" + strconv.Itoa(day) + "</td>"
			}
		}
		html += "</tr>"
	}
	html += "</tbody></table>"
	return html, nil
}
