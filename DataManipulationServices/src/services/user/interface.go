package userService

import (
	"context"

	"github.com/TimDebug/FitByte/src/model/dtos/request"
	"github.com/TimDebug/FitByte/src/model/dtos/response"
)

type UserServiceInterface interface {
	Login(ctx context.Context, input request.UserRegister) (response.UserRegister, error)
	Register(ctx context.Context, input request.UserRegister) (response.UserRegister, error)
	Update(ctx context.Context, id string, input request.UpdateProfile) (*response.UpdateProfile, error)
}
