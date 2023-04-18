package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUserId(c *fiber.Ctx) string {
	uid := c.Locals("user_id")
	if uid == nil {
		return ""
	}
	return fmt.Sprintf("%#v", uid)
}
