package httpServer

import (
	"github.com/TimDebug/FitByte/src/config"
	"github.com/TimDebug/FitByte/src/di"
	userController "github.com/TimDebug/FitByte/src/http/controllers/user"
	"github.com/TimDebug/FitByte/src/http/routes"
	userroutes "github.com/TimDebug/FitByte/src/http/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type HttpServer struct{}

func (s *HttpServer) Listen() {
	app := fiber.New(fiber.Config{
		ServerHeader: "TIM-DEBUG",
	})

	//? Depedency Injection
	//? UserController
	uc := do.MustInvoke[userController.UserControllerInterface](di.Injector)

	routes := routes.SetRoutes(app)
	userroutes.SetRouteUsers(routes, uc)
	app.Listen(config.GetPort())
}
