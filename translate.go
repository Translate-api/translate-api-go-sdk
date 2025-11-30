// Package translateapi provides the official Go SDK for Translate API
//
// Get your API key at: https://translate-api.com
// Documentation: https://translate-api.com/documentation
package translateapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client is the Translate API client
type Client struct {
	APIKey  string
	BaseURL string
	HTTP    *http.Client
}

// TranslateRequest represents a translation request
type TranslateRequest struct {
	Text           string      `json:"text"`
	TargetLanguage interface{} `json:"target_language"`
}

// TranslateResponse represents a translation response
type TranslateResponse struct {
	Success        bool              `json:"success"`
	Translations   map[string]string `json:"translations"`
	SourceLanguage string            `json:"source_language,omitempty"`
	CharactersUsed int               `json:"characters_used"`
}

// NewClient creates a new Translate API client
//
// Example:
//
//	client := translateapi.NewClient("your-api-key")
func NewClient(apiKey string) *Client {
	if apiKey == "" {
		panic("API key is required. Get one at https://translate-api.com")
	}
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://translate-api.com/v1",
		HTTP:    &http.Client{},
	}
}

// Translate translates text to one or more target languages
//
// targetLanguage can be a single string ("es") or a slice ([]string{"es", "fr"})
//
// Example - Single language:
//
//	result, err := client.Translate("Hello", "es")
//	fmt.Println(result.Translations["es"]) // "Hola"
//
// Example - Multiple languages:
//
//	result, err := client.Translate("Hello", []string{"es", "fr"})
//	fmt.Println(result.Translations) // map[es:Hola fr:Bonjour]
func (c *Client) Translate(text string, targetLanguage interface{}) (*TranslateResponse, error) {
	reqBody := TranslateRequest{
		Text:           text,
		TargetLanguage: targetLanguage,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/translate", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp map[string]string
		json.NewDecoder(resp.Body).Decode(&errResp)
		errMsg := errResp["error"]
		if errMsg == "" {
			errMsg = fmt.Sprintf("translation failed (HTTP %d)", resp.StatusCode)
		}
		return nil, fmt.Errorf(errMsg)
	}

	var result TranslateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
