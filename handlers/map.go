package handlers

import (
	"net/http"

	"github.com/ThanFX/G3/models"
	"github.com/julienschmidt/httprouter"
)

func MapHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	maps := models.GetMap()
	SendJsonResponse(w, r, http.StatusOK, maps, len(maps), "")
}
