package rtp

import "encoding/json"

func (h *Hub) ByteToRTPRequest(message []byte) (RTPRequest, error) {
	req := &RTPRequest{}
	err := json.Unmarshal(message, &req)
	return *req, err
}
