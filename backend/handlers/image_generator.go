package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type GenerateRequest struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	NumImages int    `json:"num_images"`
	Size      string `json:"size"`
}

type GenerateResponse struct {
	ImageURL string `json:"image_url"`
}

func GenerateImage(prompt string) (string, error) {
	// Create a new GenerateRequest with the necessary fields
	req := GenerateRequest{
		Model:     os.Getenv("MODEL_ID"),
		Prompt:    prompt,
		NumImages: 1,
		Size:      "512x512",
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	// Create a new HTTP request
	reqURL := "https://api.openai.com/v1/images/generations"
	reqMethod := "POST"
	reqHeaders := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + os.Getenv("OPENAI_API_KEY"),
	}
	reqBodyReader := bytes.NewReader(reqBody)

	httpReq, err := http.NewRequest(reqMethod, reqURL, reqBodyReader)
	if err != nil {
		return "", err
	}
	for k, v := range reqHeaders {
		httpReq.Header.Set(k, v)
	}

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the HTTP response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the response JSON to a GenerateResponse struct
	var respData struct {
		Data struct {
			Images []struct {
				URL string `json:"url"`
			} `json:"images"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &respData); err != nil {
		return "", err
	}
	imageURL := respData.Data.Images[0].URL

	return imageURL, nil
}
