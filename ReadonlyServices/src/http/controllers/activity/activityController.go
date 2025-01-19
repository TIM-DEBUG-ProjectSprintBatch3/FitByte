package activityController

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rafitanujaya/go-fiber-template/src/exceptions"
	Entity "github.com/rafitanujaya/go-fiber-template/src/model/entities/activity"
	activityService "github.com/rafitanujaya/go-fiber-template/src/services/activity"
	"github.com/samber/do/v2"
)

type ActivityController struct {
	activityService activityService.ActivityServiceInterface
}

func NewActivityController(activityService activityService.ActivityServiceInterface) ActivityControllerInterface {
	return &ActivityController{activityService: activityService}
}

func NewActivityControllerInject(i do.Injector) (ActivityControllerInterface, error) {
	_activityService := do.MustInvoke[activityService.ActivityServiceInterface](i)
	return NewActivityController(_activityService), nil
}

func (ac *ActivityController) GetAll(c *fiber.Ctx) error {
	id, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(exceptions.NewUnauthorizedError(fiber.ErrUnauthorized.Error(), fiber.StatusUnauthorized))
	}

	params := make(map[string]string)
	params["id"] = id
	params["activityType"] = c.Query("activityType", "")
	params["doneAtFrom"] = c.Query("doneAtFrom", "")
	params["doneAtTo"] = c.Query("doneAtTo", "")
	params["caloriesBurnedMin"] = c.Query("caloriesBurnedMin", "")
	params["caloriesBurnedMax"] = c.Query("caloriesBurnedMax", "")

	response, err := ac.activityService.GetAll(context.Background(), buildQueryParams(c, params))

	if err != nil {
		return err
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(fiber.StatusOK).JSON(response)
}

func GetCaloriesPerMinute(activityType Entity.ActivityType) (float64, error) {
	switch activityType {
	case Entity.Walking, Entity.Yoga, Entity.Stretching:
		return 4.0, nil
	case Entity.Cycling, Entity.Swimming, Entity.Dancing:
		return 8.0, nil
	case Entity.Hiking, Entity.Running, Entity.HIIT, Entity.JumpRope:
		return 10.0, nil
	default:
		return 0, fiber.ErrBadRequest
	}
}

func IsValidActivityType(activityType Entity.ActivityType) bool {
	_, err := GetCaloriesPerMinute(activityType)
	return err == nil
}

func getQueryInt(ctx *fiber.Ctx, key string, defaultValue int) int {
	queries := ctx.Queries()
	value, exists := queries[key]
	if !exists {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func buildQueryParams(ctx *fiber.Ctx, params map[string]string) []interface{} {
	limit := getQueryInt(ctx, "limit", 5)
	offset := getQueryInt(ctx, "offset", 0)
	activityType := params["activityType"]
	doneAtFrom := params["doneAtFrom"]
	doneAtTo := params["doneAtTo"]
	caloriesBurnedMin := getQueryInt(ctx, "caloriesBurnedMin", 0)
	caloriesBurnedMax := getQueryInt(ctx, "caloriesBurnedMax", 0)
	id := params["id"]

	queryArgs := []interface{}{id, nil, nil, nil, nil, nil, limit, offset}

	// validate activityType
	if activityType != "" && IsValidActivityType(Entity.ActivityType(activityType)) {
		queryArgs[1] = activityType
	}

	// validate doneAtFrom
	if doneAtFrom != "" {
		if parsedDate, err := time.Parse(time.RFC3339, doneAtFrom); err == nil {
			queryArgs[2] = parsedDate
		}
	}

	// validate doneAtTo
	if doneAtTo != "" {
		if parsedDate, err := time.Parse(time.RFC3339, doneAtTo); err == nil {
			queryArgs[3] = parsedDate
		}
	}

	// validate caloriesBurnedMin
	if caloriesBurnedMin > 0 {
		queryArgs[4] = caloriesBurnedMin
	}

	// validate caloriesBurnedMax
	if caloriesBurnedMax > 0 {
		queryArgs[5] = caloriesBurnedMax
	}

	return queryArgs
}
