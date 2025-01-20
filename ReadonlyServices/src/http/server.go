package httpServer

import (
	"fmt"

	"github.com/TimDebug/FitByte/src/config"
	"github.com/TimDebug/FitByte/src/di"
	activityController "github.com/TimDebug/FitByte/src/http/controllers/activity"
	controllers "github.com/TimDebug/FitByte/src/http/controllers/userProfile"
	"github.com/TimDebug/FitByte/src/http/routes"
	activityroutes "github.com/TimDebug/FitByte/src/http/routes/activity"
	userroutes "github.com/TimDebug/FitByte/src/http/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type HttpServer struct{}

func (s *HttpServer) Listen() {
	fmt.Printf("New Fiber\n")
	app := fiber.New(fiber.Config{
		ServerHeader: "TIM-DEBUG",
	})

	fmt.Printf("Inject Controllers\n")
	//? Depedency Injection
	//? ActivityController
	ac := do.MustInvoke[activityController.ActivityControllerInterface](di.Injector)
	uc := do.MustInvoke[controllers.UserController](di.Injector)

	routes := routes.SetRoutes(app)
	activityroutes.SetRouteActivities(routes, ac)
	userroutes.SetRouteActivities(routes, uc)

	fmt.Printf("Start Listener\n")
	app.Listen(fmt.Sprintf("%s:%s", "0.0.0.0", config.GetPort()))
}
