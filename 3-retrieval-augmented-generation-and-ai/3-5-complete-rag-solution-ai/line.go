package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getTokenStateless() (string, error) {
	endpoint := "https://api.line.me/oauth2/v3/token"
	data := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     os.Getenv("CHANNEL_ID"),
		"client_secret": os.Getenv("CHANNEL_SECRET"),
	}

	resp, err := PostFormRequest(endpoint, data)
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

	resp, err := PostJSONRequest(endpoint, accessToken, payload)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("error replying message: %s", resp.String())
	}

	return nil
}

func LoadingMessage(lineUserId string) error {
	accessToken, err := getTokenStateless()
	if err != nil {
		return err
	}

	endpoint := "https://api.line.me/v2/bot/chat/loading/start"
	payload := map[string]interface{}{
		"chatId":         lineUserId,
		"loadingSeconds": 20,
	}

	resp, err := PostJSONRequest(endpoint, accessToken, payload)

	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("error replying message: %s", resp.String())
	}

	return nil
}

func GetProfile(userID string) (*Profile, error) {
	accessToken, err := getTokenStateless()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://api.line.me/v2/bot/profile/%s", userID)

	resp, err := GetRequest(endpoint, accessToken)

	if err != nil {
		return nil, err
	}

	var profile Profile
	if err := json.Unmarshal(resp, &profile); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &profile, nil
}
