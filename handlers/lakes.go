package handlers

import (
	"net/http"

	"github.com/ThanFX/G3/models"
	"github.com/julienschmidt/httprouter"
)

func LakesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lakes := models.GetLakes()
	SendJsonResponse(w, r, http.StatusOK, lakes, len(lakes), "")
}
