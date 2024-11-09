package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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
