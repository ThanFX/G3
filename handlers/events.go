package handlers

import (
	"net/http"

	"github.com/ThanFX/G3/models"
	"github.com/julienschmidt/httprouter"
)

func GetEventsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	events := models.GetEvents()
	SendJsonResponse(w, r, http.StatusOK, events, len(events), "")
}
