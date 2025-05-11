package config

import (
	"log"
	"os"
	"sync"

	"github.com/sashabaranov/go-openai"
)

var (
	OpenAIClient *openai.Client
	once         sync.Once
)

func NewOpenAIClient() {
	once.Do(func() {
		apiKey := os.Getenv("OPEN_AI_API_KEY")
		if len(apiKey) == 0 {
			log.Fatal("OPEN_AI_API_KEY not provided in env.")
		}

		OpenAIClient = openai.NewClient(apiKey)
	})
}
