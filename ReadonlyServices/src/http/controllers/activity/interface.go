package activityController

import "github.com/gofiber/fiber/v2"

type ActivityControllerInterface interface {
	GetAll(C *fiber.Ctx) error
}
