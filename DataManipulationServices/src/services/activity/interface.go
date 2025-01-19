package activityService

import (
	"context"

	"github.com/TimDebug/FitByte/src/model/dtos/request"
	"github.com/TimDebug/FitByte/src/model/dtos/response"
)

type ActivityServiceInterface interface {
	Create(ctx context.Context, req request.RequestActivity) (response.ResponseActivity, error)
}
