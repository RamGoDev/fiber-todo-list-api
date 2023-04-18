package repositories_v1

import (
	"errors"
	filters_v1 "todo-list/app/filters/v1"
	"todo-list/app/models"
	queries_v1 "todo-list/app/queries/v1"
	responses_v1 "todo-list/app/responses/v1"
	"todo-list/database"
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User interface {
	List(filter map[string]interface{}, pagination helpers.Pagination) *helpers.Pagination
	Show(id string) (*responses_v1.User, error)
	Store(c *fiber.Ctx) (*responses_v1.User, error)
	Update(c *fiber.Ctx, id string) (*responses_v1.User, error)
	Destroy(c *fiber.Ctx, id string) (*responses_v1.User, error)
	ForceDestroy(c *fiber.Ctx, id string) (*responses_v1.User, error)
}

type userImpl struct {
	//
}

func NewUser() User {
	return &userImpl{}
}

func IsExistsUser(query *gorm.DB, user *models.User) (*models.User, string) {
	err := query.Find(&user).Error

	if err != nil {
		return nil, err.Error()
	}

	if user.ID == 0 {
		return nil, "not found"
	}

	return user, ""
}

func (impl userImpl) List(filter map[string]interface{}, pagination helpers.Pagination) *helpers.Pagination {
	db := database.DB
	var users []models.User

	db = filters_v1.ByName(filter, db)
	db = filters_v1.ByEmail(filter, db)
	db = filters_v1.ById(filter, db)

	db.Scopes(helpers.Paginate(users, &pagination, db)).Find(&users)

	// TODO: simplify mapping/parse response
	// Mapping response
	resp := []responses_v1.User{}
	for _, v := range users {
		data := responses_v1.UserMapToResponse(&v)
		resp = append(resp, *data)
	}
	pagination.Rows = resp

	return &pagination
}

func (impl userImpl) Show(id string) (*responses_v1.User, error) {
	db := database.DB
	var user models.User

	db = queries_v1.ById(id, db)

	model, message := IsExistsUser(db, &user)
	if model == nil {
		return nil, errors.New(message)
	}

	resp := responses_v1.UserMapToResponse(&user)

	return resp, nil
}

func (impl userImpl) Store(c *fiber.Ctx) (*responses_v1.User, error) {
	db := database.DB
	var user models.User

	err := c.BodyParser(&user)

	if err != nil {
		return nil, err
	}

	err = db.Create(&user).Error

	resp := responses_v1.UserMapToResponse(&user)

	return resp, err
}

func (impl userImpl) Update(c *fiber.Ctx, id string) (*responses_v1.User, error) {
	db := database.DB
	var user models.User

	db = queries_v1.ById(id, db)

	model, message := IsExistsUser(db, &user)
	if model == nil {
		return nil, errors.New(message)
	}

	err := c.BodyParser(&user)

	if err != nil {
		return nil, err
	}

	db.Model(&user).Where("id = ?", id).UpdateColumns(&user)

	resp := responses_v1.UserMapToResponse(&user)

	return resp, err
}

func (impl userImpl) Destroy(c *fiber.Ctx, id string) (*responses_v1.User, error) {
	db := database.DB
	var user models.User

	db = queries_v1.ById(id, db)

	model, message := IsExistsUser(db, &user)
	if model == nil {
		return nil, errors.New(message)
	}

	err := db.Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (impl userImpl) ForceDestroy(c *fiber.Ctx, id string) (*responses_v1.User, error) {
	db := database.DB
	var user models.User

	db = queries_v1.ById(id, db)
	db = db.Unscoped()

	model, message := IsExistsUser(db, &user)
	if model == nil {
		return nil, errors.New(message)
	}

	err := db.Delete(&user).Error
	if err != nil {
		return nil, err
	}

	return nil, err
}
