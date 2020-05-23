package rtp

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/meftunca/websocket/models"
)

type ExampleHandleFucntion interface {
	HandlerFunctionInterface
}
type CustomEndpointTypes struct {
}

type PostDataType struct {
	Name string
}

/*
	TODO Gelen İstekleri İstenilen Şekilde Parse Edebilmek İçin yardımcı fonksiyonlar yaz
	!    Eğer Yukarıdaki Adımı es geçersen bu kısım çok karmaşık olacak !!!
*/

func (c *CustomEndpointTypes) POST(rtp RTPRequest, db *gorm.DB) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("\n\n----[POST]--- IS RUNNING\n\n", rtp.POSTDATA)
	postData := PostDataType{}
	bytePostData, _ := json.Marshal(rtp.POSTDATA)
	err := json.Unmarshal(bytePostData, &postData)
	if err != nil {
		panic(err)
	}
	testModels := &models.RTPTESTMODELS{}
	testModels.Name = postData.Name
	db.Create(&testModels)
	// resp.DATA = testModels
	resp.INSERTEDROWID = int64(testModels.ID)
	fmt.Println("\n\n----[POST]--- IS Successfull\t-- ID =>", testModels.ID)
	return *resp
}

func (c *CustomEndpointTypes) GET(rtp RTPRequest, db *gorm.DB) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("\n\n----[GET]--- IS RUNNING\n\n")

	return *resp
}

func (c *CustomEndpointTypes) PUT(rtp RTPRequest, db *gorm.DB) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("\n\n----[PUT]--- IS RUNNING\n\n")

	return *resp
}

func (c *CustomEndpointTypes) DELETE(rtp RTPRequest, db *gorm.DB) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("\n\n----[DELETE]--- IS RUNNING\n\n")

	return *resp
}

var ExampleHandleFucntionController ExampleHandleFucntion = new(CustomEndpointTypes)
