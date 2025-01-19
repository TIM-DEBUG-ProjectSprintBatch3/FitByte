package userService

import (
	"context"

	"github.com/TimDebug/FitByte/src/model/dtos/request"
	"github.com/TimDebug/FitByte/src/model/dtos/response"
)

type UserServiceInterface interface {
	Register(ctx context.Context, input request.UserRegister) (response.UserRegister, error)
}
