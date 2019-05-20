package handlers

import (
	"net/http"

	"github.com/ThanFX/G3/models"
	"github.com/julienschmidt/httprouter"
)

func DateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	curDate := models.GetCalendarDate()
	SendJsonResponse(w, r, http.StatusOK, curDate, 0, "")
}
