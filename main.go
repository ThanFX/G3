package main

import (
	stdlog "log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/ThanFX/G3/config"
	"github.com/ThanFX/G3/handlers"
	"github.com/ThanFX/G3/middlewares"
	"github.com/ThanFX/G3/models"
	"github.com/braintree/manners"
	"github.com/julienschmidt/httprouter"
)

func start() {
	models.CreatePerson(20)
	models.SetDate(9842)
	models.SetCalendar()
}

func getRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/persons", handlers.PersonsHandler)
	router.GET("/api/date", handlers.GetDateHandler)
	router.GET("/api/nextdate", handlers.NextDateHandler)
	router.GET("/", handlers.HomeHandler)
	router.ServeFiles("/public/*filepath", http.Dir("./public/"))
	return router
}

func main() {
	conf, err := config.Load()
	start()
	rand.Seed(time.Now().UTC().UnixNano())
	if err != nil {
		stdlog.Fatal(err.Error())
	}

	httpAddr := net.JoinHostPort(conf.Api.Host, conf.Api.Port)
	router := middlewares.RecoverMiddleware(getRouter())
	httpServer := manners.NewServer()
	httpServer.Addr = httpAddr
	httpServer.Handler = router
	httpServer.ReadTimeout = conf.Api.ReadTimeout
	stdlog.Fatal(httpServer.ListenAndServe())
}
