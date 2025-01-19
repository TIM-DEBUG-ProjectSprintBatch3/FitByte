package activityController

import (
	"strings"
	"time"

	"github.com/TimDebug/FitByte/src/exceptions"
	"github.com/TimDebug/FitByte/src/model/dtos/request"
	Entity "github.com/TimDebug/FitByte/src/model/entities/activity"
	activityService "github.com/TimDebug/FitByte/src/services/activity"
	"github.com/gofiber/fiber/v2"
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

func (ac *ActivityController) Create(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			exceptions.NewUnauthorizedError(
				fiber.ErrUnauthorized.Error(),
				fiber.StatusUnauthorized,
			),
		)
	}

	req := request.RequestActivity{}
	req.UserId = &userId

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			exceptions.NewBadRequestError(
				fiber.ErrBadRequest.Error(),
				fiber.StatusBadRequest,
			),
		)
	}

	errMsg := validateCreateReq(req)
	if errMsg != "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			exceptions.NewBadRequestError(
				errMsg,
				fiber.StatusBadRequest,
			),
		)
	}

	response, err := ac.activityService.Create(c.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), "23503") { // userId doesnt exist anymore
			return c.Status(fiber.StatusUnauthorized).JSON(
				exceptions.NewUnauthorizedError(
					fiber.ErrUnauthorized.Message,
					fiber.StatusUnauthorized,
				),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			exceptions.NewBadRequestError(
				"Internal server error",
				fiber.StatusInternalServerError,
			),
		)
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(fiber.StatusCreated).JSON(response)
}

func validateCreateReq(req request.RequestActivity) (errMsg string) {
	if req.ActivityType == nil || req.DoneAt == nil || req.DurationInMinutes == nil ||
		*req.ActivityType == "" || *req.DoneAt == "" || *req.DurationInMinutes < 1 {
		return "Require valid values for all properties"
	}

	if *req.DoneAt != "" {
		parsedTime, err := time.Parse(time.RFC3339, *req.DoneAt)
		if err != nil {
			return "Invalid date format, expected RFC3339"
		}
		*req.DoneAt = parsedTime.Format(time.RFC3339)
	}

	if !Entity.IsValidActivityType(*req.ActivityType) {
		return "Invalid activity format"
	}

	return ""
}
