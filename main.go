package rtp

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

/*
	TODO  Utils => Yardımcı fonksiyonlar yazılacak

	! 1
	* Request datalarını parse etme
	* Parse edilen datalardan gerekenleri çekme
	* Parse edilen datalar çekildikten sonra istenilen structa çevirme

	! 2
*/

type RTPCore struct {
	Server *mux.Router
	DB     *gorm.DB
	Hubs   map[string]Hub
}

type RTPCorePath struct {
	PATH string
}

func (rtp *RTPCore) Start(router *mux.Router, db *gorm.DB) {
	rtp.Server = router
	rtp.DB = db
	rtp.Hubs = make(map[string]Hub)
}
