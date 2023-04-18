package controllers_v1

import (
	"strconv"
	repositories_v1 "todo-list/app/repositories/v1"
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
)

type User interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Destroy(c *fiber.Ctx) error
	ForceDestroy(c *fiber.Ctx) error
}

type userImpl struct {
	response   helpers.Response
	repository repositories_v1.User
}

func NewUser(
	response helpers.Response,
	repository repositories_v1.User) User {
	return &userImpl{response, repository}
}

// @Security 		BearerAuth
// @Summary			List User
// @Description		List User
// @Tags			User
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	query	int	false	"User ID"
// @Param			name	query	string	false	"Name"
// @Param			email	query	string	false	"Email"
// @Param			limit	query	int	false	"Default 10"	default(10)
// @Param			page	query	int	false	"Default 10"	default(1)
// @Param			sort	query	string	false	"Sorting"	Enums(ID asc, ID desc, name asc, name desc)
// @Router 			/api/v1/users	[get]
func (impl userImpl) Index(c *fiber.Ctx) error {
	var pagination helpers.Pagination

	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))
	pagination.Sort = c.Query("sort")
	filter := map[string]interface{}{
		"name":  c.Query("name"),
		"email": c.Query("email"),
		"id":    c.Query("id"),
	}
	users := impl.repository.List(filter, pagination)

	return impl.response.Success(c, users, "success")
}

// @Security		BearerAuth
// @Summary			Detail User
// @Description		Detail User
// @Tags			User
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"User ID"
// @Router 			/api/v1/users/{id}	[get]
func (impl userImpl) Show(c *fiber.Ctx) error {
	user, err := impl.repository.Show(c.Params("id"))

	if err != nil {
		return impl.response.NotFound(c, err.Error())
	}

	return impl.response.Success(c, user, "success")
}

// @Security		BearerAuth
// @Summary			Create new User
// @Description		Create new User
// @Tags			User
// @Accept			application/json
// @Produce 		json
// @Success			201	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			usersRequest	body	validators_v1.User	true	"name"
// @Router 			/api/v1/users	[post]
func (impl userImpl) Store(c *fiber.Ctx) error {
	user, err := impl.repository.Store(c)

	if err != nil {
		return impl.response.UnprocessableEntity(c, user, err.Error())
	}

	return impl.response.Success(c, user, "success")
}

// @Security		BearerAuth
// @Summary			Update User
// @Description		Update User
// @Tags			User
// @Accept			application/json
// @Produce 		json
// @Success			201	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"User ID"
// @Param			usersRequest	body	validators_v1.User	true	"name"
// @Router 			/api/v1/users/{id}	[put]
func (impl userImpl) Update(c *fiber.Ctx) error {
	user, err := impl.repository.Update(c, c.Params("id"))

	if err != nil {
		return impl.response.UnprocessableEntity(c, user, err.Error())
	}

	return impl.response.Success(c, user, "success")
}

// @Security		BearerAuth
// @Summary			Delete User
// @Description		Delete User with ID
// @Tags			User
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"User ID"
// @Router 			/api/v1/users/{id}	[delete]
func (impl userImpl) Destroy(c *fiber.Ctx) error {
	user, err := impl.repository.Destroy(c, c.Params("id"))

	if err != nil {
		return impl.response.UnprocessableEntity(c, user, err.Error())
	}

	return impl.response.Success(c, user, "Delete successfully")
}

// @Security		BearerAuth
// @Summary			Force Delete User
// @Description		Force Delete User with ID
// @Tags			User
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			id	path	int	true	"User ID"
// @Router 			/api/v1/users/{id}/force	[delete]
func (impl userImpl) ForceDestroy(c *fiber.Ctx) error {
	user, err := impl.repository.ForceDestroy(c, c.Params("id"))

	if err != nil {
		return impl.response.UnprocessableEntity(c, user, err.Error())
	}

	return impl.response.Success(c, user, "Force Delete successfully")
}
