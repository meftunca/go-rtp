package rtp

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
)

type ConnectionStatus struct {
	Connected      bool    `json:"connected"`
	Message        *string `json:"message"`
	TotalSubscribe int     `json:"totalCount"`
	DeviceToken    string  `json:"deviceToken"`
}

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
	HandlerFunction HandlerFunctionInterface
	unregister      chan *Client
	DB              *gorm.DB
}

func newHub(db *gorm.DB, handler HandlerFunctionInterface) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),

		HandlerFunction: handler,
		DB:              db,
	}
}

func (h *Hub) run() {
	for {
		fmt.Println("-----------\t Hub is started")
		select {
		case client := <-h.register:
			client.DeviceToken = GenerateUniqueID(8)
			h.clients[client] = true
			client.Conn.WriteJSON(ConnectionStatus{Connected: true, TotalSubscribe: len(h.clients), DeviceToken: client.DeviceToken})
		case client := <-h.unregister:
			client.Conn.WriteJSON(ConnectionStatus{Connected: false, TotalSubscribe: len(h.clients), DeviceToken: client.DeviceToken})
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Println("message", string(message))
			req, err := h.ByteToRTPRequest(message)
			byteSend := make([]byte, 0)
			if req.METHOD == "GET" {
				fmt.Println("----\tmessage", req.METHOD, err)
				resp := h.HandlerFunction.GET(req, h.DB)
				byt, _ := json.Marshal(resp)
				byteSend = byt

			} else {
				fmt.Println("----\tmessage", req.METHOD, err)
				resp := h.SwitchHandlerMethod(req, h.HandlerFunction)
				byt, _ := json.Marshal(resp)
				byteSend = byt
			}
			for client := range h.clients {
				//
				if req.METHOD == "GETONE" {
					if client.DeviceToken == req.DeviceToken {
						client.Conn.WriteMessage(1, byteSend)
						// select {
						// case client.send <- byteSend:
						// default:
						// 	close(client.send)
						// 	delete(h.clients, client)
						// }
					}
				} else {
					client.Conn.WriteMessage(1, byteSend)

					// select {
					// case client.send <- byteSend:
					// default:
					// 	close(client.send)
					// 	delete(h.clients, client)
					// }
				}
			}
		}
	}
}

func (h *Hub) SwitchHandlerMethod(request RTPRequest, handler HandlerFunctionInterface) RTPResponse {
	newResponse := &RTPResponse{}
	switch request.METHOD {
	case "POST":
		*newResponse = handler.POST(request, h.DB)
	case "GET":
		*newResponse = handler.GETONE(request, h.DB)
	case "GETONE":
		*newResponse = handler.GETONE(request, h.DB)
	case "STAT":
		*newResponse = handler.STAT(request, h.DB)
	case "PUT":
		*newResponse = handler.PUT(request, h.DB)
	case "DELETE":
		*newResponse = handler.DELETE(request, h.DB)

	}
	return *newResponse
}
