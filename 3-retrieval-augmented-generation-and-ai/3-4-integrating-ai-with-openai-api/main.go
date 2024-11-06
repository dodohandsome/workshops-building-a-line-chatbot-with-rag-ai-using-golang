package main

import (
	"log"
	"os"

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

type Message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}

type Profile struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

type ChatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type messageHandler func(event Event) (interface{}, error)

var handlers = map[string]messageHandler{
	"text":     handleTextMessage,
	"image":    handleImageMessage,
	"audio":    handleAudioMessage,
	"file":     handleFileMessage,
	"video":    handleVideoMessage,
	"location": handleLocationMessage,
	"sticker":  handleStickerMessage,
	"postback": handlePostbackMessage,
	"beacon":   handleBeaconMessage,
	"follow":   handleFollowMessage,
	"unfollow": handleUnfollowMessage,
}

var messages = make(map[string][]Message)

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
		LoadingMessage(event.Source.UserID)
		initMessages(event)
		result, err := handler(event)
		if err != nil {
			log.Println("Error handling event:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		}
		results[i] = result
	}

	return c.JSON(results)
}

// func initMessages(event Event) {
// 	userID := event.Source.UserID
// 	if _, exists := messages[userID]; !exists || len(messages[userID]) == 0 {
// 		profile, err := GetProfile(userID)
// 		if err != nil {
// 			fmt.Printf("Error getting profile: %v\n", err)
// 		}
// 		messages[userID] = []Message{
// 			{Role: "system", Content: "คุณชื่อลูฟี่ คุณเป็นผู้ช่วยส่วนตัวของฉัน คุณเป็นผู้เชี่ยวชาญในทุกๆ เรื่อง และคุณจะตอบฉันเป็นภาษาไทยเท่านั้น"},
// 			{Role: "user", Content: fmt.Sprintf("ฉันมีชื่อว่า %s", profile.DisplayName)},
// 			{Role: "assistant", Content: fmt.Sprintf("คุณชื่อ %s", profile.DisplayName)},
// 		}
// 	}

// }

func handleTextMessage(event Event) (interface{}, error) {
	// prompt := event.Message.Text
	// userID := event.Source.UserID
	// responseChat, err := GenerateChatResponse(messages[userID], prompt)
	// if err != nil {
	// 	fmt.Printf("Error generating chat response: %v\n", err)
	// 	return nil, err
	// }

	// messages[userID] = append(messages[userID], Message{Role: "user", Content: prompt})
	// messages[userID] = append(messages[userID], Message{Role: "assistant", Content: responseChat})

	// reply := map[string]string{
	// 	"type": "text",
	// 	"text": responseChat,
	// }
	// if err := ReplyMessage(event.ReplyToken, reply); err != nil {
	// 	return nil, err
	// }

	// return reply, nil
}

func handleImageMessage(event Event) (interface{}, error) {
	// imageContent, err := GetFileBytes(event.Message.ID)
	// if err != nil {
	// 	fmt.Printf("Error getting image content: %v\n", err)
	// 	return nil, err
	// }

	// imageBase64 := base64.StdEncoding.EncodeToString(imageContent)

	// userID := event.Source.UserID

	// messages[userID] = append(messages[userID], Message{Role: "user", Content: []interface{}{
	// 	map[string]interface{}{
	// 		"type": "image_url",
	// 		"image_url": map[string]interface{}{
	// 			"url": "data:image/jpeg;base64," + imageBase64,
	// 		},
	// 	},
	// }})
	// messages[userID] = append(messages[userID], Message{Role: "assistant", Content: "จากรูปคุณให้ฉันช่วยอะไรไหม"})

	// reply := map[string]string{
	// 	"type":       "text",
	// 	"text":       "จากรูปคุณให้ฉันช่วยอะไรไหม",
	// 	"quoteToken": event.Message.QuoteToken,
	// }

	// if err := ReplyMessage(event.ReplyToken, reply); err != nil {
	// 	return nil, err
	// }

	// return reply, nil
}

func handleAudioMessage(event Event) (interface{}, error) {
	// voiceContent, err := GetFileBytes(event.Message.ID)
	// if err != nil {
	// 	fmt.Printf("Error getting image content: %v\n", err)
	// 	return nil, err
	// }
	// timestamp := time.Now().Format("20060102150405")
	// fileName := fmt.Sprintf("/tmp/%s.m4a", timestamp)

	// err = ioutil.WriteFile(fileName, voiceContent, 0644)
	// if err != nil {
	// 	return nil, err
	// }
	// prompt, _ := TranscribeVoice(fileName)
	// os.Remove(fileName)

	// messagesAudio := []Message{
	// 	{Role: "system", Content: "โปรดปรับข้อความนี้ให้ถูกต้องทั้งในด้านไวยากรณ์และการสะกดและความถูกต้องของประโยค เนื่องจากอาจมีการเรียงคำหรือการใส่คำผิด จากนั้นให้แปลข้อความที่ถูกปรับแล้ว ถ้าเป็นภาษาจีนหรือต้นฉบับให้สะกดคำอ่านออกมาเป็นภาษาไทยให้ฉันด้วย"},
	// }
	// responsePrompt := fmt.Sprintf(`
	// โปรดแปลประโยคตามนี้:
	// 1. ถ้าเป็นภาษาไทย: ให้แปลเป็นภาษาอังกฤษ
	// 2. ถ้าเป็นภาษาอื่นที่ไม่ใช่ภาษาไทย: ให้แปลเป็นภาษาไทย ภาษาอังกฤษ และภาษาจีน
	// 3. ผลลัพธ์ควรประกอบไปด้วยภาษาอังกฤษ ภาษาไทย และภาษาต้นฉบับเสมอ
	// ตัวอย่าง output ที่ต้องการ ให้เป็นแบบด้านล่างนี้เท่านั้น :
	// Original (Japanese): おはようございまーす。私はリュウです。
	// คำอ่าน Original: โอะฮาโย โกะไซมา-สุ วะตะชิวะ ริว เดสุ

	// Thai: สวัสดีตอนเช้าครับ ผมชื่อริวครับ

	// English: Good morning. My name is Ryu.

	// chinese : 早安。我的名字是龙。
	// คำอ่าน chinese  : จ่าว-อาน หว่อ-เตอ-หมิง-จื๋อ-ชื่อ หลง

	// ประโยคด้านล่างนี้คือประโยคที่ต้องแปล:

	// %s
	// `, prompt)

	// responseChat, err := GenerateChatResponse(messagesAudio, responsePrompt)
	// if err != nil {
	// 	fmt.Printf("Error generating chat response: %v\n", err)
	// 	return nil, err
	// }

	// reply := map[string]string{
	// 	"type": "text",
	// 	"text": responseChat,
	// }

	// if err := ReplyMessage(event.ReplyToken, reply); err != nil {
	// 	return nil, err
	// }

	// return reply, nil
}

func handleFileMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello file",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handleVideoMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello video",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handleLocationMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello location",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handleStickerMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello sticker",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handlePostbackMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello postback",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handleFollowMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello follow",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handleUnfollowMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello unfollow",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handleBeaconMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello beacon",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}
