package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetEventsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	events := []string{"Тестовая строка №1", "Тестовая строка №2", "Тестовая строка №3"}
	SendJsonResponse(w, r, http.StatusOK, events, len(events), "")
}
