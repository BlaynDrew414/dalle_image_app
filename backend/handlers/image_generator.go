package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BlaynDrew414/dalle_image_app/backend/models"
	"github.com/sethgrid/pester"
)

func GenerateImage(description string) ([]byte, error) {
	openaiURL := "https://api.openai.com/v1/engines/davinci/images/generate"

	openaiRequestBody := models.OpenAIRequestBody{
		Model:     "image-alpha-001",
		Prompt:    "Generate image of " + description,
		MaxTokens: 1024,
		Temperature: 0.5,
	}

	requestBodyBytes, err := json.Marshal(openaiRequestBody)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey != "" {
		request.Header.Set("Authorization", "Bearer "+apiKey)
	}

	client := pester.New()
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var openaiResponseBody models.OpenAIResponseBody
	err = json.Unmarshal(responseBodyBytes, &openaiResponseBody)
	if err != nil {
		return nil, err
	}

	imageUrl := openaiResponseBody.Choices[0].Text
	imageResponse, err := http.Get(imageUrl)
	if err != nil {
		return nil, err
	}
	defer imageResponse.Body.Close()

	imageBytes, err := ioutil.ReadAll(imageResponse.Body)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}