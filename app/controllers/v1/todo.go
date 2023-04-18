package controllers_v1

import (
	"strconv"
	repositories_v1 "todo-list/app/repositories/v1"
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
)

type Todo interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
	ForceDestroy(c *fiber.Ctx) error
	Complete(c *fiber.Ctx) error
}

type todoImpl struct {
	response   helpers.Response
	repository repositories_v1.Todo
}

func NewTodo(
	response helpers.Response,
	repository repositories_v1.Todo) Todo {
	return &todoImpl{response, repository}
}

// @Security 		BearerAuth
// @Summary			List Todo
// @Description		List Todo of current user
// @Tags			Todo
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			title	query	string	false	"Title"
// @Param			description	query	string	false	"Description"
// @Param			is_done	query	bool	false	"Is Done"
// @Param			limit	query	int	false	"Default 10"	default(10)
// @Param			page	query	int	false	"Default 10"	default(1)
// @Param			sort	query	string	false	"Sorting"	Enums(ID asc, ID desc, title asc, title desc)
// @Router 			/api/v1/todos	[get]
func (impl todoImpl) Index(c *fiber.Ctx) error {
	var pagination helpers.Pagination

	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Sort = c.Query("sort")
	filter := map[string]interface{}{
		"title":       c.Query("title"),
		"description": c.Query("description"),
		"is_done":     c.Query("is_done"),
		"user_id":     helpers.GetCurrentUserId(c),
	}
	todos := impl.repository.List(filter, pagination)

	return impl.response.Success(c, todos, "success")
}

// @Security		BearerAuth
// @Summary			Detail Todo
// @Description		Detail Todo
// @Tags			Todo
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"Todo ID"
// @Router 			/api/v1/todos/{id}	[get]
func (impl todoImpl) Show(c *fiber.Ctx) error {
	todo, err := impl.repository.Show(c, c.Params("id"))

	if err != nil {
		return impl.response.NotFound(c, err.Error())
	}

	return impl.response.Success(c, todo, "success")
}

// @Security		BearerAuth
// @Summary			Create new Todo
// @Description		Create new Todo for current user
// @Tags			Todo
// @Accept			application/json
// @Produce 		json
// @Success			201	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			todosRequest	body	validators_v1.Todo	true	"title"
// @Router 			/api/v1/todos	[post]
func (impl todoImpl) Store(c *fiber.Ctx) error {
	todo, err := impl.repository.Store(c)

	if err != nil {
		return impl.response.UnprocessableEntity(c, todo, err.Error())
	}

	return impl.response.Success(c, todo, "success")
}

// @Security		BearerAuth
// @Summary			Update Todo
// @Description		Update Todo
// @Tags			Todo
// @Accept			application/json
// @Produce 		json
// @Success			201	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"Todo ID"
// @Param			todosRequest	body	validators_v1.Todo	true	"title"
// @Router 			/api/v1/todos/{id}	[put]
func (impl todoImpl) Update(c *fiber.Ctx) error {
	todo, err := impl.repository.Update(c, c.Params("id"))

	if err != nil {
		return impl.response.UnprocessableEntity(c, todo, err.Error())
	}

	return impl.response.Success(c, todo, "success")
}

// @Security		BearerAuth
// @Summary			Delete Todo
// @Description		Delete Todo with ID
// @Tags			Todo
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"Todo ID"
// @Router 			/api/v1/todos/{id}	[delete]
func (impl todoImpl) Destroy(c *fiber.Ctx) error {
	todo, err := impl.repository.Destroy(c, c.Params("id"))

	if err != nil {
		return impl.response.UnprocessableEntity(c, todo, err.Error())
	}

	return impl.response.Success(c, todo, "Delete successfully")
}

// @Security		BearerAuth
// @Summary			Force Delete Todo
// @Description		Force Delete Todo with ID
// @Tags			Todo
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"Todo ID"
// @Router 			/api/v1/todos/{id}/force	[delete]
func (impl todoImpl) ForceDestroy(c *fiber.Ctx) error {
	todo, err := impl.repository.ForceDestroy(c, c.Params("id"))

	if err != nil {
		return impl.response.UnprocessableEntity(c, todo, err.Error())
	}

	return impl.response.Success(c, todo, "Force Delete successfully")
}

// @Security		BearerAuth
// @Summary			Completed / Uncompleted Todo
// @Description		Completed / Uncompleted Todo with ID
// @Tags			Todo
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"Todo ID"
// @Router 			/api/v1/todos/{id}/complete	[put]
func (impl todoImpl) Complete(c *fiber.Ctx) error {
	todo, message, err := impl.repository.Complete(c, c.Params("id"))

	if err != nil {
		return impl.response.UnprocessableEntity(c, todo, err.Error())
	}

	return impl.response.Success(c, todo, message)
}
