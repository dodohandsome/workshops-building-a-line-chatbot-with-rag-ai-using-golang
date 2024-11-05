package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/pinecone-io/go-pinecone/pinecone"
	"google.golang.org/protobuf/types/known/structpb"
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

type Profile struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

type Message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
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

type ChatCompletionResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type messageHandler func(event Event) (interface{}, error)

var messages = make(map[string][]Message)

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
var client = resty.New()
var pc *pinecone.Client
var idxConnection *pinecone.IndexConnection

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx := context.Background()
	pc, err = pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: os.Getenv("PINECONE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create Client: %v", err)
	}
	indexName := os.Getenv("INDEX_NAME")
	idxModel, err := pc.DescribeIndex(ctx, indexName)
	if err != nil {
		log.Fatalf("Failed to describe index \"%v\": %v", indexName, err)
	}

	idxConnection, err = pc.Index(pinecone.NewIndexConnParams{Host: idxModel.Host})
	if err != nil {
		log.Fatalf("Failed to create IndexConnection for Host %v: %v", idxModel.Host, err)
	}

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
		loadingMessage(event.Source.UserID)
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

func initMessages(event Event) {
	userID := event.Source.UserID
	if _, exists := messages[userID]; !exists || len(messages[userID]) == 0 {
		profile, err := getProfile(userID)
		if err != nil {
			fmt.Printf("Error getting profile: %v\n", err)
		}
		messages[userID] = []Message{
			{Role: "system", Content: "คุณเป็นผู้ช่วยอัจฉริยะที่ออกแบบมาเพื่อตอบคำถามโดยใช้ Retrieval-Augmented Generation (RAG) วิธีการของคุณคือการประเมินข้อมูลที่ดึงมาจากฐานข้อมูลตามความเกี่ยวข้องกับคำถามของผู้ใช้ หากข้อมูลที่ดึงมามีคะแนนความเกี่ยวข้องมากกว่า 0.90 และตรงกับบริบทหรือคำหลักที่ผู้ใช้ระบุ ให้ตอบกลับด้วยรายการสินค้าที่สอดคล้องในลักษณะที่เป็นมิตร แต่หากไม่มีข้อมูลที่เกี่ยวข้อง ให้แจ้งผู้ใช้ว่ารอสักครู่เพื่อติดต่อเจ้าหน้าที่"},
			{Role: "user", Content: fmt.Sprintf("ฉันมีชื่อว่า %s", profile.DisplayName)},
			{Role: "assistant", Content: fmt.Sprintf("คุณชื่อ %s", profile.DisplayName)},
		}
	}

}

func handleTextMessage(event Event) (interface{}, error) {
	// prompt := event.Message.Text
	// embedding, _ := getEmbeddings(prompt)
	// results, _ := queryVectors(context.Background(), embedding)
	// ragResultsText := "ผลลัพธ์จาก RAG:\n"
	// if len(results) == 0 {
	// 	// return nil, nil
	// }
	// for i, match := range results {
	// 	ragResultsText += fmt.Sprintf("%d. text: %s, score: %f\n", i+1, match.Metadata.Text, match.Score)
	// }

	// fmt.Println("ragResultsText :", ragResultsText)

	// userID := event.Source.UserID
	// responseChat, err := generateChatResponse(messages[userID], prompt, ragResultsText)
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

func getTokenStateless() (string, error) {
	endpoint := "https://api.line.me/oauth2/v3/token"
	data := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     os.Getenv("CHANNEL_ID"),
		"client_secret": os.Getenv("CHANNEL_SECRET"),
	}

	resp, err := postFormRequest(endpoint, data)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	if accessToken, ok := result["access_token"].(string); ok {
		return accessToken, nil
	}

	return "", fmt.Errorf("failed to get access token")
}

func ReplyMessage(replyToken string, messages interface{}) error {
	messagesToSend, ok := messages.([]interface{})
	if !ok {
		messagesToSend = []interface{}{messages}
	}

	accessToken, err := getTokenStateless()
	if err != nil {
		return err
	}

	endpoint := "https://api.line.me/v2/bot/message/reply"
	payload := map[string]interface{}{
		"replyToken": replyToken,
		"messages":   messagesToSend,
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

func loadingMessage(lineUserId string) error {
	accessToken, err := getTokenStateless()
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

func getProfile(userID string) (*Profile, error) {
	accessToken, err := getTokenStateless()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://api.line.me/v2/bot/profile/%s", userID)

	resp, err := getRequest(endpoint, accessToken)

	if err != nil {
		return nil, err
	}

	var profile Profile
	if err := json.Unmarshal(resp, &profile); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &profile, nil
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

func getEmbeddings(text string) ([]float64, error) {
	endpoint := "https://api.openai.com/v1/embeddings"
	payload := map[string]interface{}{
		"input": text,
		"model": "text-embedding-ada-002",
	}

	resp, err := postJSONRequest(endpoint, os.Getenv("OPENAI_API_KEY"), payload)
	if err != nil {
		return nil, err
	}

	var embeddingsResponse EmbeddingsResponse
	if err := json.Unmarshal(resp.Body(), &embeddingsResponse); err != nil {
		return nil, err
	}

	if len(embeddingsResponse.Data) > 0 {
		return embeddingsResponse.Data[0].Embedding, nil
	}

	return nil, fmt.Errorf("no embeddings returned from OpenAI")
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

func generateVectors(ctx context.Context) {

	segmentLength := 5
	overlap := 2
	batchSize := 10

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	var wg sync.WaitGroup
	var batchVectors []*pinecone.Vector

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		segments := splitTextWithOverlap(line, segmentLength, overlap)

		for lineSeg, segment := range segments {
			wg.Add(1)
			go func(line string, segment string, lineNumber int, lineSeg int) {
				defer wg.Done()

				embedding, _ := getEmbeddings(segment)
				id := "vec-" + strconv.Itoa(lineNumber) + "-" + strconv.Itoa(lineSeg)
				metadataMap := map[string]interface{}{
					"text": line,
					"cut":  segment,
				}

				metadata, _ := convertMetadata(metadataMap)
				vectors := &pinecone.Vector{
					Id:       id,
					Values:   convertToFloat32(embedding),
					Metadata: metadata,
				}
				batchVectors = append(batchVectors, vectors)

				if len(batchVectors) >= batchSize {
					batch := batchVectors
					batchVectors = nil

					_, err := idxConnection.UpsertVectors(ctx, batch)
					if err != nil {
						log.Printf("Failed to upsert batch vectors: %v", err)
					} else {
						fmt.Printf("Successfully upserted %d vector(s) for batch\n", len(batch))
					}
				}

			}(line, segment, lineNumber, lineSeg)
		}
		lineNumber++
	}

	// Wait for all goroutines to finish
	wg.Wait()
	if len(batchVectors) > 0 {
		_, err := idxConnection.UpsertVectors(ctx, batchVectors)
		if err != nil {
			log.Printf("Failed to upsert remaining vectors: %v", err)
		} else {
			fmt.Printf("Successfully upserted %d remaining vector(s)!\n", len(batchVectors))
		}
	}
	fmt.Println("All vectors upserted successfully.")
}

func convertToFloat32(input []float64) []float32 {
	output := make([]float32, len(input))
	for i, v := range input {
		output[i] = float32(v)
	}
	return output
}

func convertMetadata(metadata map[string]interface{}) (*structpb.Struct, error) {
	return structpb.NewStruct(metadata)
}
func convertQueryMetadata(metadata *structpb.Struct) (Metadata, error) {
	meta := Metadata{}
	for key, value := range metadata.GetFields() {
		switch key {
		case "cut":
			meta.Cut = value.GetStringValue()
		case "text":
			meta.Text = value.GetStringValue()
		}
	}
	return meta, nil
}

func queryVectors(ctx context.Context, vector []float64) ([]QueryMatch, error) {
	queryVector := convertToFloat32(vector)
	queryResponse, err := idxConnection.QueryByVectorValues(ctx, &pinecone.QueryByVectorValuesRequest{
		TopK:            5, // Number of nearest vectors to retrieve
		Vector:          queryVector,
		IncludeMetadata: true, // Include metadata in the response
	})
	if err != nil {
		return nil, err
	}

	// Process the response
	var results []QueryMatch
	fmt.Println("Query Results:")
	for _, match := range queryResponse.Matches {
		metadata, err := convertQueryMetadata(match.Vector.Metadata)
		if err != nil {
			log.Fatalf("Error converting metadata: %v", err)
		}
		if match.Score >= 0.9 {
			results = append(results, QueryMatch{
				ID:       match.Vector.Id,
				Score:    match.Score,
				Metadata: metadata,
			})

		}

	}

	return results, nil
}

func generateChatResponse(messages []Message, promt string, assistant string) (string, error) {
	endpoint := "https://api.openai.com/v1/chat/completions"

	messages = append(messages, Message{Role: "user", Content: promt}, Message{Role: "assistant", Content: assistant})

	payload := map[string]interface{}{
		"model":    "gpt-4o",
		"messages": messages,
	}

	resp, err := postJSONRequest(endpoint, os.Getenv("OPENAI_API_KEY"), payload)
	if err != nil {
		return "", err
	}

	var completion ChatCompletionResponse
	if err := json.Unmarshal(resp.Body(), &completion); err != nil {
		return "", err
	}

	if len(completion.Choices) > 0 {
		return completion.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}
