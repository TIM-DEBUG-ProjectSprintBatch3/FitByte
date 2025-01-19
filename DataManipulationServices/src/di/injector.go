package di

import (
	authJwt "github.com/TimDebug/FitByte/src/auth/jwt"
	"github.com/TimDebug/FitByte/src/database/postgre"
	activityController "github.com/TimDebug/FitByte/src/http/controllers/activity"
	userController "github.com/TimDebug/FitByte/src/http/controllers/user"
	loggerZap "github.com/TimDebug/FitByte/src/logger/zap"
	activityRepository "github.com/TimDebug/FitByte/src/repositories/activity"
	userRepository "github.com/TimDebug/FitByte/src/repositories/user"
	activityService "github.com/TimDebug/FitByte/src/services/activity"
	userService "github.com/TimDebug/FitByte/src/services/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

var Injector *do.RootScope

func init() {
	Injector = do.New()

	//? Setup Database Connection
	do.Provide[*pgxpool.Pool](Injector, postgre.NewPgxConnectInject)

	//? Logger
	//? Zap
	do.Provide[loggerZap.LoggerInterface](Injector, loggerZap.NewLogHandlerInject)

	//? Setup Auth
	//? JWT Service
	do.Provide[authJwt.JwtServiceInterface](Injector, authJwt.NewJwtServiceInject)

	//? Setup Repositories
	//? User Repository
	do.Provide[userRepository.UserRepositoryInterface](Injector, userRepository.NewUserRepositoryInject)

	//? Setup Services
	//? User Service
	do.Provide[userService.UserServiceInterface](Injector, userService.NewUserServiceInject)

	//? Setup Controller/Handler
	//? User Controller
	do.Provide[userController.UserControllerInterface](Injector, userController.NewUserControllerInject)

	//? Setup Repositories
	//? Activity Repository
	do.Provide[activityRepository.ActivityRepositoryInterface](Injector, activityRepository.NewActivityRepositoryInject)

	//? Setup Services
	//? Activity Service
	do.Provide[activityService.ActivityServiceInterface](Injector, activityService.NewActivityServiceInject)

	//? Setup Controller/Handler
	//? Activity Controller
	do.Provide[activityController.ActivityControllerInterface](Injector, activityController.NewActivityControllerInject)
}
