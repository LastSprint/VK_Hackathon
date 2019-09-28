package chat

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/websocket"
// )

// type Client struct {
// 	id   string
// 	name string
// }

// type Message struct {

// 	Recipient   string `json:"recipient"`
// 	Text        string `json:"text"`
// 	AuthorName  string `json:"authorName"`
// 	AuthorImage string `json:"authorImage"`
// 	MessageType int    `json:"messageType"`
// }

// var clients = new(map[string]interface{})

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func StartChat(router *mux.Router) {
// 	router.HandleFunc("/ws", chatHandler)
// }

// func chatHandler(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {

// 		log.Println("KEK PECK")

// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			delete(clients)
// 			log.Println(err)
// 			return
// 		}

// 		var msg Message

// 		json.Unmarshal(p, &msg)

// 		log.Println(msg)

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }
