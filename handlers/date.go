package handlers

import (
	"net/http"

	"github.com/ThanFX/G3/models"
	"github.com/julienschmidt/httprouter"
)

func GetDateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	curDate := models.GetCalendarDate()
	SendJsonResponse(w, r, http.StatusOK, curDate, 0, "")
}

func NextDateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	models.IncDate()
	models.LakesNextDate()
	models.PersonsNextDate()
	SendJsonResponse(w, r, http.StatusOK, nil, 0, "")
}
