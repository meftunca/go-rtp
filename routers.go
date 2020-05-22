package rtp

type HandlerFunctionInterface interface {
	POST(RTPRequest) RTPResponse
	GET(RTPRequest) RTPResponse
	DELETE(RTPRequest) RTPResponse
	PUT(RTPRequest) RTPResponse
}
type HandlerFunc func(RTPRequest)

func (h *RTPRequest) HandlerFunc(newRequest *RTPRequest, f HandlerFunc) *RTPResponse {
	// return r.Handler(http.HandlerFunc(f))
	resp := &RTPResponse{}
	/*
		TODO Methoda göre sınıflandırılma ayarlanacak
		TODO Methoda uygun Fonksiyon çağırılacak
		TODO Çağırılan fonksiyondan dönen değer return ile channel'a gönderilip kullanıcıya iletilecek
	*/
	return resp
}
