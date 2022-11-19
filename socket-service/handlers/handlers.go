package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
var wsConn *websocket.Conn

type Ihandler interface {
	WScable(http.ResponseWriter, *http.Request)
}

type handler struct{}

func NewHandler() Ihandler {
	return &handler{}
}

type Data struct {
	Message string `json:"message"`
}

func (h *handler) WScable(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	fmt.Printf("Web socket open and waiting for messages")
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}
	defer wsConn.Close()
	for {
		var data Data
		err := wsConn.ReadJSON(&data)
		if err != nil {
			fmt.Printf("err: %v\n", "did not get a json string with correct structure... probably")
			return
		}
	}
}

func (h *handler) SendMessage(ws *websocket.Conn, msg string) {
	err := ws.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("err: %v\n", "could not send message with ws")
		return
	}
}
