package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ChunkHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//inv := models.GetPersonInventory(p.ByName("id"))
	SendJsonResponse(w, r, http.StatusOK, "inv", 0, "")
}
