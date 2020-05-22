package rtp

/*
	! Request  body
	USERID: Auth işlemleri için bu değeri istek gövdesinde kullanıcı göndermek zorunda
	METHOD: POST, DELETE, GET, GETONE, PUT
	SEARCHOBJECT(POST, DELETE, GET, GETONE, PUT): Bu kısım kullanıcının standart isteklerinde gönderdiği ‘URLSEARCHPARAMS’ a karşılık gelmektedir.
	POSTDATA(POST;PUT): Eklenecek data nesnesi
*/
type MapStringFace = map[string]interface{}

type RTPRequest struct {
	PATH         string        `json:"path"`
	USERID       string        `json:"userId"`
	METHOD       string        `json:"method"`
	SEARCHOBJECT MapStringFace `json:"searchObject"`
	POSTDATA     MapStringFace `json:"postData"`
	// ! Auth (v2) de eklenecek
	AUTHID string `json:"-"` //`json:"authTokenId"`

}

/*
	! Response  body
 	DATA(GET,GETONE): Sonuçların listeleneceği method
	DELETEDROWID(DELETE): Silinen satırın id değeri
	UPDATEDROWID(PUT): Güncellenen satır id değeri
	INSERTEDROWID(POST): Eklenen satır id değeri
*/

type RTPResponse struct {
	PATH          string          `json:"path"`
	DATA          []MapStringFace `json:"data"`
	DELETEDROWID  int64           `json:"deletedRowId"`
	UPDATEDROWID  int64           `json:"updatedRowId"`
	INSERTEDROWID int64           `json:"insertedRowId"`
}

const (
	POSTMETHOD   = "POST"
	PUTMETHOD    = "PUT"
	DELETEMETHOD = "DELETE"
	GETMETHOD    = "GET"
	GETONEMETHOD = "GETONE"
)
