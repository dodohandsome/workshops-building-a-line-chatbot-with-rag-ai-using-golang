package main

import (
	"context"
	"log"
	"os"
	"strings"

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

type EmbeddingsResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
}

type Metadata struct {
	Cut  string `json:"cut"`
	Text string `json:"text"`
}

type QueryMatch struct {
	ID       string   `json:"id"`
	Score    float32  `json:"score"`
	Metadata Metadata `json:"metadata"`
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

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx := context.Background()
	// InitPinecone(ctx)
	// generateVectors(ctx)

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
		result, err := handler(event)
		if err != nil {
			log.Println("Error handling event:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
		}
		results[i] = result
	}

	return c.JSON(results)
}

func handleTextMessage(event Event) (interface{}, error) {
	// embedding, _ := GetEmbeddings(event.Message.Text)
	// results, _ := QueryVectors(context.Background(), embedding)
	if len(results) == 0 {
		return nil, nil
	}

	reply := map[string]string{
		"type": "text",
		"text": results[0].Metadata.Text,
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handleImageMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello image",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func handleAudioMessage(event Event) (interface{}, error) {
	reply := map[string]string{
		"type": "text",
		"text": "hello audio",
	}

	if err := ReplyMessage(event.ReplyToken, reply); err != nil {
		return nil, err
	}

	return reply, nil
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

func splitTextWithOverlap(text string, length int, overlap int) []string {
	words := strings.Fields(text) // แยกคำโดยใช้ whitespace
	var result []string

	for i := 0; i < len(words); i += (length - overlap) {
		end := i + length
		if end > len(words) {
			end = len(words)
		}
		part := strings.Join(words[i:end], " ")
		result = append(result, part)

		if end == len(words) { // หยุดเมื่อตัดข้อความทั้งหมดแล้ว
			break
		}
	}
	return result
}

// func generateVectors(ctx context.Context) {

// 	segmentLength := 5
// 	overlap := 2
// 	batchSize := 10

// 	file, err := os.Open("example.txt")
// 	if err != nil {
// 		fmt.Printf("Error opening file: %v\n", err)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	lineNumber := 0
// 	var wg sync.WaitGroup
// 	var batchVectors []*pinecone.Vector

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fmt.Println(line)
// 		segments := splitTextWithOverlap(line, segmentLength, overlap)

// 		for lineSeg, segment := range segments {
// 			wg.Add(1)
// 			go func(line string, segment string, lineNumber int, lineSeg int) {
// 				defer wg.Done()

// 				embedding, _ := GetEmbeddings(segment)
// 				id := "vec-" + strconv.Itoa(lineNumber) + "-" + strconv.Itoa(lineSeg)
// 				metadataMap := map[string]interface{}{
// 					"text": line,
// 					"cut":  segment,
// 				}

// 				metadata, _ := ConvertMetadata(metadataMap)
// 				vectors := &pinecone.Vector{
// 					Id:       id,
// 					Values:   ConvertToFloat32(embedding),
// 					Metadata: metadata,
// 				}
// 				batchVectors = append(batchVectors, vectors)

// 				if len(batchVectors) >= batchSize {
// 					batch := batchVectors
// 					batchVectors = nil

// 					_, err := idxConnection.UpsertVectors(ctx, batch)
// 					if err != nil {
// 						log.Printf("Failed to upsert batch vectors: %v", err)
// 					} else {
// 						fmt.Printf("Successfully upserted %d vector(s) for batch\n", len(batch))
// 					}
// 				}

// 			}(line, segment, lineNumber, lineSeg)
// 		}
// 		lineNumber++
// 	}

// 	wg.Wait()
// 	if len(batchVectors) > 0 {
// 		_, err := idxConnection.UpsertVectors(ctx, batchVectors)
// 		if err != nil {
// 			log.Printf("Failed to upsert remaining vectors: %v", err)
// 		} else {
// 			fmt.Printf("Successfully upserted %d remaining vector(s)!\n", len(batchVectors))
// 		}
// 	}
// 	fmt.Println("All vectors upserted successfully.")
// }
