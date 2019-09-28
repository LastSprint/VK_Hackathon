package chat

import (
	"encoding/json"
	"suncity/commod"
	"suncity/notifications"
	"suncity/reps"

	"github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	rep *reps.ChatRep
}

func NewHub(rep *reps.ChatRep) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rep:        rep,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			logrus.Infoln("HUB UNREG")
			logrus.Infoln(client)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			var msg reps.MessageModel

			json.Unmarshal(message, &msg)

			client := h.findRecivier(msg.Sender, msg.Recipient, msg.Text)

			msg.IsMe = false

			if client == nil {
				go h.rep.SaveMessageById(&msg, msg.Recipient)
				continue
			}

			h.rep.SaveMessage(&msg, client.user)

			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
			}

			// for client := range h.clients {
			// 	logrus.Infoln("HUB BR")
			// 	logrus.Infoln(client)
			// 	select {
			// 	case client.send <- message:
			// 	default:
			// 		close(client.send)
			// 		delete(h.clients, client)
			// 	}
			// }
		}
	}
}

func (h *Hub) findRecivier(requester *commod.ServiceUser, reciverId string, data string) *Client {
	for client := range h.clients {
		if client.user.ID.Hex() == reciverId {
			return client
		}
	}

	go notifications.SendNotificationByUserId(requester, reciverId, data)
	return nil
}
