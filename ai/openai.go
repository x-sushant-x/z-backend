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

func (o OpenAISvc) SuggestTasks(taskStats *model.TasksStats) (string, error) {
	payloadJSON, err := json.MarshalIndent(taskStats, "", " ")
	if err != nil {
		log.Println("marshal error: " + err.Error())

		return "", customErrors.ErrInternalServerError
	}

	prompt := fmt.Sprintf(`
    You are an AI task assignment assistant.

    Given the following users and new tasks:
    %s

    Analyze the current workload of each user (number of tasks and total estimated hours) and suggest the most balanced task assignments. 
    Consider distributing the new tasks fairly, avoiding overloading any single user.

    Provide your suggestions as a bulleted list, with each item clearly stating which task should be assigned to which user and the reasoning behind the suggestion.

    For example:
    * [TASK_NAME_1] should be assigned to [USER_NAME_A] because they currently have a lighter workload (e.g., fewer tasks, lower total estimated hours).
    * [TASK_NAME_2] should be assigned to [USER_NAME_B] as their current workload is moderate and they have the capacity for more.
    * [TASK_NAME_3] should be assigned to [USER_NAME_C] to balance their heavier workload with a smaller task.
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
		return "", customErrors.ErrInternalServerError
	}

	taskSuggestions := resp.Choices[0].Message.Content

	return taskSuggestions, nil
}
