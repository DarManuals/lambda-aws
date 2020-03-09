package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

type Model struct {
	Name string `json:"name"`
}

func HandleRequest(_ context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var m Model

	err := json.Unmarshal([]byte(event.Body), &m)
	if err != nil {
		log.Fatal("err: ", err)
	}

	m.Name = "Name is " + m.Name

	b, _ := json.Marshal(m)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
