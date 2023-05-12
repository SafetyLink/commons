package jwt

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Config struct {
	Unauthorized fiber.Handler
}

func unauthorized(c *fiber.Ctx) error {
	c.Set(`Content-Type`, `application/json; charset=utf-8`)
	c.Status(fiber.StatusUnauthorized)
	return c.SendString(`{
    "data": null,
    "error": "Unauthorized",
    "message": null
}`)
}

func New(cfg Config) fiber.Handler {

	return func(c *fiber.Ctx) error {
		authorizationBearer := c.Get("authorization")

		claims, err := VerifyJwt(strings.Split(authorizationBearer, " ")[1])
		if err == nil {
			c.Locals("userID", claims.UserId)
			return c.Next()
		}

		return unauthorized(c)
	}
}
