package controllers_v1

import (
	repositories_v1 "todo-list/app/repositories/v1"
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
)

type Auth interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	MyProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
}

type authImpl struct {
	response       helpers.Response
	repository     repositories_v1.Auth
	userRepository repositories_v1.User
}

func NewAuth(
	response helpers.Response,
	repository repositories_v1.Auth,
	userRepository repositories_v1.User) Auth {
	return &authImpl{response, repository, userRepository}
}

// @Summary			Login
// @Description		Login with Email and Password
// @Tags			Auth
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			loginRequest	body	validators_v1.Login	true	"email"
// @Router 			/api/v1/auth/login	[post]
func (impl authImpl) Login(c *fiber.Ctx) error {
	resp, err := impl.repository.Login(c)

	if err != nil {
		return impl.response.UnprocessableEntity(c, nil, err.Error())
	}

	return impl.response.Success(c, resp, "login successfully")
}

// @Summary			Register
// @Description		Register new user
// @Tags			Auth
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			registerRequest	body	validators_v1.Register	true	"email"
// @Router 			/api/v1/auth/register	[post]
func (impl authImpl) Register(c *fiber.Ctx) error {
	resp, err := impl.repository.Register(c)

	if err != nil {
		return impl.response.UnprocessableEntity(c, nil, err.Error())
	}

	return impl.response.Success(c, resp, "register successfully")
}

// @Security		BearerAuth
// @Summary			Get My Profile
// @Description		Get profile of current user
// @Tags			Auth
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Router 			/api/v1/auth/profile	[get]
func (impl authImpl) MyProfile(c *fiber.Ctx) error {
	uid := helpers.GetCurrentUserId(c)
	user, err := impl.userRepository.Show(uid)

	if err != nil {
		return impl.response.NotFound(c, err.Error())
	}

	return impl.response.Success(c, user, "success")
}

// @Security		BearerAuth
// @Summary			Update Profile
// @Description		Update profile of current user
// @Tags			Auth
// @Accept			application/json
// @Produce 		json
// @Success			200	{object}  helpers.Success
// @Failure			400 {object}  helpers.Failed
// @Failure     	404 {object}  helpers.Error
// @Failure     	500 {object}  helpers.Error
// @Param			profileRequest	body	validators_v1.Profile	true	"name"
// @Router 			/api/v1/auth/profile	[put]
func (impl authImpl) UpdateProfile(c *fiber.Ctx) error {
	uid := helpers.GetCurrentUserId(c)
	user, err := impl.userRepository.Update(c, uid)

	if err != nil {
		return impl.response.NotFound(c, err.Error())
	}

	return impl.response.Success(c, user, "success")
}
