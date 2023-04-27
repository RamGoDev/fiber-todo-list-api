package repositories_v1

import (
	"encoding/json"
	"errors"
	"todo-list/app/indices"
	"todo-list/app/models"
	responses_v1 "todo-list/app/responses/v1"
	validators_v1 "todo-list/app/validators/v1"
	"todo-list/database"
	"todo-list/helpers"

	"github.com/gofiber/fiber/v2"
)

type Auth interface {
	Login(c *fiber.Ctx) (*responses_v1.Login, error)
	Register(c *fiber.Ctx) (*responses_v1.User, error)
}

type authImpl struct {
	elastic database.Elasticsearch
}

func NewAuth(elastic database.Elasticsearch) Auth {
	return &authImpl{elastic}
}

func (impl authImpl) Login(c *fiber.Ctx) (*responses_v1.Login, error) {
	var err error
	var user *models.User
	var structure validators_v1.Login

	err = json.Unmarshal(c.Body(), &structure)

	if err != nil {
		return nil, errors.New("failed to convert login structure")
	}

	db := database.DB
	err = db.Find(&user, "email = ?", structure.Email).Error

	if err != nil {
		return nil, err
	}

	errHash := helpers.CompareHash(user.Password, structure.Password)
	if !errHash {
		return nil, errors.New("credentials is invalid")
	}

	user, jwt, errResp := helpers.GenerateJwt(user)

	if errResp != nil {
		return nil, errResp
	}

	userResp := responses_v1.UserMapToResponse(user)
	resp := responses_v1.LoginMapToResponse(userResp, jwt)

	return resp, nil
}

func (impl authImpl) Register(c *fiber.Ctx) (*responses_v1.User, error) {
	var err error
	var user models.User
	var structure validators_v1.Register
	var userIndex *indices.User

	err = json.Unmarshal(c.Body(), &structure)

	if err != nil {
		return nil, errors.New("failed to convert register structure")
	}

	db := database.DB
	err = c.BodyParser(&user)

	if err != nil {
		return nil, err
	}

	err = db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	// Store to elasticsearch
	_, _ = helpers.ConvertToOtherStruct(user, &userIndex)
	dataByte, _ := json.Marshal(userIndex)
	err = impl.elastic.AddDocument(userIndex.IndexName(), dataByte)
	if err != nil {
		return nil, err
	}

	resp := responses_v1.RegisterMapToResponse(&user)

	return resp, nil
}
