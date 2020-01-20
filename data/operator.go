package data

import (
	"strconv"
	"time"
	"us-stock-trade-date/utils"
)

func GenHolidays() {
	csvName := "holidays.csv"
	row, err := utils.ReadCsvLastRow(csvName)
	var startYear int
	if err != nil || row == nil {
		startYear = 2010
	} else {
		startYear, _ = strconv.Atoi(row[0])
	}
	endYear := time.Now().Year()
	if startYear == endYear {
		return
	}
	for year := startYear; year <= endYear; year++ {
		rows := CrawlHolidays(year)
		utils.WriteCsv(csvName, rows)
	}
}

func GetHoliday(t time.Time) *Holiday {
	csvName := "holidays.csv"
	tf := t.Format("2006-01-02")
	row, err := utils.ReadCsvContains(csvName, 2, tf)
	if err != nil {
		return nil
	}
	return &Holiday{
		Name:   row[1],
		Date:   t,
		Status: row[3],
	}
}
