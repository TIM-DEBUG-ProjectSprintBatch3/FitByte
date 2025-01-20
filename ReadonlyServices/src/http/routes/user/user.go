package userroutes

import (
	controllers "github.com/TimDebug/FitByte/src/http/controllers/userProfile"
	"github.com/TimDebug/FitByte/src/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetRouteActivities(router fiber.Router, uc controllers.UserController) {
	router.Get("/user", middlewares.AuthMiddleware, uc.GetProfile)
}
