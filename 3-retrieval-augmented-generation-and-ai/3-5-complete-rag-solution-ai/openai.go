package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GenerateChatResponse(messages []Message, promt string, assistant string) (string, error) {
	endpoint := "https://api.openai.com/v1/chat/completions"

	messages = append(messages, Message{Role: "user", Content: promt}, Message{Role: "assistant", Content: assistant})

	payload := map[string]interface{}{
		"model":    "gpt-4o",
		"messages": messages,
	}

	resp, err := PostJSONRequest(endpoint, os.Getenv("OPENAI_API_KEY"), payload)
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

func GetEmbeddings(text string) ([]float64, error) {
	endpoint := "https://api.openai.com/v1/embeddings"
	payload := map[string]interface{}{
		"input": text,
		"model": "text-embedding-ada-002",
	}

	resp, err := PostJSONRequest(endpoint, os.Getenv("OPENAI_API_KEY"), payload)
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
