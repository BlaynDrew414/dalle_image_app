package models

type OpenAIRequestBody struct {
	Model     string   `json:"model"`
	Prompt    string   `json:"prompt"`
	MaxTokens int      `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

type OpenAIResponseBody struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}