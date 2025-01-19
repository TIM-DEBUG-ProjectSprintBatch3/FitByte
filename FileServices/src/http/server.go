package httpServer

import (
	"fmt"

	"github.com/TimDebug/FitByte/src/config"
	"github.com/TimDebug/FitByte/src/di"
	fileController "github.com/TimDebug/FitByte/src/http/controllers/file"
	"github.com/TimDebug/FitByte/src/http/routes"
	fileRoutes "github.com/TimDebug/FitByte/src/http/routes/file"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type HttpServer struct{}

func (s *HttpServer) Listen() {
	app := fiber.New(fiber.Config{
		ServerHeader: "TIM-DEBUG",
	})

	//? Depedency Injection
	fileController := do.MustInvoke[fileController.FileController](di.Injector)

	routes := routes.SetRoutes(app)
	fileRoutes.SetRouteUsers(routes, fileController)

	app.Listen(fmt.Sprintf("%s:%s", "0.0.0.0", config.GetPort()))
}
