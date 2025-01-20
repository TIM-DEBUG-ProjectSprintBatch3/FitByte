package fileRoutes

import (
	fileController "github.com/TimDebug/FitByte/src/http/controllers/file"
	"github.com/TimDebug/FitByte/src/http/controllers/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetRouteUsers(router fiber.Router, controller fileController.FileController) {
	router.Post("/file", middlewares.AuthMiddleware, controller.Upload)
}
