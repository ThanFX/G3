package main

import (
	"database/sql"
	"fmt"
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
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func start() {
	rand.Seed(time.Now().UTC().UnixNano())
	models.CreatePerson(10)
	models.CreateLakes(3)
	models.LakesStart()
	models.PersonsStart()
	go models.EventLoop()
	models.SetDate(9842)
	models.SetCalendar()
	fmt.Println("Запускаем сервер...")
}

func getRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ws/events/", handlers.GetWSEventsHandler)
	router.GET("/api/persons", handlers.PersonsHandler)
	router.GET("/api/date", handlers.GetDateHandler)
	router.GET("/api/nextdate", handlers.NextDateHandler)
	router.GET("/api/lakes", handlers.LakesHandler)
	router.GET("/api/events", handlers.GetEventsHandler)
	router.GET("/", handlers.HomeHandler)
	router.ServeFiles("/public/*filepath", http.Dir("./public/"))
	return router
}

func main() {
	conf, err := config.Load()
	DB, err = sql.Open("sqlite3", "data/g3.db")
	if err != nil {
		stdlog.Printf("Ошибка открытия файла БД: %s", err)
	}
	models.DB = DB
	defer DB.Close()
	start()
	handlers.RunHub()
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
