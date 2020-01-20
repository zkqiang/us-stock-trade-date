package router

import (
	"net/http"
	"us-stock-trade-date/service/handlers"

	"github.com/julienschmidt/httprouter"
)

func New() *httprouter.Router {
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(handlers.HandleCors)

	router.PanicHandler = handlers.HandlePanic

	router.GET("/date", handlers.HandleDate)

	return router
}
