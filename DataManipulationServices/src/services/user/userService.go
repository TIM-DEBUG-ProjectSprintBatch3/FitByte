package userService

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	authJwt "github.com/rafitanujaya/go-fiber-template/src/auth/jwt"
	"github.com/rafitanujaya/go-fiber-template/src/exceptions"
	functionCallerInfo "github.com/rafitanujaya/go-fiber-template/src/logger/helper"
	loggerZap "github.com/rafitanujaya/go-fiber-template/src/logger/zap"
	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/request"
	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/response"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities/user"
	userRepository "github.com/rafitanujaya/go-fiber-template/src/repositories/user"
	"github.com/samber/do/v2"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	UserRepository userRepository.UserRepositoryInterface
	Db             *pgxpool.Pool
	jwtService     authJwt.JwtServiceInterface
	Logger         loggerZap.LoggerInterface
}

func NewUserService(userRepo userRepository.UserRepositoryInterface, db *pgxpool.Pool, jwtService authJwt.JwtServiceInterface, logger loggerZap.LoggerInterface) UserServiceInterface {
	return &userService{UserRepository: userRepo, Db: db, jwtService: jwtService, Logger: logger}
}

func NewUserServiceInject(i do.Injector) (UserServiceInterface, error) {
	_db := do.MustInvoke[*pgxpool.Pool](i)
	_userRepo := do.MustInvoke[userRepository.UserRepositoryInterface](i)
	_jwtService := do.MustInvoke[authJwt.JwtServiceInterface](i)
	_logger := do.MustInvoke[loggerZap.LoggerInterface](i)

	return NewUserService(_userRepo, _db, _jwtService, _logger), nil
}

func (us *userService) Register(ctx context.Context, input request.UserRegister) (response.UserRegister, error) {
	// TODO Validate Input

	user := Entity.User{}

	timeNow := time.Now()
	user.CreatedAt = timeNow
	user.UpdatedAt = timeNow
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return response.UserRegister{}, err
	}

	user.PasswordHash = string(passwordHash)

	user.Id, err = us.UserRepository.CreateUser(ctx, us.Db, user)

	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return response.UserRegister{}, exceptions.NewConflictError("Data Conflict", 409)
		} else {
			us.Logger.Error(err.Error(), functionCallerInfo.UserServiceRegister)
			return response.UserRegister{}, err
		}
	}

	token, err := us.jwtService.GenerateToken(user.Id)

	if err != nil {
		return response.UserRegister{}, err
	}

	return response.UserRegister{
		Email: user.Email,
		Token: token,
	}, nil

}

func (us *userService) Update(ctx context.Context, id string, req request.UpdateProfile) (*response.UpdateProfile, error) {
	profile, err := us.UserRepository.FindById(context.Background(), id)
	if err != nil {
		us.Logger.Error(err.Error(), functionCallerInfo.UserServiceUpdate, err)
		return nil, err
	}

	profile.Preference = &req.Preference
	profile.WeightUnit = &req.WeightUnit
	profile.HeightUnit = &req.HeightUnit
	profile.Weight = &req.Weight
	profile.Height = &req.Height
	if req.Name != nil {
		profile.Name = req.Name
	}
	if req.ImageUri != nil {
		profile.ImageUri = req.ImageUri
	}

	_, err = us.UserRepository.Update(context.Background(), id, *profile)
	if err != nil {
		us.Logger.Error(err.Error(), functionCallerInfo.UserServiceUpdate, err)
		return nil, err
	}

	return &response.UpdateProfile{
		Preference: *profile.Preference,
		WeightUnit: *profile.WeightUnit,
		HeightUnit: *profile.HeightUnit,
		Weight:     *profile.Weight,
		Height:     *profile.Height,
		Name:       profile.Name,
		ImageUri:   profile.ImageUri,
	}, nil
}
