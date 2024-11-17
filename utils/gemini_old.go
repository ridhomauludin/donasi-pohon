package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type GeminiRequest struct {
	Prompt string `json:"prompt"`
}

type GeminiResponse struct {
	Goal string `json:"goal"`
}

func GenerateGoal(prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	baseURL := os.Getenv("GEMINI_BASE_URL")
	url := baseURL + "?key=" + apiKey

	// Siapkan data untuk permintaan
	// requestData := GeminiRequest{Prompt: prompt}
	requestData := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}
	requestBody, _ := json.Marshal(requestData)

	// Kirim permintaan ke Gemini AI
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Cek jika respon berhasil
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(string(body))
	}

	// Decode respon
	var geminiResp GeminiResponse
	err = json.NewDecoder(resp.Body).Decode(&geminiResp)
	if err != nil {
		return "", err
	}

	return geminiResp.Goal, nil
}