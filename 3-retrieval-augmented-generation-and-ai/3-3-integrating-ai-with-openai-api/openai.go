package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GenerateChatResponse(messages []Message, promt string) (string, error) {
	endpoint := "https://api.openai.com/v1/chat/completions"

	messages = append(messages, Message{Role: "user", Content: promt})

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

func TranscribeVoice(filepath string) (string, error) {
	formData := map[string]string{
		"model": "whisper-1",
	}

	response, err := PostFormDataRequest("https://api.openai.com/v1/audio/transcriptions", os.Getenv("OPENAI_API_KEY"), formData, filepath)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}

	err = json.Unmarshal([]byte(response), &result)
	if err != nil {
		return "", err
	}

	text, _ := result["text"].(string)

	return text, err
}
