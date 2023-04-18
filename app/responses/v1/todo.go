package responses_v1

import "todo-list/app/models"

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func TodoMapToResponse(data *models.Todo) *Todo {
	return &Todo{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		IsDone:      data.IsDone,
	}
}
