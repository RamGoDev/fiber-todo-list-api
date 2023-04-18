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
	//
}

func NewTodo() Todo {
	return &todoImpl{}
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
	db := database.DB

	db = filters_v1.ByUserId(filter, db)
	db = filters_v1.ByTitle(filter, db)
	db = filters_v1.ByDescription(filter, db)
	db = filters_v1.ByIsDone(filter, db)

	db.Scopes(helpers.Paginate(todos, &pagination, db)).Find(&todos)

	// TODO: simplify mapping/parse response
	// Mapping response
	resp := []responses_v1.Todo{}
	for _, v := range todos {
		data := responses_v1.TodoMapToResponse(&v)
		resp = append(resp, *data)
	}
	pagination.Rows = resp

	return &pagination
}

func (impl todoImpl) Show(c *fiber.Ctx, id string) (*responses_v1.Todo, error) {
	db := database.DB
	var todo models.Todo

	userId := helpers.GetCurrentUserId(c)
	db = queries_v1.ByUserId(userId, db)
	db = queries_v1.ById(id, db)

	model, message := IsExistsTodo(db, &todo)
	if model == nil {
		return nil, errors.New(message)
	}

	// TODO: simplify mapping/parse response
	// Mapping response
	resp := responses_v1.TodoMapToResponse(&todo)

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

	// TODO: simplify mapping/parse response
	// Mapping response
	resp := responses_v1.TodoMapToResponse(&todo)

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

	// TODO: simplify mapping/parse response
	// Mapping response
	resp := responses_v1.TodoMapToResponse(&todo)

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

	// TODO: simplify mapping/parse response
	// Mapping response
	resp := responses_v1.TodoMapToResponse(&todo)

	return resp, msgStatus, nil
}
