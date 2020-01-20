package handlers

import (
	"log"
	"net/http"
	"us-stock-trade-date/result"
	"us-stock-trade-date/result/code"
)

func HandlePanic(w http.ResponseWriter, _ *http.Request, err interface{}) {
	result.WriteJSON(w, result.Response{
		Code:    code.SystemError,
		Message: "System error",
	}.Result())
	log.Print(err)
}
