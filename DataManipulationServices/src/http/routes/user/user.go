package userroutes

import (
	userController "github.com/TimDebug/FitByte/src/http/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetRouteUsers(router fiber.Router, uc userController.UserControllerInterface) {
	router.Post("/login", uc.Login)
	router.Post("/register", uc.Register)
	router.Patch("/user", uc.Update)
}
