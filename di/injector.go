package di

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/levensspel/go-gin-template/database"
	"github.com/levensspel/go-gin-template/domain"
	authHandler "github.com/levensspel/go-gin-template/handler/auth"
	fileHandler "github.com/levensspel/go-gin-template/handler/file"
	userHandler "github.com/levensspel/go-gin-template/handler/user"

	"github.com/levensspel/go-gin-template/infrastructure/storage"
	"github.com/levensspel/go-gin-template/logger"

	fileService "github.com/levensspel/go-gin-template/service/file"
	userService "github.com/levensspel/go-gin-template/service/user"

	fileRepository "github.com/levensspel/go-gin-template/repository/file"
	userRepository "github.com/levensspel/go-gin-template/repository/user"

	"github.com/samber/do/v2"
)

var Injector *do.RootScope

func init() {
	Injector = do.New()

	// Jika ada dependensi, tolong tambahkan sesuai dengan hirarki
	// Setup client
	envMode := os.Getenv("MODE")
	fmt.Print("Mode :%s", envMode)
	if envMode == "DEBUG" {
		do.Provide[domain.StorageClient](Injector, storage.NewMockStorageClientInject)
	} else {
		do.Provide[domain.StorageClient](Injector, storage.NewS3StorageClientInject)
	}

	// Setup database connection
	do.Provide[*pgxpool.Pool](Injector, database.NewUserRepositoryInject)
	// setup logger
	do.Provide[logger.LogHandler](Injector, logger.NewlogHandlerInject)

	// Setup repositories
	// UserRepository
	do.Provide[userRepository.UserRepository](Injector, userRepository.NewUserRepositoryInject)
	do.Provide[fileRepository.FileRepository](Injector, fileRepository.NewInject)

	// Setup Services
	do.Provide[userService.UserService](Injector, userService.NewUserServiceInject)
	do.Provide[fileService.FileService](Injector, fileService.NewInject)

	// Setup Handlers
	do.Provide[userHandler.UserHandler](Injector, userHandler.NewUserHandlerInject)
	do.Provide[authHandler.AuthorizationHandler](Injector, authHandler.NewHandlerInject)
	do.Provide[fileHandler.FileHandler](Injector, fileHandler.NewInject)

}
