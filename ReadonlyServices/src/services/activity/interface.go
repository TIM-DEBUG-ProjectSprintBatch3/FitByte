package activityService

import (
	"context"

	"github.com/rafitanujaya/go-fiber-template/src/model/dtos/response"
)

type ActivityServiceInterface interface {
	GetAll(ctx context.Context, queryArgs []interface{}) ([]response.ResponseActivity, error)
}
