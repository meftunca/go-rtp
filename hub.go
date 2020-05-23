package rtp

import (
	"fmt"

	"github.com/jinzhu/gorm"
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
	HandlerFunction HandlerFunctionInterface
	unregister      chan *Client
	DB              *gorm.DB
}

func newHub(db *gorm.DB, handler HandlerFunctionInterface) *Hub {
	return &Hub{
		broadcast:       make(chan []byte),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		clients:         make(map[*Client]bool),
		HandlerFunction: handler,
		DB:              db,
	}
}

func (h *Hub) run() {
	for {
		fmt.Println("\n\n-----------\t Hub is started\n\n")
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			req, err := h.ByteToRTPRequest(message)
			fmt.Println("\n\n----\tmessage", req.METHOD, err)
			resp := SwitchHandlerMethod(req, h.HandlerFunction)
			for client := range h.clients {
				client.Conn.WriteJSON(resp)
				// select {
				// case client.send <- message:
				// default:
				// 	close(client.send)
				// 	delete(h.clients, client)
				// }
			}
		}
	}
}

func SwitchHandlerMethod(request RTPRequest, handler HandlerFunctionInterface) RTPResponse {
	newResponse := &RTPResponse{}
	switch request.METHOD {
	case "POST":
		*newResponse = handler.POST(request)
	case "GET":
		*newResponse = handler.GET(request)
	case "PUT":
		*newResponse = handler.PUT(request)
	case "DELETE":
		*newResponse = handler.DELETE(request)

	}
	return *newResponse
}
