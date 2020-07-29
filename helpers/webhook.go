package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yevhenshymotiuk/telegram-lambda-webhook/apigateway"
)

func setWebhook(
	request events.APIGatewayProxyRequest,
	telegramToken string,
) (apigateway.Response, error) {
	responseFail := apigateway.Response{StatusCode: 404}

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		return responseFail, err
	}

	url := fmt.Sprintf(
		"https://%s/%s/",
		request.Headers["Host"],
		request.RequestContext.Stage,
	)
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		return responseFail, err
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		return responseFail, err
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	var buf bytes.Buffer

	body, err := json.Marshal(
		map[string]interface{}{"message": "Webhook was successfully set!"},
	)
	if err != nil {
		return responseFail, err
	}
	json.HTMLEscape(&buf, body)

	resp := apigateway.Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers:         map[string]string{"Content-Type": "application/json"},
	}

	return resp, nil
}
