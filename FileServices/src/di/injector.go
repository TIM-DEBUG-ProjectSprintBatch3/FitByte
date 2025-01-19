package di

import (
	authJwt "github.com/TimDebug/FitByte/src/auth/jwt"
	"github.com/TimDebug/FitByte/src/database/postgre"
	fileController "github.com/TimDebug/FitByte/src/http/controllers/file"
	loggerZap "github.com/TimDebug/FitByte/src/logger/zap"
	fileService "github.com/TimDebug/FitByte/src/services/file"
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

	//? Setup Services
	do.Provide[fileService.FileServiceInterface](Injector, fileService.NewInject)

	//? Setup Controller/Handler
	do.Provide[fileController.FileController](Injector, fileController.NewInject)
}
