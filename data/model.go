package data

import "time"

type Holiday struct {
	Name   string
	Date   time.Time
	Status string
}

type TradeData struct {
	DateNYC        string `json:"dateNYC"`
	IsTradingHours bool   `json:"isTradingHours"`
	IsTradingDay   bool   `json:"isTradingDay"`
	OpenNYC        string `json:"openNYC,omitempty"`
	CloseNYC       string `json:"closeNYC,omitempty"`
	Reason         string `json:"reason,omitempty"`
	DateLoc        string `json:"dateLoc,omitempty"`
	OpenLoc        string `json:"openLoc,omitempty"`
	CloseLoc       string `json:"closeLoc,omitempty"`
	NextTradingDay string `json:"nextTradingDay"`
}
