package handlers

import (
	"net/http"

	"github.com/ThanFX/G3/models"
	"github.com/julienschmidt/httprouter"
)

func PersonsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	persons := models.GetPersons()
	SendJsonResponse(w, r, http.StatusOK, persons, len(persons), "")
}
