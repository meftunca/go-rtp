package rtp

type HandlerFunction interface {
	POST(RTPRequest) RTPResponse
	GET(RTPRequest) RTPResponse
	DELETE(RTPRequest) RTPResponse
	PUT(RTPRequest) RTPResponse
}
