package rtp

import "github.com/gorilla/mux"

type RTPCore struct {
	Server *mux.Router
}

type RTPCorePath struct {
	PATH string
}
