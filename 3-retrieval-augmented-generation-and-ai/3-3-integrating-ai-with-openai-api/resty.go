package main

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()

func GetRequest(endpoint string, accessToken string) ([]byte, error) {
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		Get(endpoint)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func PostFormRequest(endpoint string, data map[string]string) ([]byte, error) {
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(data).
		Post(endpoint)

	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func PostJSONRequest(endpoint, accessToken string, payload interface{}) (*resty.Response, error) {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		SetBody(payload).
		Post(endpoint)

	return resp, err
}

func PostFormDataRequest(endpoint, accessToken string, formData map[string]string, filePath string) ([]byte, error) {
	client := resty.New()

	req := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		SetFile("file", filePath)

	for key, value := range formData {
		req.SetFormData(map[string]string{key: value})
	}

	resp, err := req.Post(endpoint)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("request failed: status code %d", resp.StatusCode())
	}

	return resp.Body(), nil
}
