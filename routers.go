package rtp

import (
	"net/http"
)

// handlerPath için interface örneği
type HandlerFunctionInterface interface {
	POST(RTPRequest) RTPResponse
	GET(RTPRequest) RTPResponse
	DELETE(RTPRequest) RTPResponse
	PUT(RTPRequest) RTPResponse
}

type HandlerFunc func(RTPRequest)

type HandlerPathFunc func(HandlerFunctionInterface)

// RouteWrapper
func (rtp *RTPCore) HandlePath(path string, handleFunc HandlerFunctionInterface) {
	testRequest := new(RTPRequest)
	rtp.Server.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		resp := &RTPResponse{}
		switch r.Method {
		case POSTMETHOD:
			*resp = handleFunc.POST(*testRequest)
		case GETMETHOD:
			*resp = handleFunc.GET(*testRequest)
		case PUTMETHOD:
			*resp = handleFunc.PUT(*testRequest)
		case DELETEMETHOD:
			*resp = handleFunc.DELETE(*testRequest)
		}
	})
}
