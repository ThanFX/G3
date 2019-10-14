package handlers

import (
	"net/http"

	"github.com/ThanFX/G3/models"
	"github.com/julienschmidt/httprouter"
)

func ChunkHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	chInfo := models.GetChunkTerrainsInfo(p.ByName("id"))
	SendJsonResponse(w, r, http.StatusOK, chInfo, 0, "")
}
