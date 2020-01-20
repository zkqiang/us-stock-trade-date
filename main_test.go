package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"us-stock-trade-date/data"
	"us-stock-trade-date/result"
	"us-stock-trade-date/result/code"
	"us-stock-trade-date/router"

	"github.com/goinggo/mapstructure"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	router_ := router.New()
	test := testRequest(t, router_)

	var resp result.Response
	var err error

	// 测试无参数
	resp = test("GET", "/date")
	assert.Equal(t, resp.Code, 0)
	var tradeData1 data.TradeData
	err = mapstructure.Decode(resp.Data, &tradeData1)
	assert.Nil(t, err)

	// 测试交易日参数
	resp = test("GET", "/date?ts=1575559447000")
	assert.Equal(t, resp.Code, 0)
	var tradeData2 data.TradeData
	err = mapstructure.Decode(resp.Data, &tradeData2)
	assert.Nil(t, err)
	assert.Equal(t, tradeData2.DateNYC, "2019-12-05 10:24:07 -0500 EST")
	assert.Equal(t, tradeData2.IsTradingDay, true)
	assert.Equal(t, tradeData2.IsTradingHours, true)
	assert.Equal(t, tradeData2.OpenNYC, "09:30")
	assert.Equal(t, tradeData2.CloseNYC, "16:00")
	assert.Equal(t, tradeData2.NextTradingDay, "2019-12-06")
	assert.Empty(t, tradeData2.DateLoc)
	assert.Empty(t, tradeData2.OpenLoc)
	assert.Empty(t, tradeData2.CloseLoc)

	// 测试半交易日参数
	resp = test("GET", "/date?ts=1530609847000")
	assert.Equal(t, resp.Code, 0)
	var tradeData3 data.TradeData
	err = mapstructure.Decode(resp.Data, &tradeData3)
	assert.Nil(t, err)
	assert.Equal(t, tradeData3.DateNYC, "2018-07-03 05:24:07 -0400 EDT")
	assert.Equal(t, tradeData3.IsTradingDay, true)
	assert.Equal(t, tradeData3.IsTradingHours, false)
	assert.Equal(t, tradeData3.OpenNYC, "09:30")
	assert.Equal(t, tradeData3.CloseNYC, "13:00")
	assert.Equal(t, tradeData3.NextTradingDay, "2018-07-05")
	assert.Empty(t, tradeData3.DateLoc)
	assert.Empty(t, tradeData3.OpenLoc)
	assert.Empty(t, tradeData3.CloseLoc)

	// 测试非交易日参数
	resp = test("GET", "/date?ts=1562253847000")
	assert.Equal(t, resp.Code, 0)
	var tradeData4 data.TradeData
	err = mapstructure.Decode(resp.Data, &tradeData4)
	assert.Nil(t, err)
	assert.Equal(t, tradeData4.DateNYC, "2019-07-04 11:24:07 -0400 EDT")
	assert.Equal(t, tradeData4.IsTradingDay, false)
	assert.Equal(t, tradeData4.IsTradingHours, false)
	assert.Empty(t, tradeData4.OpenNYC)
	assert.Empty(t, tradeData4.CloseNYC)
	assert.Equal(t, tradeData4.NextTradingDay, "2019-07-05")
	assert.Empty(t, tradeData4.DateLoc)
	assert.Empty(t, tradeData4.OpenLoc)
	assert.Empty(t, tradeData4.CloseLoc)
	assert.Equal(t, tradeData4.Reason, "Independence Day")

	// 测试本地时间参数
	resp = test("GET", "/date?ts=1591644998000&loc=PRC")
	assert.Equal(t, resp.Code, 0)
	var tradeData5 data.TradeData
	err = mapstructure.Decode(resp.Data, &tradeData5)
	assert.Nil(t, err)
	assert.Equal(t, tradeData5.DateNYC, "2020-06-08 15:36:38 -0400 EDT")
	assert.Equal(t, tradeData5.IsTradingDay, true)
	assert.Equal(t, tradeData5.IsTradingHours, true)
	assert.Equal(t, tradeData5.OpenNYC, "09:30")
	assert.Equal(t, tradeData5.CloseNYC, "16:00")
	assert.Equal(t, tradeData5.NextTradingDay, "2020-06-09")
	assert.Equal(t, tradeData5.DateLoc, "2020-06-09 03:36:38 +0800 CST")
	assert.Equal(t, tradeData5.OpenLoc, "21:30")
	assert.Equal(t, tradeData5.CloseLoc, "04:00")
	assert.Empty(t, tradeData5.Reason)

	// 测试过早的年份报错
	resp = test("GET", "/date?ts=1246721047000")
	assert.Equal(t, resp.Code, code.YearOutOfRange)
	assert.Nil(t, resp.Data)

	// 测试未来的年份报错
	nextYear := time.Now().AddDate(1, 0, 0)
	resp = test("GET", "/date?ts="+strconv.FormatInt(nextYear.UnixNano(), 10)[:13])
	assert.Equal(t, resp.Code, code.YearOutOfRange)
	assert.Nil(t, resp.Data)
}

func testRequest(t *testing.T, rt *httprouter.Router) func(method, url string) result.Response {
	return func(method, url string) result.Response {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(method, url, nil)
		rt.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Fatalf("Invalid status: %s", rr.Body)
		}

		var res result.Response
		err := json.Unmarshal(rr.Body.Bytes(), &res)
		if err != nil {
			t.Fatalf("Invalid body: %s", rr.Body)
		}
		return res
	}
}
