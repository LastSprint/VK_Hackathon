package chat

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"suncity/commod"

// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/websocket"
// )

// type Client struct {
// 	user *commod.ServiceUser
// 	name string
// 	con  *websocket.Conn
// }

// func (cl *Client) Send(msg *Message) {
// 	cl.con.WriteJSON(msg)
// }

// type Message struct {
// 	Recipient   string `json:"recipient"`
// 	Text        string `json:"text"`
// 	AuthorName  string `json:"authorName"`
// 	AuthorImage string `json:"authorImage"`
// 	MessageType int    `json:"messageType"`
// }

// type ConnectMsg struct {
// 	Token string
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

// 		cl := Client{}

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

// package main

// import (
// 	"flag"
// 	"log"
// 	"net/http"
// )

// var addr = flag.String("addr", ":8080", "http service address")

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL)
// 	if r.URL.Path != "/" {
// 		http.Error(w, "Not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	http.ServeFile(w, r, "home.html")
// }

// func main() {
// 	flag.Parse()
// 	hub := newHub()
// 	go hub.run()
// 	http.HandleFunc("/", serveHome)
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		serveWs(hub, w, r)
// 	})
// 	err := http.ListenAndServe(*addr, nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
