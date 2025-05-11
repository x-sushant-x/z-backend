package config

import (
	"context"
	"google.golang.org/genai"
	"log"
	"os"
	"sync"
)

var GeminiAPIClient *genai.Client
var geminiOnce sync.Once

func NewGeminiClient() {
	geminiOnce.Do(func() {
		apiKey := os.Getenv("GEMINI_API_KEY")
		if len(apiKey) == 0 {
			log.Fatal("GEMINI_API_KEY not provided in env.")
		}

		ctx := context.Background()

		client, err := genai.NewClient(ctx, &genai.ClientConfig{
			APIKey:  apiKey,
			Backend: genai.BackendGeminiAPI,
		})
		if err != nil {
			log.Fatal("unable to initialize Gemini API Client")
		}

		GeminiAPIClient = client
	})
}
