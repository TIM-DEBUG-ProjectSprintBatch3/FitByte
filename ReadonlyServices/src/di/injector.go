package di

import (
	authJwt "github.com/TimDebug/FitByte/src/auth/jwt"
	"github.com/TimDebug/FitByte/src/database/postgre"
	"github.com/TimDebug/FitByte/src/http/controllers"
	activityController "github.com/TimDebug/FitByte/src/http/controllers/activity"
	loggerZap "github.com/TimDebug/FitByte/src/logger/zap"
	"github.com/TimDebug/FitByte/src/repositories"
	activityRepository "github.com/TimDebug/FitByte/src/repositories/activity"
	"github.com/TimDebug/FitByte/src/services"
	activityService "github.com/TimDebug/FitByte/src/services/activity"
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
	//? Activity Repository
	do.Provide[activityRepository.ActivityRepositoryInterface](Injector, activityRepository.NewActivityRepositoryInject)
	do.Provide[repositories.UserRepository](Injector, repositories.NewUserRepositoryImplInject)

	//? Setup Services
	//? Activity Service
	do.Provide[activityService.ActivityServiceInterface](Injector, activityService.NewActivityServiceInject)
	do.Provide[services.UserService](Injector, services.NewUserServiceImplInject)

	//? Setup Controller/Handler
	//? Activity Controller
	do.Provide[activityController.ActivityControllerInterface](Injector, activityController.NewActivityControllerInject)
	do.Provide[controllers.UserController](Injector, controllers.NewUserControllerInject)
}
