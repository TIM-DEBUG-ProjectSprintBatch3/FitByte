package userroutes

import (
	"github.com/gofiber/fiber/v2"
	activityController "github.com/rafitanujaya/go-fiber-template/src/http/controllers/activity"
	"github.com/rafitanujaya/go-fiber-template/src/http/middlewares"
)

func SetRouteActivities(router fiber.Router, uc activityController.ActivityControllerInterface) {
	router.Get("/activity", middlewares.AuthMiddleware, uc.GetAll)
}
