package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	"lambda.com/nspark/controller"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ResponseBody struct {
	Message string                 `json:message`
	Data    controller.ReturnValue `json:data`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("test")
	send := controller.UserController()
	data, err := json.Marshal(ResponseBody{Message: "test", Data: send})
	if err != nil {
		log.Panic(err)
	}
	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: 200,
	}, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lambda.Start(Handler)
}
