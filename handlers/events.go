package handlers

import (
	"net/http"

	"github.com/ThanFX/G3/models"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

func GetEventsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	events := models.GetEvents()
	SendJsonResponse(w, r, http.StatusOK, events, len(events), "")
}

func GetWSEventsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	user := &User{hub: hub, ws: ws, send: make(chan []byte)}
	user.hub.register <- user
	go user.writePump()
	user.readPump()
}
