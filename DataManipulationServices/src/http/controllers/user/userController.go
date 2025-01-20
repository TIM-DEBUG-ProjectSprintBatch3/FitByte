package userController

import (
	"context"
	"net/http"

	"github.com/TimDebug/FitByte/src/exceptions"
	"github.com/TimDebug/FitByte/src/model/dtos/request"
	userService "github.com/TimDebug/FitByte/src/services/user"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type UserController struct {
	userService userService.UserServiceInterface
}

func NewUserController(userService userService.UserServiceInterface) UserControllerInterface {
	return &UserController{userService: userService}
}

func NewUserControllerInject(i do.Injector) (UserControllerInterface, error) {
	_userService := do.MustInvoke[userService.UserServiceInterface](i)
	return NewUserController(_userService), nil
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	bodyParsed := request.UserRegister{}
	if err := c.BodyParser(&bodyParsed); err != nil {
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.Login(context.Background(), bodyParsed)
	if err != nil {
		return c.Status(int(err.(exceptions.ErrorResponse).StatusCode)).
			JSON(err)
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(http.StatusOK).JSON(response)
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	userRequestParse := request.UserRegister{}

	if err := c.BodyParser(&userRequestParse); err != nil {
		return c.Status(http.StatusBadRequest).JSON(exceptions.ErrBadRequest(err.Error()))
	}

	response, err := uc.userService.Register(context.Background(), userRequestParse)
	if err != nil {
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

	updateRequest := request.UpdateProfile{}

	if err := c.BodyParser(&updateRequest); err != nil {
		panic(exceptions.NewBadRequestError(err.Error(), 400))
	}
	response, err := uc.userService.Update(context.Background(), userId, updateRequest)
	if err != nil {
		return err
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(201).JSON(response)
}
