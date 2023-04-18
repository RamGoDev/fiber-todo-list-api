package controllers_v1

import "github.com/gofiber/fiber/v2"

func Test(c *fiber.Ctx) error {
	return c.SendString("V1 Test")
}

func TestPost(c *fiber.Ctx) error {
	return c.SendString("V1 Test POST")
}
