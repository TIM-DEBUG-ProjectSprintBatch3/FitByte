package userroutes

import (
	userController "github.com/TimDebug/FitByte/src/http/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetRouteUsers(router fiber.Router, uc userController.UserControllerInterface) {
	router.Post("/register", uc.Register)
}
