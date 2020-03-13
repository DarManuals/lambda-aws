package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Result string `json:"result"`
}

func HandleRequest(_ context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request
	_ = json.Unmarshal([]byte(event.Body), &req)

	result := "Name was: " + req.Name

	log.Println("INFO: log: GOT payload: ", event.Body)
	fmt.Println("INFO: fmt: GOT payload: ", event.Body)

	b, _ := json.Marshal(Response{Result: result})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
