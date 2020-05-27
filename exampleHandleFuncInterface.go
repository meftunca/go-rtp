package rtp

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/meftunca/websocket/models"
)

type ExampleHandleFucntion = HandlerFunctionInterface
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
	resp := &RTPResponse{Method: "POST"}
	fmt.Println("----[POST]--- IS RUNNING", rtp.POSTDATA)
	postData := PostDataType{}
	err := GenerateCustomStruct(rtp.POSTDATA, &postData)
	if err != nil {
		panic(err)
	}
	testModels := models.RTPTESTMODELS{Name: postData.Name}

	//Create Row
	db.Create(&testModels)

	// Get EffectedRowId
	resp.EFFECTEDROWID = int64(testModels.ID)

	fmt.Println("----[POST]--- IS Successfull\t-- ID =>", testModels.ID)
	return *resp
}

func (c *CustomEndpointTypes) GET(rtp RTPRequest, db *gorm.DB) RTPMultiGetResponse {
	resp := &RTPMultiGetResponse{Method: "GET"}
	fmt.Println("----[GET]--- IS RUNNING")
	allList := []models.RTPTESTMODELS{}
	allListByte := []map[string]interface{}{}
	db.Order("created_at DESC").Find(&allList)
	err := GenerateCustomStruct(allList, &allListByte)
	if err != nil {
		panic(err)
	}
	resp.DATA = allListByte
	return *resp
}

func (c *CustomEndpointTypes) PUT(rtp RTPRequest, db *gorm.DB) RTPResponse {
	resp := &RTPResponse{Method: "PUT"}
	fmt.Println("----[PUT]--- IS RUNNING", rtp.POSTDATA)
	putData := models.RTPTESTMODELS{}
	err := GenerateCustomStruct(rtp.POSTDATA, &putData)
	if err != nil {
		panic(err)
	}
	testModels := models.RTPTESTMODELS{}
	//Create Row
	allListByte := map[string]interface{}{}
	db.Model(&testModels).Where("id=?", putData.ID).Update("name", putData.Name)

	// Get EffectedRowId
	resp.EFFECTEDROWID = int64(putData.ID)
	err = GenerateCustomStruct(testModels, &allListByte)
	allListByte["id"] = putData.ID
	resp.DATA = allListByte
	fmt.Println("----[PUT]--- IS Successfull\t-- ID =>", testModels.ID)
	return *resp
}

func (c *CustomEndpointTypes) DELETE(rtp RTPRequest, db *gorm.DB) RTPResponse {
	fmt.Println("----[DELETE]--- IS RUNNING")
	models := models.RTPTESTMODELS{}

	err := GenerateCustomStruct(rtp.SEARCHOBJECT, &models)
	if err != nil {
	}
	db.Delete(&models)
	resp := &RTPResponse{Method: "DELETE", EFFECTEDROWID: int64(models.ID)}
	return *resp
}

func (c *CustomEndpointTypes) GETONE(rtp RTPRequest, db *gorm.DB) RTPResponse {
	resp := &RTPResponse{Method: "GETONE"}
	fmt.Println("----[GETONE]--- IS RUNNING")
	allList := models.RTPTESTMODELS{}
	allListByte := map[string]interface{}{}
	err := GenerateCustomStruct(rtp.SEARCHOBJECT, &allList)

	fmt.Println(rtp.SEARCHOBJECT)
	db.Order("created_at DESC").First(&allList)
	err = GenerateCustomStruct(allList, &allListByte)
	if err != nil {
		fmt.Println()
		fmt.Println("Hata Burada")
		panic(err)
	}
	resp.DATA = allListByte
	return *resp
}

func (c *CustomEndpointTypes) STAT(rtp RTPRequest, db *gorm.DB) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("----[STAT]--- IS RUNNING")

	return *resp
}

var ExampleHandleFucntionController ExampleHandleFucntion = new(CustomEndpointTypes)
