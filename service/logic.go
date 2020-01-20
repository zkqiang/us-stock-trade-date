package service

import (
	"time"
	"us-stock-trade-date/data"
)

func IsTradingDay(t time.Time) (bool, string) {
	if b, r := isWeekend(t); b {
		return false, r
	}
	if b, r := isHoliday(t); b {
		return false, r
	}
	return true, ""
}

func GetTradingHours(t time.Time) (*time.Time, *time.Time) {
	if b, _ := isWeekend(t); b {
		return nil, nil
	}
	open := time.Date(
		t.Year(), t.Month(), t.Day(), 9, 30, 0, 0, t.Location())

	var close_ time.Time

	d := data.GetHoliday(t)
	if d != nil {
		if d.Status == "Closed" {
			return nil, nil
		} else {
			close_, _ = time.ParseInLocation("3:04 PM", d.Status, t.Location())
		}
	} else {
		close_ = time.Date(
			t.Year(), t.Month(), t.Day(), 16, 0, 0, 0, t.Location())
	}
	return &open, &close_
}

func isWeekend(t time.Time) (bool, string) {
	wd := t.Weekday()
	if wd == time.Saturday {
		return true, "Saturday"
	} else if wd == time.Sunday {
		return true, "Sunday"
	}
	return false, ""
}

func isHoliday(t time.Time) (bool, string) {
	d := data.GetHoliday(t)
	if d != nil && d.Status == "Closed" {
		return true, d.Name
	}
	return false, ""
}
