package responses_v1

import "todo-list/app/models"

type Login struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
	Type  string `json:"type" default:"Bearer"`
}

type Register struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func LoginMapToResponse(user *User, token string) *Login {
	return &Login{
		User:  user,
		Token: token,
		Type:  "Bearer",
	}
}

func RegisterMapToResponse(data *models.User) *User {
	return &User{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
}
