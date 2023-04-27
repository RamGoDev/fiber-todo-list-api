package repositories_v1

import (
	"encoding/json"
	"errors"
	"fmt"
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
	cache   database.CacheDriver
	elastic database.Elasticsearch
}

func NewUser(cache database.CacheDriver,
	elastic database.Elasticsearch) User {
	return &userImpl{cache, elastic}
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
	var users []models.User
	var user models.User
	db := database.DB

	queryString := helpers.ConvertToQueryString(filter)
	cacheKey := fmt.Sprintf("%s_%s", user.CacheBaseKey(), queryString)

	// Checking the cache
	cache := impl.cache.Get(cacheKey)
	if cache != "" {
		json.Unmarshal([]byte(cache), &pagination)
		fmt.Println("resp from the cache")
		return &pagination
	}

	db = filters_v1.ByName(filter, db)
	db = filters_v1.ByEmail(filter, db)
	db = filters_v1.ById(filter, db)

	db.Scopes(helpers.Paginate(users, &pagination, db)).Find(&users)

	resp := []responses_v1.User{}
	for _, v := range users {
		data := responses_v1.UserMapToResponse(&v)
		resp = append(resp, *data)
	}
	pagination.Rows = resp

	// Set the cache
	respString, _ := json.Marshal(pagination)
	impl.cache.Set(cacheKey, string(respString))

	return &pagination
}

func (impl userImpl) Show(id string) (*responses_v1.User, error) {
	db := database.DB
	var user models.User

	// Checking the cache
	cacheKey := user.CacheShowKey(id)
	cache := impl.cache.Get(cacheKey)
	if cache != "" {
		resp := responses_v1.UserMapToResponse(&user)
		json.Unmarshal([]byte(cache), resp)
		fmt.Println("resp from the cache")
		return resp, nil
	}

	db = queries_v1.ById(id, db)

	model, message := IsExistsUser(db, &user)
	if model == nil {
		return nil, errors.New(message)
	}

	resp := responses_v1.UserMapToResponse(&user)

	// Set the cache
	respString, _ := json.Marshal(resp)
	impl.cache.Set(cacheKey, string(respString))

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

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", user.CacheBaseKey())
	impl.cache.Clear(cachePattern)

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

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", user.CacheBaseKey())
	impl.cache.Clear(cachePattern)

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

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", user.CacheBaseKey())
	impl.cache.Clear(cachePattern)

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

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", user.CacheBaseKey())
	impl.cache.Clear(cachePattern)

	return nil, err
}
