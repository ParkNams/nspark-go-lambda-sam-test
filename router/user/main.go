package user

import (
	"encoding/json"
	"log"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ResponseBody struct {
	Message string `json:message`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data, err := json.Marshal(ResponseBody{Message: "user"})
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
