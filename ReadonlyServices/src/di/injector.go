package di

import (
	"github.com/jackc/pgx/v5/pgxpool"
	authJwt "github.com/rafitanujaya/go-fiber-template/src/auth/jwt"
	"github.com/rafitanujaya/go-fiber-template/src/database/postgre"
	activityController "github.com/rafitanujaya/go-fiber-template/src/http/controllers/activity"
	loggerZap "github.com/rafitanujaya/go-fiber-template/src/logger/zap"
	activityRepository "github.com/rafitanujaya/go-fiber-template/src/repositories/activity"
	activityService "github.com/rafitanujaya/go-fiber-template/src/services/activity"
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

	//? Setup Services
	//? Activity Service
	do.Provide[activityService.ActivityServiceInterface](Injector, activityService.NewActivityServiceInject)

	//? Setup Controller/Handler
	//? Activity Controller
	do.Provide[activityController.ActivityControllerInterface](Injector, activityController.NewActivityControllerInject)
}
