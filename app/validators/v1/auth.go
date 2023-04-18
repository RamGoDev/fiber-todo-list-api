package validators_v1

import (
	"todo-list/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Login struct {
	Email    string `json:"email" default:"super@todo.id" validate:"required"`
	Password string `json:"password" default:"password123" validate:"required"`
}

type Register struct {
	Name     string `json:"name" default:"New user 1" validate:"required"`
	Email    string `json:"email" default:"newuser1@todo.id" validate:"required,email"` // FIXME: add unique validation
	Password string `json:"password" default:"password123" validate:"required"`
}

type Profile struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

func LoginValidator(c *fiber.Ctx) error {
	var errors []*IError
	body := new(Login)
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

func RegisterValidator(c *fiber.Ctx) error {
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

func ProfileValidator(c *fiber.Ctx) error {
	var errors []*IError
	body := new(Profile)
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
