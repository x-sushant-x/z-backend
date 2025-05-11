package requests

type TaskRequest struct {
	Description    string `json:"description"`
	Status         string `json:"status"`
	AssignedTo     *uint  `json:"assignedTo"`
	EstimatedHours int    `json:"estimatedHours"`
}

type UpdateStatusRequest struct {
	TaskId uint   `json:"taskId"`
	Status string `json:"status"`
}
