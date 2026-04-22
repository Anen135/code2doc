package llm

import (
	"bytes"
	"code2doc/internal/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	APIKey string
	Model  string
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c *Client) Ask(prompt string) (string, error) {
	url := "https://openrouter.ai/api/v1/chat/completions"

	logger.Debug(fmt.Sprintf("Sending request to LLM API with model: %s", c.Model))

	body, _ := json.Marshal(map[string]interface{}{
		"model": c.Model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(fmt.Sprintf("HTTP request failed: %v", err))
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errorBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read error body: %w", err)
		}
		errAPI := fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(errorBody))
		logger.ErrorV(errAPI)
		return "", errAPI
	}

	var res Response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		logger.Error(fmt.Sprintf("Failed to decode API response: %v", err))
		return "", err
	}

	if len(res.Choices) > 0 {
		logger.Debug("Successfully received response from LLM API")
		return res.Choices[0].Message.Content, nil
	}
	logger.Warning("Empty response received from LLM API")
	return "No response", nil
}
