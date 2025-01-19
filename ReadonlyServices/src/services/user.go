package services

import (
	"context"

	functionCallerInfo "github.com/TimDebug/FitByte/src/logger/helper"
	loggerZap "github.com/TimDebug/FitByte/src/logger/zap"
	"github.com/TimDebug/FitByte/src/model/dtos/response"
	"github.com/TimDebug/FitByte/src/repositories"
	"github.com/samber/do/v2"
)

type UserService interface {
	GetProfile(ctx context.Context, id string) (*response.ProfileResponse, error)
}

type UserServiceImpl struct {
	userRepo repositories.UserRepository
	logger   loggerZap.LoggerInterface
}

func (u UserServiceImpl) GetProfile(
	ctx context.Context,
	id string,
) (*response.ProfileResponse, error) {
	profile, err := u.userRepo.FindById(ctx, id)
	if err != nil {
		u.logger.Error(err.Error(), functionCallerInfo.UserServiceGetProfile, err)
		return nil, err
	}

	return &response.ProfileResponse{
		Preference: profile.Preference,
		WeightUnit: profile.WeightUnit,
		HeightUnit: profile.HeightUnit,
		Weight:     profile.Weight,
		Height:     profile.Height,
		Email:      profile.Email,
		Name:       profile.Name,
		ImageUri:   profile.ImageUri,
	}, nil
}

func NewUserServiceImpl(
	userRepository repositories.UserRepository,
	logger loggerZap.LoggerInterface,
) UserService {
	return &UserServiceImpl{
		userRepo: userRepository,
		logger:   logger,
	}
}

func NewUserServiceImplInject(i do.Injector) (UserService, error) {
	return NewUserServiceImpl(
		do.MustInvoke[repositories.UserRepository](i),
		do.MustInvoke[loggerZap.LoggerInterface](i),
	), nil
}
