package handlers

import (
	"net/http"
	"strconv"
	"time"
	"us-stock-trade-date/data"
	"us-stock-trade-date/result"
	"us-stock-trade-date/result/code"
	"us-stock-trade-date/service"
	"us-stock-trade-date/utils"

	"github.com/julienschmidt/httprouter"
)

func HandleDate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loc, _ := time.LoadLocation("America/New_York")

	var locReq *time.Location
	var err error
	locStr := r.FormValue("loc")
	if locStr != "" {
		locReq, err = time.LoadLocation(locStr)
		if err != nil {
			result.WriteJSON(w, result.Response{
				Code:    code.ParamInvalid,
				Message: "loc param invalid",
			}.Result())
			return
		}
	}

	tsStr := r.FormValue("ts")
	now := time.Now().In(loc)
	var t time.Time
	if tsStr == "" {
		t = now
	} else {
		ts, err := strconv.ParseInt(tsStr, 10, 64)
		if err != nil {
			result.WriteJSON(w, result.Response{
				Code:    code.ParamInvalid,
				Message: "ts param invalid",
			}.Result())
			return
		}
		if len(tsStr) == 13 {
			t = utils.TimeFromTimestamp(ts).In(loc)
		} else if len(tsStr) == 10 {
			t = utils.TimeFromUnix(ts).In(loc)
		} else {
			result.WriteJSON(w, result.Response{
				Code:    code.ParamInvalid,
				Message: "ts param invalid",
			}.Result())
			return
		}
		if t.Year() < 2010 || t.Year() > now.Year() {
			result.WriteJSON(w, result.Response{
				Code:    code.YearOutOfRange,
				Message: "year out of range: 2010 ~ " + strconv.Itoa(now.Year()),
			}.Result())
			return
		}
	}

	isTradingDay, reason := service.IsTradingDay(t)

	open, close_ := service.GetTradingHours(t)

	isTradingHours := false
	if isTradingDay {
		if open != nil && close_ != nil {
			if t.After(*open) && t.Before(*close_) {
				isTradingHours = true
			}
		}
	}

	var openNYC string
	var closeNYC string
	if open != nil {
		openNYC = open.Format("15:04")
	}
	if close_ != nil {
		closeNYC = close_.Format("15:04")
	}

	nextTradingDay := t
	for {
		nextTradingDay = nextTradingDay.AddDate(0, 0, 1)
		if b, _ := service.IsTradingDay(nextTradingDay); b {
			break
		}
	}

	resData := data.TradeData{
		DateNYC:        t.Format("2006-01-02 15:04:05 -0700 MST"),
		IsTradingHours: isTradingHours,
		IsTradingDay:   isTradingDay,
		OpenNYC:        openNYC,
		CloseNYC:       closeNYC,
		Reason:         reason,
		NextTradingDay: nextTradingDay.Format("2006-01-02"),
	}

	if locReq != nil {
		var openLoc string
		var closeLoc string
		if open != nil {
			openLoc = open.In(locReq).Format("15:04")
		}
		if close_ != nil {
			closeLoc = close_.In(locReq).Format("15:04")
		}

		resData.DateLoc = t.In(locReq).Format("2006-01-02 15:04:05 -0700 MST")
		resData.OpenLoc = openLoc
		resData.CloseLoc = closeLoc
	}

	result.WriteJSON(w, result.Response{
		Data:    resData,
		Message: "Success",
	}.Result())
}
