package rtp

import (
	"encoding/json"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (h *Hub) ByteToRTPRequest(message []byte) (RTPRequest, error) {
	req := &RTPRequest{}
	err := json.Unmarshal(message, &req)
	return *req, err
}

type respJson = map[string]interface{}

func StructToJson(v interface{}) respJson {
	respJsonModel := respJson{}
	byteArr := StructToByte(v)
	err := json.Unmarshal(byteArr, respJsonModel)
	if err != nil {
		panic(err)
	}
	return respJsonModel
}

func StructToByte(v interface{}) []byte {
	resp, _ := json.Marshal(v)
	return resp
}

func GenerateCustomStruct(v interface{}, target interface{}) error {
	bte := StructToByte(v)
	err := json.Unmarshal(bte, target)
	return err
}

func StringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano() * int64(rand.Intn(999999))))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateUniqueID(length int) string {
	return StringWithCharset(length, charset)
}
