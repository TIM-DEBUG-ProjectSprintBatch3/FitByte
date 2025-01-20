package userController

import (
	"net/http"

	"github.com/TimDebug/FitByte/src/exceptions"
	functionCallerInfo "github.com/TimDebug/FitByte/src/logger/helper"
	loggerZap "github.com/TimDebug/FitByte/src/logger/zap"
	"github.com/TimDebug/FitByte/src/model/dtos/request"
	userService "github.com/TimDebug/FitByte/src/services/user"
	"github.com/TimDebug/FitByte/src/services/user/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type UserController struct {
	userService userService.UserServiceInterface
	logger      loggerZap.LoggerInterface
}

func NewUserController(userService userService.UserServiceInterface, logger loggerZap.LoggerInterface) UserControllerInterface {
	return &UserController{userService: userService, logger: logger}
}

func NewUserControllerInject(i do.Injector) (UserControllerInterface, error) {
	_userService := do.MustInvoke[userService.UserServiceInterface](i)
	_logger := do.MustInvoke[loggerZap.LoggerInterface](i)
	return NewUserController(_userService, _logger), nil
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	bodyParsed := request.UserRegister{}
	if err := c.BodyParser(&bodyParsed); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLogin)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.Login(c, bodyParsed)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerLogin, bodyParsed)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).
			JSON(err)
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(http.StatusOK).JSON(response)
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	userRequestParse := request.UserRegister{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerRegister)
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.Register(c, userRequestParse)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.UserControllerRegister, userRequestParse)
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).
			JSON(err)
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(201).JSON(response)
}

func (uc *UserController) Update(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(
			exceptions.NewUnauthorizedError(
				fiber.ErrUnauthorized.Error(),
				fiber.StatusUnauthorized,
			),
		)
	}

	updateRequestCustom := request.UpdateProfileCustom{}

	if err := c.BodyParser(&updateRequestCustom); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		// panic(exceptions.NewBadRequestError(err.Error(), 400))
	}

	if updateRequestCustom.Name.IsPresent {
		if updateRequestCustom.Name.IsNull || len(updateRequestCustom.Name.Value) < 2 || len(updateRequestCustom.Name.Value) > 60 {
			return c.Status(fiber.StatusBadRequest).JSON(
				exceptions.NewUnauthorizedError(
					"Invalid Name",
					fiber.StatusBadRequest,
				),
			)
		}
	}

	if updateRequestCustom.ImageUri.IsPresent {
		if updateRequestCustom.ImageUri.IsNull || updateRequestCustom.ImageUri.Value == "" {
			return c.Status(fiber.StatusBadRequest).JSON(
				exceptions.NewUnauthorizedError(
					"Invalid Name",
					fiber.StatusBadRequest,
				),
			)
		}
	}

	updateRequest := request.UpdateProfile{
		Preference: updateRequestCustom.Preference,
		WeightUnit: updateRequestCustom.WeightUnit,
		HeightUnit: updateRequestCustom.HeightUnit,
		Weight:     updateRequestCustom.Weight,
		Height:     updateRequestCustom.Height,
		Name:       &updateRequestCustom.Name.Value,
		ImageUri:   &updateRequestCustom.ImageUri.Value,
	}

	err := validator.ValidateUpdateProfile(updateRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			exceptions.NewUnauthorizedError(
				"Invalid Name",
				fiber.StatusBadRequest,
			),
		)
	}

	response, err := uc.userService.Update(c, userId, updateRequest)
	if err != nil {
		uc.logger.Error(err.Error(), functionCallerInfo.ProfileControllerUpdate, userId, updateRequest)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(200).JSON(response)
}
