package go_chat

import (
	"chatapp/go-chat"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	chatUseCase := go_chat.newChatUseCase()

	handlers := go_chat.chatDelivery{useCase: chatUseCase}
	router.HandleFunc("/chat/{userID}", handlers.chatHandler)

	http.ListenAndServe(":8086", router)
}

