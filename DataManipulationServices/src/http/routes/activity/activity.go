package activityroutes

import (
	"github.com/gofiber/fiber/v2"
	activityController "github.com/rafitanujaya/go-fiber-template/src/http/controllers/activity"
	"github.com/rafitanujaya/go-fiber-template/src/http/middlewares"
)

func SetRouteActivities(router fiber.Router, ac activityController.ActivityControllerInterface) {
	router.Post("/activity", middlewares.AuthMiddleware, middlewares.ContentTypeJsonApplicationMiddleware, ac.Create)
}
