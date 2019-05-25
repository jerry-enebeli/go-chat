package go_chat

import (
	"chatapp/go-chat"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type chatDelivery struct {
	useCase go_chat.chatUsecase
}

var store map[string]interface{}

func init() {

	store = make(map[string]interface{})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (c *chatDelivery) chatHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	conn, err := upgrader.Upgrade(w, r, nil)

	println(conn.RemoteAddr().String())
	if err != nil {
		log.Println(err.Error())
		return
	}

	go c.useCase.ReadData(conn)
	go c.useCase.BroadcastMessage(store)

	//saves user's connection to global store
	store[userID] = conn

}
