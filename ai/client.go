package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Provider string

const (
	ProviderOllama    Provider = "ollama"
	ProviderOpenAI    Provider = "openai"
	ProviderAnthropic Provider = "anthropic"
	ProviderGroq      Provider = "groq"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type ChatResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   Message   `json:"message"`
	Done      bool      `json:"done"`
}

type Client struct {
	provider Provider
	endpoint string
	apiKey   string
	model    string
}

func New(provider Provider, endpoint, model string, apiKey string) *Client {
	return &Client{
		provider: provider,
		endpoint: strings.TrimSuffix(endpoint, "/"),
		model:    model,
		apiKey:   apiKey,
	}
}

func (c *Client) Chat(messages []Message) (string, error) {
	reqBody := ChatRequest{
		Model:    c.model,
		Messages: messages,
		Stream:   false,
	}

	var response ChatResponse
	var err error

	switch c.provider {
	case ProviderOllama:
		response, err = c.ollamaChat(reqBody)
	default:
		err = fmt.Errorf("unsupported provider: %s", c.provider)
	}

	if err != nil {
		return "", err
	}

	return response.Message.Content, nil
}

func (c *Client) ollamaChat(req ChatRequest) (ChatResponse, error) {
	url := fmt.Sprintf("%s/api/chat", c.endpoint)

	data, err := json.Marshal(req)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpClient := &http.Client{Timeout: 60 * time.Second}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return ChatResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != 200 {
		return ChatResponse{}, fmt.Errorf("API error: %s", string(body))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return ChatResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return chatResp, nil
}
