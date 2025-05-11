package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	customErrors "github.com/x-sushant-x/Zocket/errors"
	"github.com/x-sushant-x/Zocket/model"
	"log"
)

type OpenAISvc struct {
	client *openai.Client
}

func NewOpenAISvc(client *openai.Client) OpenAISvc {
	return OpenAISvc{
		client: client,
	}
}

func (o OpenAISvc) SuggestTasks(taskStats *model.TasksStats) ([]model.TaskAssignment, error) {
	payloadJSON, err := json.MarshalIndent(taskStats, "", " ")
	if err != nil {
		log.Println("marshal error: " + err.Error())

		return nil, customErrors.ErrInternalServerError
	}

	prompt := fmt.Sprintf(`
	You are an AI task assignment assistant.

	Given the following users and new tasks:
	%s

	Suggest how to assign the tasks to users in a way that balances the workload. 
	Take into account the current number of tasks and total estimated hours per user. 
	Try to avoid overloading any user and distribute tasks fairly. 
	Respond in a clear JSON format with "assignments" like:
	[
  		{ "user": "John", "task": "Create submission video." },
  		{ "user": "Kevin", "task": "Another task" }
	]
	`, payloadJSON)

	resp, err := o.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-4o",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful assistant that suggests optimal task assignments.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		log.Println("chatgpt api error: " + err.Error())
		return nil, customErrors.ErrInternalServerError
	}

	taskSuggestions := resp.Choices[0].Message.Content

	fmt.Println("Response: " + taskSuggestions)

	var assignments []model.TaskAssignment

	err = json.Unmarshal([]byte(taskSuggestions), &assignments)
	if err != nil {
		log.Println("unmarshal error: " + err.Error())
		return nil, customErrors.ErrInternalServerError

	}

	return assignments, nil
}
