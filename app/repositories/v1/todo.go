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

type Todo interface {
	List(filter map[string]interface{}, pagination helpers.Pagination) *helpers.Pagination
	Show(c *fiber.Ctx, id string) (*responses_v1.Todo, error)
	Store(c *fiber.Ctx) (*responses_v1.Todo, error)
	Update(c *fiber.Ctx, id string) (*responses_v1.Todo, error)
	Destroy(c *fiber.Ctx, id string) (*responses_v1.Todo, error)
	ForceDestroy(c *fiber.Ctx, id string) (*responses_v1.Todo, error)
	Complete(c *fiber.Ctx, id string) (*responses_v1.Todo, string, error)
}

type todoImpl struct {
	cache database.CacheDriver
}

func NewTodo(cache database.CacheDriver) Todo {
	return &todoImpl{cache}
}

func IsExistsTodo(query *gorm.DB, todo *models.Todo) (*models.Todo, string) {
	err := query.Find(&todo).Error

	if err != nil {
		return nil, err.Error()
	}

	if todo.ID == 0 {
		return nil, "not found"
	}

	return todo, ""
}

func (impl todoImpl) List(filter map[string]interface{}, pagination helpers.Pagination) *helpers.Pagination {
	var todos []models.Todo
	var todo models.Todo
	db := database.DB

	queryString := helpers.ConvertToQueryString(filter)
	cacheKey := fmt.Sprintf("%s_%s", todo.CacheBaseKey(), queryString)

	// Checking the cache
	cache := impl.cache.Get(cacheKey)
	if cache != "" {
		json.Unmarshal([]byte(cache), &pagination)
		fmt.Println("resp from the cache")
		return &pagination
	}

	db = filters_v1.ByUserId(filter, db)
	db = filters_v1.ByTitle(filter, db)
	db = filters_v1.ByDescription(filter, db)
	db = filters_v1.ByIsDone(filter, db)

	db.Scopes(helpers.Paginate(todos, &pagination, db)).Find(&todos)

	resp := []responses_v1.Todo{}
	for _, v := range todos {
		data := responses_v1.TodoMapToResponse(&v)
		resp = append(resp, *data)
	}
	pagination.Rows = resp

	// Set the cache
	respString, _ := json.Marshal(pagination)
	impl.cache.Set(cacheKey, string(respString))

	return &pagination
}

func (impl todoImpl) Show(c *fiber.Ctx, id string) (*responses_v1.Todo, error) {
	db := database.DB
	var todo models.Todo

	userId := helpers.GetCurrentUserId(c)

	// Checking the cache
	cacheKey := todo.CacheShowKey(userId, id)
	cache := impl.cache.Get(cacheKey)
	if cache != "" {
		resp := responses_v1.TodoMapToResponse(&todo)
		json.Unmarshal([]byte(cache), resp)
		fmt.Println("resp from the cache")
		return resp, nil
	}

	db = queries_v1.ByUserId(userId, db)
	db = queries_v1.ById(id, db)

	model, message := IsExistsTodo(db, &todo)
	if model == nil {
		return nil, errors.New(message)
	}

	resp := responses_v1.TodoMapToResponse(&todo)

	// Set the cache
	respString, _ := json.Marshal(resp)
	impl.cache.Set(cacheKey, string(respString))

	return resp, nil
}

func (impl todoImpl) Store(c *fiber.Ctx) (*responses_v1.Todo, error) {
	var user models.User
	var todo models.Todo
	db := database.DB
	userId := helpers.GetCurrentUserId(c)

	err := c.BodyParser(&todo)

	if err != nil {
		return nil, err
	}

	err = db.Find(&user, "id = ?", userId).Error

	if err != nil {
		return nil, err
	}

	todo.UserId = user.ID

	err = db.Create(&todo).Error

	if err != nil {
		return nil, err
	}

	resp := responses_v1.TodoMapToResponse(&todo)

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", todo.CacheBaseKey())
	impl.cache.Clear(cachePattern)

	return resp, err
}

func (impl todoImpl) Update(c *fiber.Ctx, id string) (*responses_v1.Todo, error) {
	db := database.DB
	var todo models.Todo

	userId := helpers.GetCurrentUserId(c)
	db = queries_v1.ByUserId(userId, db)
	db = queries_v1.ById(id, db)

	model, message := IsExistsTodo(db, &todo)
	if model == nil {
		return nil, errors.New(message)
	}

	err := c.BodyParser(&todo)

	if err != nil {
		return nil, err
	}

	db.Model(&todo).Where("id = ?", id).UpdateColumns(&todo)

	resp := responses_v1.TodoMapToResponse(&todo)

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", todo.CacheBaseKey())
	impl.cache.Clear(cachePattern)

	return resp, err
}

func (impl todoImpl) Destroy(c *fiber.Ctx, id string) (*responses_v1.Todo, error) {
	db := database.DB
	var todo models.Todo

	userId := helpers.GetCurrentUserId(c)
	db = queries_v1.ByUserId(userId, db)
	db = queries_v1.ById(id, db)

	model, message := IsExistsTodo(db, &todo)
	if model == nil {
		return nil, errors.New(message)
	}

	err := db.Delete(&todo).Error
	if err != nil {
		return nil, err
	}

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", todo.CacheBaseKey())
	impl.cache.Clear(cachePattern)

	return nil, err
}

func (impl todoImpl) ForceDestroy(c *fiber.Ctx, id string) (*responses_v1.Todo, error) {
	db := database.DB
	var todo models.Todo

	// .Unscoped(): find with soft delete data
	userId := helpers.GetCurrentUserId(c)
	db = queries_v1.ByUserId(userId, db)
	db = queries_v1.ById(id, db)
	db = db.Unscoped()

	model, message := IsExistsTodo(db, &todo)
	if model == nil {
		return nil, errors.New(message)
	}

	err := db.Delete(&todo).Error
	if err != nil {
		return nil, err
	}

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", todo.CacheBaseKey())
	impl.cache.Clear(cachePattern)

	return nil, err
}

func (impl todoImpl) Complete(c *fiber.Ctx, id string) (*responses_v1.Todo, string, error) {
	db := database.DB
	var todo models.Todo

	userId := helpers.GetCurrentUserId(c)
	db = queries_v1.ByUserId(userId, db)
	db = queries_v1.ById(id, db)

	model, message := IsExistsTodo(db, &todo)
	if model == nil {
		return nil, message, errors.New(message)
	}

	status := true
	msgStatus := "Completed"
	if todo.IsDone {
		status = false
		msgStatus = "Uncompleted"
	}
	todo.IsDone = status
	db.Model(&todo).Update("is_done", status)
	msgStatus = msgStatus + " is successfully"

	resp := responses_v1.TodoMapToResponse(&todo)

	// Clear cache
	cachePattern := fmt.Sprintf("*%s*", todo.CacheBaseKey())
	impl.cache.Clear(cachePattern)

	return resp, msgStatus, nil
}
