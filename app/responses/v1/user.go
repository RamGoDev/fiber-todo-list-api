package responses_v1

import "todo-list/app/models"

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserMapToResponse(data *models.User) *User {
	return &User{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
}
