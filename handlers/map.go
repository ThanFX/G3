package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func MapHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//inv := models.GetPersonInventory(p.ByName("id"))
	SendJsonResponse(w, r, http.StatusOK, "inv", 0, "")
}
