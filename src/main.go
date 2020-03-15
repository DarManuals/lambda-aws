package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	bot          *botapi.BotAPI
	weatherToken string
	httpCli      = http.DefaultClient
)

func init() {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		fmt.Printf("ERROR: no token")
		os.Exit(1)
	}
	weatherToken, ok = os.LookupEnv("WEATHER_TOKEN")
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

	var replyStr = `echo: ` + payload.Msg.Text // default

	if strings.HasPrefix(payload.Msg.Text, W_CMD) {
		w, err := getWeather(strings.TrimPrefix(payload.Msg.Text, W_CMD))
		if err != nil {
			fmt.Printf("w err: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 200}, nil
		}
		replyStr = w.String()
	}

	if _, err := bot.Send(botapi.NewMessage(payload.Msg.Chat.ID, replyStr)); err != nil {
		fmt.Printf("ERROR: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 200}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
