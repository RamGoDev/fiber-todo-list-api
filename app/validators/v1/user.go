package validators_v1

import (
	"todo-list/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name     string `json:"name" default:"New user 1" validate:"required"`
	Email    string `json:"email" default:"newuser1@todo.id" validate:"required,email"` // FIXME: add unique validation
	Password string `json:"password" default:"password123" validate:"required"`
}

func UserValidator(c *fiber.Ctx) error {
	var errors []*IError
	body := new(Register)
	err := c.BodyParser(body)
	response := helpers.NewResponse()

	if err != nil {
		return response.UnprocessableEntity(c, nil, err.Error())
	}

	err = Validator.Struct(body)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return response.UnprocessableEntity(c, errors, "unprocessable entity")
	}

	return c.Next()
}
