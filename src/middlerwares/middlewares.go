package middlerwares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	if len(c.Request().Header.Peek("X-API-KEY")) > 0 {
		return c.Next()
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func Recover(c *fiber.Ctx) error {
	defer func(c *fiber.Ctx) error {
		if recover() != nil {
			log.Println("we got a panic")
			return c.SendStatus(500)
		}
		return nil
	}(c)
	return c.Next()
}
