package validators_v1

import (
	"todo-list/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Title       string `json:"title" default:"title" validate:"required"`
	Description string `json:"description" default:"description" validate:"required"`
	IsDone      bool   `json:"is_done" default:"false" validate:"boolean"`
}

func TodoValidator(c *fiber.Ctx) error {
	var errors []*IError
	body := new(Todo)
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
