package ai

import "github.com/x-sushant-x/Zocket/model"

type Suggestions interface {
	SuggestTasks(taskStats *model.TasksStats) (string, error)
}
