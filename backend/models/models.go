package models

type OpenAIRequestBody struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

type OpenAIResponseBody struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

type Image struct {
	ID     string `bson:"_id"`
	Image  []byte `bson:"image"`
	Prompt string `bson:"prompt"`
}

type GenerateRequest struct {
    Prompt string `json:"prompt"`
    NumImages int `json:"n"`
    Size string `json:"size"`
}

type GenerateResponse struct {
    Created int64 `json:"created"`
    Data []struct {
        URL string `json:"url"`
    } `json:"data"`
}

type GenerateImageRequestBody struct {
    Description string `json:"description"`
}

type GenerateImageResponseBody struct {
    ImageUrls []string `json:"image_urls"`
}