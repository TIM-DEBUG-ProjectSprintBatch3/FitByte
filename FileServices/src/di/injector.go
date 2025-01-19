package di

import (
	"fmt"
	"os"

	authJwt "github.com/TimDebug/FitByte/src/auth/jwt"
	"github.com/TimDebug/FitByte/src/database/postgre"
	fileController "github.com/TimDebug/FitByte/src/http/controllers/file"
	"github.com/TimDebug/FitByte/src/infrastructure/domain"
	"github.com/TimDebug/FitByte/src/infrastructure/storage"
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

	envMode := os.Getenv("MODE")
	fmt.Printf("Mode :%s", envMode)
	if envMode == "DEBUG" {
		do.Provide[domain.StorageClient](Injector, storage.NewMockStorageClientInject)
	} else {
		do.Provide[domain.StorageClient](Injector, storage.NewS3StorageClientInject)
	}

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
