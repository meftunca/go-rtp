package rtp

import "fmt"

type ExampleHandleFucntion interface {
	HandlerFunctionInterface
}
type CustomEndpointTypes struct {
}

func (c *CustomEndpointTypes) POST(RTPRequest) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("\n\n----[POST]--- IS RUNNING\n\n")

	return *resp
}

func (c *CustomEndpointTypes) GET(RTPRequest) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("\n\n----[GET]--- IS RUNNING\n\n")

	return *resp
}

func (c *CustomEndpointTypes) PUT(RTPRequest) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("\n\n----[PUT]--- IS RUNNING\n\n")

	return *resp
}

func (c *CustomEndpointTypes) DELETE(RTPRequest) RTPResponse {
	resp := &RTPResponse{}
	fmt.Println("\n\n----[DELETE]--- IS RUNNING\n\n")

	return *resp
}

var ExampleHandleFucntionController ExampleHandleFucntion = new(CustomEndpointTypes)
