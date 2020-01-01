package main

import (
	"database/sql"
	"fmt"
	stdlog "log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/ThanFX/G3/areas"
	"github.com/ThanFX/G3/config"
	"github.com/ThanFX/G3/handlers"
	"github.com/ThanFX/G3/libs"
	"github.com/ThanFX/G3/middlewares"
	"github.com/ThanFX/G3/models"
	"github.com/braintree/manners"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func start() {
	rand.Seed(time.Now().UTC().UnixNano())
	libs.ReadMasteryItemsCatalog(db)
	libs.ReadMastershipsCatalog(db)
	models.ReadDate()
	models.MapInitialize()
	models.CreateTerrains()
	areas.AreasStart()
	models.ReadPersonsCatalog()
	models.PersonsStart()
	//go models.EventLoop()
	fmt.Println("Запускаем сервер...")
}

func getRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ws/events/", handlers.GetWSEventsHandler)
	router.GET("/api/map", handlers.MapHandler)
	router.GET("/api/chunk/:id", handlers.ChunkHandler)
	router.GET("/api/persons", handlers.PersonsHandler)
	router.GET("/api/person/:id/inventory", handlers.PersonInventoryHandler)
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
	db, err = sql.Open("postgres", "postgres://cdwjrqix:m-zQFIhz8jCYgY4XmyvV9Z40h1ShZGSR@balarama.db.elephantsql.com:5432/cdwjrqix")
	if err != nil {
		stdlog.Printf("Ошибка открытия файла БД: %s", err)
	}
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)
	db.SetConnMaxLifetime(0)
	models.DB = db
	defer db.Close()
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
