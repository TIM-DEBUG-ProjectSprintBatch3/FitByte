package activityService

import (
	"context"

	"github.com/TimDebug/FitByte/src/model/dtos/response"
)

type ActivityServiceInterface interface {
	GetAll(ctx context.Context, queryArgs []interface{}) ([]response.ResponseActivity, error)
}
