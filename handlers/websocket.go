package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type User struct {
	hub *Hub
	// The websocket connection.
	ws *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

type Hub struct {
	// Пользователи
	users map[*User]bool
	// Исходящие сообщения для всех активных клиентов
	broadcast chan []byte
	// Запросы на добавление новых активных клиентов
	register chan *User
	// Запросы на удаление клиентов
	unregister chan *User
}

func (h *Hub) run() {
	for {
		select {
		case user := <-h.register:
			h.users[user] = true
			users++
		case user := <-h.unregister:
			if _, ok := h.users[user]; ok {
				delete(h.users, user)
				close(user.send)
				users--
			}
		case message := <-h.broadcast:
			for user := range h.users {
				select {
				case user.send <- message:
				default:
					close(user.send)
					delete(h.users, user)
				}
			}
		}
	}
}

var (
	hub     = newHub()
	newline = []byte{'\n'}
	space   = []byte{' '}
	users   int
)

func RunHub() {
	go hub.run()
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *User),
		unregister: make(chan *User),
		users:      make(map[*User]bool),
	}
}

func (u *User) readPump() {
	defer func() {
		u.hub.unregister <- u
		u.ws.Close()
	}()
	u.ws.SetReadLimit(maxMessageSize)
	u.ws.SetReadDeadline(time.Now().Add(pongWait))
	u.ws.SetPongHandler(func(string) error {
		u.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := u.ws.ReadMessage()
		if err != nil {
			fmt.Printf("%s", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Fatal(err.Error())
			}
			break
		}
		fmt.Print(string(message))
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		//u.hub.broadcast <- message
	}
}

func (u *User) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		u.ws.Close()
	}()
	for {
		select {
		case message, ok := <-u.send:
			u.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				u.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := u.ws.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			u.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
