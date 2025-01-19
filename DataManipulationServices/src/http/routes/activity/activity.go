package activityroutes

import (
	activityController "github.com/TimDebug/FitByte/src/http/controllers/activity"
	"github.com/TimDebug/FitByte/src/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetRouteActivities(router fiber.Router, ac activityController.ActivityControllerInterface) {
	router.Post("/activity", middlewares.AuthMiddleware, middlewares.ContentTypeJsonApplicationMiddleware, ac.Create)
}
