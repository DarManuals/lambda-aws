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

func HandleRequest(_ context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	token, _ := os.LookupEnv("BOT_TOKEN")
	bot, err := botapi.NewBotAPI(token)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	fmt.Printf("INFO: fmt: GOT body: %+v", event.Body)

	var msg botapi.Message
	_ = json.Unmarshal([]byte(event.Body), &msg)

	fmt.Printf("INFO: fmt: GOT msg: %+v", msg)

	m := botapi.NewMessage(msg.Chat.ID, `echo: `+msg.Text)
	_, err = bot.Send(m)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
