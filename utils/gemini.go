package utils

import (
	"context"
	"donasiPohon/config"

	"github.com/google/generative-ai-go/genai"
)



func GenerateContent(prompt string) (text string, err error) {
	client := config.Gemini
    ctx := context.Background()
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if partString, ok := part.(genai.Text); ok {
					text += string(partString)
				}
			}
		}
	}

	return text, nil
}
