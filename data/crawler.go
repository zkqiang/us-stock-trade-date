package data

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

func CrawlHolidays(year int) [][]string {
	var res [][]string

	url := "https://www.tradinghours.com/exchanges/nyse/market-holidays/" + strconv.Itoa(year)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		log.Fatal("Crawl holidays error: ", err)
		return res
	}

	trs, _ := htmlquery.QueryAll(doc, "//table/tbody/tr")

	for _, tr := range trs {
		// 解析年份
		yearNode, _ := htmlquery.Query(tr, "./td[1]/text()")
		year := strings.TrimSpace(htmlquery.InnerText(yearNode))

		// 解析节日名
		holidayNode, _ := htmlquery.Query(tr, "./td[2]//text()")
		holiday := strings.TrimSpace(htmlquery.InnerText(holidayNode))
		holiday = strings.Trim(holiday, "\"")

		// 解析日期
		dateNode, _ := htmlquery.Query(tr, "./td[3]/text()")
		date := strings.TrimSpace(htmlquery.InnerText(dateNode))
		t, _ := time.Parse("January 2, 2006", date)
		date = t.Format("2006-01-02")

		// 解析开市状态  Closed - 休市 || X:00 PM - 下午X点提前收盘
		statusNode, _ := htmlquery.Query(tr, "./td[4]/text()")
		status := strings.TrimSpace(htmlquery.InnerText(statusNode))
		status = strings.ReplaceAll(status, "\n", "")
		status = strings.ReplaceAll(status, "‡", "")
		re, _ := regexp.Compile("\\d:\\d{2} PM")
		if re.MatchString(status) {
			status = re.FindString(status)
		}

		res = append(res, []string{year, holiday, date, status})
	}
	return res
}
