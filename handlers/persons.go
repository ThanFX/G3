package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func PersonsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	SendJsonResponse(w, r, http.StatusOK, "", 0, "")
}
