package activityService

import (
	"context"
	"fmt"

	authJwt "github.com/TimDebug/FitByte/src/auth/jwt"
	functionCallerInfo "github.com/TimDebug/FitByte/src/logger/helper"
	loggerZap "github.com/TimDebug/FitByte/src/logger/zap"
	"github.com/TimDebug/FitByte/src/model/dtos/response"
	activityRepository "github.com/TimDebug/FitByte/src/repositories/activity"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

type activityService struct {
	ActivityRepository activityRepository.ActivityRepositoryInterface
	Db                 *pgxpool.Pool
	jwtService         authJwt.JwtServiceInterface
	Logger             loggerZap.LoggerInterface
}

func NewActivityService(activityRepo activityRepository.ActivityRepositoryInterface, db *pgxpool.Pool, jwtService authJwt.JwtServiceInterface, logger loggerZap.LoggerInterface) ActivityServiceInterface {
	return &activityService{ActivityRepository: activityRepo, Db: db, jwtService: jwtService, Logger: logger}
}

func NewActivityServiceInject(i do.Injector) (ActivityServiceInterface, error) {
	_db := do.MustInvoke[*pgxpool.Pool](i)
	_activityRepo := do.MustInvoke[activityRepository.ActivityRepositoryInterface](i)
	_jwtService := do.MustInvoke[authJwt.JwtServiceInterface](i)
	_logger := do.MustInvoke[loggerZap.LoggerInterface](i)

	return NewActivityService(_activityRepo, _db, _jwtService, _logger), nil
}

func (us *activityService) GetAll(ctx context.Context, queryArgs []interface{}) ([]response.ResponseActivity, error) {
	fmt.Printf("queryArgs: %v", queryArgs...)
	rawActivities, err := us.ActivityRepository.GetAll(ctx, us.Db, queryArgs)

	if err != nil {
		us.Logger.Error(err.Error(), functionCallerInfo.ActivityServiceGetAll, rawActivities)

		return []response.ResponseActivity{}, err
	}

	returnedActivities := make([]response.ResponseActivity, 0)
	for _, elem := range rawActivities {
		var activity response.ResponseActivity
		activity.Id = *elem.ActivityId
		activity.ActivityType = *elem.ActivityType
		activity.DoneAt = fmt.Sprintf("%s", *elem.DoneAt)
		activity.DurationInMinutes = int(*elem.DurationInMinutes)
		activity.CaloriesBurned = float64(*elem.CaloriesBurned)
		activity.CreatedAt = fmt.Sprintf("%s", *elem.CreatedAt)
		returnedActivities = append(returnedActivities, activity)
	}
	return returnedActivities, nil

}
