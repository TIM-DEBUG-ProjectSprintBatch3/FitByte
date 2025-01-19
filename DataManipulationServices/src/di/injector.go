package di

import (
	"github.com/jackc/pgx/v5/pgxpool"
	authJwt "github.com/rafitanujaya/go-fiber-template/src/auth/jwt"
	"github.com/rafitanujaya/go-fiber-template/src/database/postgre"
	activityController "github.com/rafitanujaya/go-fiber-template/src/http/controllers/activity"
	userController "github.com/rafitanujaya/go-fiber-template/src/http/controllers/user"
	loggerZap "github.com/rafitanujaya/go-fiber-template/src/logger/zap"
	activityRepository "github.com/rafitanujaya/go-fiber-template/src/repositories/activity"
	userRepository "github.com/rafitanujaya/go-fiber-template/src/repositories/user"
	activityService "github.com/rafitanujaya/go-fiber-template/src/services/activity"
	userService "github.com/rafitanujaya/go-fiber-template/src/services/user"
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
