package rtp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// handlerPath için interface örneği
type HandlerFunctionInterface interface {
	POST(RTPRequest, *gorm.DB) RTPResponse
	GET(RTPRequest, *gorm.DB) RTPMultiGetResponse
	GETONE(RTPRequest, *gorm.DB) RTPResponse
	DELETE(RTPRequest, *gorm.DB) RTPResponse
	PUT(RTPRequest, *gorm.DB) RTPResponse
	STAT(RTPRequest, *gorm.DB) RTPResponse
}

type HandlerFunc func(RTPRequest)

type HandlerPathFunc func(HandlerFunctionInterface)

// RouteWrapper
func (rtp *RTPCore) HandlePath(path string, handleFunc HandlerFunctionInterface) {
	// testRequest := new(RTPRequest)
	hub := newHub(rtp.DB, handleFunc)
	rtp.Hubs[path] = hub
	go hub.run()

	rtp.Server.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		Serve(hub, w, r, rtp.DB, handleFunc)
		fmt.Println("---\tRequest is accepted")
	})
}

func Serve(hub *Hub, w http.ResponseWriter, r *http.Request, db *gorm.DB, handler HandlerFunctionInterface) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, Conn: conn, DB: db, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
