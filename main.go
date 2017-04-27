package main

import (
	stdlog "log"
	"net"

	"net/http"

	"github.com/ThanFX/G3/config"
	"github.com/ThanFX/G3/handlers"
	"github.com/ThanFX/G3/middlewares"
	"github.com/braintree/manners"
	"github.com/julienschmidt/httprouter"
)

func getRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/persons", handlers.PersonsHandler)
	router.GET("/", handlers.HomeHandler)
	router.ServeFiles("/public/*filepath", http.Dir("./public/"))
	return router
}

func main() {
	conf, err := config.Load()
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
