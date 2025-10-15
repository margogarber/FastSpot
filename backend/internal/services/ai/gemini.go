package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GeminiService handles AI recommendations using Gemini API
type GeminiService struct {
	APIKey string
}

// NewGeminiService creates a new Gemini service instance
func NewGeminiService(apiKey string) *GeminiService {
	return &GeminiService{
		APIKey: apiKey,
	}
}

// MoodAnswer represents a user's answer to a mood question
type MoodAnswer struct {
	QuestionID      string   `json:"questionId"`
	SelectedOptions []string `json:"selectedOptions"`
}

// RecommendationRequest represents the request to Gemini API
type RecommendationRequest struct {
	Answers []MoodAnswer `json:"answers"`
}

// RecommendationResponse represents the AI recommendation
type RecommendationResponse struct {
	ProductSlugs []string          `json:"productSlugs"`
	Reasoning    string            `json:"reasoning"`
	Options      map[string]string `json:"options"`     // optional customizations
	Ingredients  []string          `json:"ingredients"` // optional ingredients
}

// GetRecommendations gets product recommendations from Gemini API
func (s *GeminiService) GetRecommendations(answers []MoodAnswer, questions []map[string]interface{}, products []map[string]interface{}) (*RecommendationResponse, error) {
	if s.APIKey == "" {
		return nil, fmt.Errorf("Gemini API key not configured")
	}

	// Build prompt for Gemini
	prompt := s.buildPrompt(answers, questions, products)

	// Call Gemini API (using gemini-2.5-flash, disable thinking mode)
	apiURL := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=%s", s.APIKey)

	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
		"generationConfig": map[string]interface{}{
			"temperature":     0.9,
			"topK":            1,
			"topP":            1,
			"maxOutputTokens": 8192, // Max allowed for gemini-2.5-flash
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Debug: log prompt size
	fmt.Printf("📊 Gemini Prompt Size: %d characters\n", len(prompt))

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ HTTP Error: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body ONCE
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Failed to read response body: %v\n", err)
		return nil, err
	}

	// Check HTTP status after reading body
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("❌ Gemini API Error (Status %d): %s\n", resp.StatusCode, string(bodyBytes))
		return nil, fmt.Errorf("Gemini API error: %s", string(bodyBytes))
	}

	// Log first 500 chars of response for debugging
	debugLen := 500
	if len(bodyBytes) < debugLen {
		debugLen = len(bodyBytes)
	}
	fmt.Printf("📥 Gemini API Response (first %d chars): %s\n", debugLen, string(bodyBytes[:debugLen]))

	var geminiResp struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
		Error *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Status  string `json:"status"`
		} `json:"error"`
	}

	if err := json.Unmarshal(bodyBytes, &geminiResp); err != nil {
		fmt.Printf("❌ JSON Decode Error: %v\n", err)
		return nil, err
	}

	// Check for API error
	if geminiResp.Error != nil {
		fmt.Printf("❌ Gemini API returned error: %s (code %d)\n", geminiResp.Error.Message, geminiResp.Error.Code)
		return nil, fmt.Errorf("Gemini API error: %s", geminiResp.Error.Message)
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		fmt.Printf("❌ No candidates in Gemini response. Full response: %s\n", string(bodyBytes))
		return nil, fmt.Errorf("no response from Gemini API")
	}

	responseText := geminiResp.Candidates[0].Content.Parts[0].Text
	// Debug log (truncate if too long)
	textDebugLen := 200
	if len(responseText) < textDebugLen {
		textDebugLen = len(responseText)
	}
	fmt.Printf("✅ Gemini Response (first %d chars): %s\n", textDebugLen, responseText[:textDebugLen])

	// Parse AI response
	recommendation, err := s.parseAIResponse(responseText)
	if err != nil {
		return nil, err
	}

	return recommendation, nil
}

func (s *GeminiService) buildPrompt(answers []MoodAnswer, questions []map[string]interface{}, products []map[string]interface{}) string {
	prompt := "FastSpot AI: recommend 2-3 food items.\n\nCUSTOMER: "

	// Collect all selected labels (super compact) - LIMIT to first 4 answers to reduce token count
	var labels []string
	answerCount := 0
	for _, answer := range answers {
		if answerCount >= 4 {
			break // Limit to 4 answers to avoid thinking mode overflow
		}
		for _, q := range questions {
			if q["id"] == answer.QuestionID {
				if opts, ok := q["options"].([]map[string]interface{}); ok {
					for _, opt := range opts {
						if optValue, ok := opt["value"].(string); ok {
							for _, selected := range answer.SelectedOptions {
								if selected == optValue {
									if label, ok := opt["label"].(string); ok {
										labels = append(labels, label)
									}
								}
							}
						}
					}
				}
				answerCount++
			}
		}
	}
	prompt += strings.Join(labels, ", ") + "\n\nMENU:\n"

	// Menu (tags only, no descriptions)
	for _, product := range products {
		tags := ""
		if t, ok := product["tags"].([]string); ok && len(t) > 0 {
			tags = "[" + strings.Join(t, ",") + "]"
		}
		prompt += fmt.Sprintf("%s (%s) %s\n", product["name"], product["slug"], tags)
	}

	prompt += `\nPick 2-3 items. JSON: {"productSlugs":["slug1","slug2"],"reasoning":"why"}`
	return prompt
}

func (s *GeminiService) parseAIResponse(text string) (*RecommendationResponse, error) {
	// Try to extract JSON from response
	// Gemini sometimes adds markdown formatting, so we need to clean it

	// Remove markdown code blocks if present
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "```json")
	text = strings.TrimPrefix(text, "```")
	text = strings.TrimSuffix(text, "```")
	text = strings.TrimSpace(text)

	start := bytes.Index([]byte(text), []byte("{"))
	end := bytes.LastIndex([]byte(text), []byte("}"))

	if start == -1 || end == -1 {
		fmt.Printf("❌ No JSON braces found. Full response: %s\n", text)
		return nil, fmt.Errorf("no JSON found in response")
	}

	jsonText := text[start : end+1]
	fmt.Printf("🔍 Extracted JSON: %s\n", jsonText)

	var recommendation RecommendationResponse
	if err := json.Unmarshal([]byte(jsonText), &recommendation); err != nil {
		fmt.Printf("❌ JSON Parse Error: %v. JSON text: %s\n", err, jsonText)
		return nil, fmt.Errorf("failed to parse AI response: %w", err)
	}

	// Validate response
	if len(recommendation.ProductSlugs) == 0 {
		return nil, fmt.Errorf("AI returned no product recommendations")
	}

	fmt.Printf("✅ Parsed recommendations: %v\n", recommendation.ProductSlugs)
	return &recommendation, nil
}
