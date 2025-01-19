package activityService

import (
	"context"

	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/request"
	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/response"
)

type ActivityServiceInterface interface {
	Create(ctx context.Context, req request.RequestActivity) (response.ResponseActivity, error)
	Update(ctx context.Context, req request.RequestActivity, userId, activityId string) (response.ResponseActivity, error)
}
