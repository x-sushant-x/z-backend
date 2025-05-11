package ai

import (
	"context"
	"encoding/json"
	"fmt"
	customErrors "github.com/x-sushant-x/Zocket/errors"
	"github.com/x-sushant-x/Zocket/model"
	"google.golang.org/genai"
	"log"
)

type GeminiService struct {
	client *genai.Client
	model  string
}

func NewGeminiService(client *genai.Client) *GeminiService {
	return &GeminiService{
		client: client,
		model:  "gemini-2.0-flash",
	}
}

func (g GeminiService) SuggestTasks(taskStats *model.TasksStats) (string, error) {
	payloadJSON, err := json.MarshalIndent(taskStats, "", "  ")
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

	res, err := g.client.Models.GenerateContent(context.Background(), g.model, genai.Text(prompt), nil)

	if err != nil {
		log.Println("gemini error: " + err.Error())
		return "", customErrors.ErrInternalServerError
	}

	return res.Text(), nil
}
