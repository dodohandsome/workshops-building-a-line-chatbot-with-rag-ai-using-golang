package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

type Event struct {
	Type    string `json:"type"`
	Message struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Text       string `json:"text"`
		QuoteToken string `json:"quoteToken"`
	} `json:"message"`
	ReplyToken string `json:"replyToken"`
	Source     struct {
		UserID  string `json:"userId"`
		GroupID string `json:"groupId"`
	} `json:"source"`
	Postback struct {
		Data string `json:"data"`
	} `json:"postback"`
	Beacon struct {
		Hwid string `json:"hwid"`
	} `json:"beacon"`
}

type WebhookRequest struct {
	Events []Event `json:"events"`
}

type messageHandler func(event Event) (interface{}, error)

var handlers = map[string]messageHandler{
	// "text":     handleTextMessage,
	// "image":    handleImageMessage,
	// "audio":    handleAudioMessage,
	// "file":     handleFileMessage,
	// "video":    handleVideoMessage,
	// "location": handleLocationMessage,
	// "sticker":  handleStickerMessage,
	// "postback": handlePostbackMessage,
	// "beacon":   handleBeaconMessage,
	// "follow":   handleFollowMessage,
	// "unfollow": handleUnfollowMessage,
}
var client = resty.New()

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app := fiber.New()
	app.Use(logger.New())

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/webhook", handleWebhook)
	log.Fatal(app.Listen(":" + port))

}

func handleWebhook(c *fiber.Ctx) error {
	var req WebhookRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	results := make([]interface{}, len(req.Events))
	for i, event := range req.Events {
		typeEvent := event.Type
		if event.Type == "message" {
			typeEvent = event.Message.Type
		}
		handler, found := handlers[typeEvent]
		if !found {
			log.Printf("No handler found for event type: %s", typeEvent)
			continue
		}
		// loadingMessage(event.Source.UserID)
		result, err := handler(event)
		if err != nil {
			log.Println("Error handling event:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		}
		results[i] = result
	}

	return c.JSON(results)
}

// func handleTextMessage(event Event) (interface{}, error) {

// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello text",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleImageMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello image",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleAudioMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello audio",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleFileMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello file",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleVideoMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello video",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleLocationMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello location",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleStickerMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello sticker",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handlePostbackMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello postback",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleFollowMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello follow",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleUnfollowMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello unfollow",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func handleBeaconMessage(event Event) (interface{}, error) {
// 	reply := map[string]string{
// 		"type": "text",
// 		"text": "hello beacon",
// 	}

// 	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }

// func getTokenStatelate() (string, error) {
// 	endpoint := "https://api.line.me/oauth2/v3/token"
// 	data := map[string]string{
// 		"grant_type":    "client_credentials",
// 		"client_id":     os.Getenv("CHANNEL_ID"),
// 		"client_secret": os.Getenv("CHANNEL_SECRET"),
// 	}

// 	resp, err := postFormRequest(endpoint, data)
// 	if err != nil {
// 		return "", err
// 	}

// 	var result map[string]interface{}
// 	if err := json.Unmarshal(resp, &result); err != nil {
// 		return "", fmt.Errorf("error unmarshalling response: %v", err)
// 	}

// 	if accessToken, ok := result["access_token"].(string); ok {
// 		return accessToken, nil
// 	}

// 	return "", fmt.Errorf("failed to get access token")
// }

// func ReplyMessage(replyToken string, messages interface{}) error {
// 	messagesToSend, ok := messages.([]interface{})
// 	if !ok {
// 		messagesToSend = []interface{}{messages}
// 	}

// 	accessToken, err := getTokenStatelate()
// 	if err != nil {
// 		return err
// 	}

// 	endpoint := "https://api.line.me/v2/bot/message/reply"
// 	payload := map[string]interface{}{
// 		"replyToken": replyToken,
// 		"messages":   messagesToSend,
// 	}

// 	resp, err := postJSONRequest(endpoint, accessToken, payload)
// 	if err != nil {
// 		return err
// 	}

// 	if resp.StatusCode() != 200 {
// 		return fmt.Errorf("error replying message: %s", resp.String())
// 	}

// 	return nil
// }

func loadingMessage(lineUserId string) error {
	accessToken, err := getTokenStatelate()
	if err != nil {
		return err
	}

	endpoint := "https://api.line.me/v2/bot/chat/loading/start"
	payload := map[string]interface{}{
		"chatId":         lineUserId,
		"loadingSeconds": 20,
	}

	resp, err := postJSONRequest(endpoint, accessToken, payload)

	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("error replying message: %s", resp.String())
	}

	return nil
}

func getRequest(endpoint string, accessToken string) ([]byte, error) {
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		Get(endpoint)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func postFormRequest(endpoint string, data map[string]string) ([]byte, error) {
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(data).
		Post(endpoint)

	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func postJSONRequest(endpoint, accessToken string, payload interface{}) (*resty.Response, error) {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		SetBody(payload).
		Post(endpoint)

	return resp, err
}

// loadingMessage(event.Source.UserID)
// var result interface{}
// var err error
// typeEvent := event.Type
// if event.Type == "message" {
// 	typeEvent = event.Message.Type
// }
// switch typeEvent {

// case "text":
// 	result, err = handleTextMessage(event)
// case "image":
// 	result, err = handleImageMessage(event)
// case "audio":
// 	result, err = handleAudioMessage(event)
// case "file":
// 	result, err = handleFileMessage(event)
// case "video":
// 	result, err = handleVideoMessage(event)
// case "location":
// 	result, err = handleLocationMessage(event)
// case "sticker":
// 	result, err = handleStickerMessage(event)
// case "postback":
// 	result, err = handlePostbackMessage(event)
// case "beacon":
// 	result, err = handleBeaconMessage(event)
// case "follow":
// 	result, err = handleFollowMessage(event)
// case "unfollow":
// 	result, err = handleUnfollowMessage(event)
// default:
// 	log.Printf("No handler found for event type: %s", event.Type)
// 	continue
// }

// if err != nil {
// 	log.Println("Error handling event:", err)
// 	return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
// }

// // แสดงการโหลดข้อมูลให้ผู้ใช้ทราบ

// results[i] = result
