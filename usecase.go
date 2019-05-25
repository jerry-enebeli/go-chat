package go_chat

import (
	"github.com/gorilla/websocket"

	"log"

	"time"
)

type chatUsecase interface {
	ReadData(conn *websocket.Conn)
	BroadcastMessage(store map[string]interface{})
}

type messageModel struct {
	UserID   string    `json:"user_id"`
	Group    bool      `json:"group"`
	Message  string    `json:"message"`
	Receiver string    `json:"receiver"`
	Room     string    `json:"room"`
	Time     time.Time `json:"time"`
}
type chatUC struct {
	ch chan messageModel
}

func newChatUseCase() chatUsecase {
	ch := make(chan messageModel)
	return &chatUC{ch: ch}
}

func (c *chatUC) ReadData(conn *websocket.Conn) {
	for {
		var message messageModel
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println(message.UserID)
			return
		}

		c.ch <- message
	}
}

func (c *chatUC) BroadcastMessage(store map[string]interface{}) {
	for message := range c.ch {
		switch message.Group {
		case true:
			sendGroupMessage(store,message)
		case false:
			sendSingleMessage(store, message)
		}
	}
}

func sendSingleMessage(store map[string]interface{}, message messageModel) {
	receiver, ok := store[message.Receiver]
	sender, _ := store[message.UserID]
	if !ok {
		v := sender.(*websocket.Conn)
		err:=v.WriteMessage(websocket.TextMessage, []byte("user is not online"))
		if err != nil {
			log.Print(err)
		}
	} else {
		v := receiver.(*websocket.Conn)
		err:=v.WriteMessage(websocket.TextMessage, []byte(message.Message))
		if err != nil {
			v := sender.(*websocket.Conn)
			err := v.WriteMessage(websocket.TextMessage, []byte("user is not online"))
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func sendGroupMessage(store map[string]interface{}, message messageModel){
	sampleUsersID := []string{"1234","5678", "9876"}

	for _, value := range sampleUsersID {
		message.UserID = value
		sendSingleMessage(store,message)
	}
}