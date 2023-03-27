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
