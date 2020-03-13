package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	bot *botapi.BotAPI
)

func init() {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		fmt.Printf("ERROR: no token")
		os.Exit(1)
	}

	var err error
	bot, err = botapi.NewBotAPI(token)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(2)
	}
}

type Payload struct {
	Msg botapi.Message `json:"message"`
}

func HandleRequest(_ context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("INFO: fmt: GOT body: %q", event.Body)

	var payload Payload
	if err := json.Unmarshal([]byte(event.Body), &payload); err != nil {
		fmt.Printf("ERROR: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 200}, nil
	}

	fmt.Printf("INFO: fmt: GOT payload: %+v", payload)

	m := botapi.NewMessage(payload.Msg.Chat.ID, `echo: `+payload.Msg.Text)

	if _, err := bot.Send(m); err != nil {
		fmt.Printf("ERROR: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 200}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
