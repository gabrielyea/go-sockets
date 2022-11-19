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
	Home(http.ResponseWriter, *http.Request)
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
	fmt.Printf("Web socket open and waiting for messages \n")
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
			res := fmt.Sprintf("connection was closed... probably %s", err.Error())
			fmt.Printf("error: %v\n", res)
			h.SendMessage(wsConn, "something went wrong!")
			return
		}
		h.SendMessage(wsConn, data.Message)
	}
}

func (h *handler) SendMessage(ws *websocket.Conn, msg string) {
	res := fmt.Sprintf("message from GO: %s \n", msg)
	err := ws.WriteMessage(websocket.TextMessage, []byte(res))
	if err != nil {
		fmt.Printf("err: %v\n", "could not send message with ws")
		return
	}
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from home")
}
