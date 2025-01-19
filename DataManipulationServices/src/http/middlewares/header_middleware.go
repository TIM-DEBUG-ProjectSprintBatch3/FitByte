package middlewares

import (
	"github.com/gofiber/fiber/v2"
	response "github.com/rafitanujaya/go-fiber-template/src/model/web"
)

func ContentTypeJsonApplicationMiddleware(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if contentType != "application/json" && contentType != "application/json; charset=utf-8" {
		return c.Status(fiber.StatusBadRequest).JSON(response.Web{
			Message: "Content-Type must be application/json",
		})
	}

	return c.Next()
}
