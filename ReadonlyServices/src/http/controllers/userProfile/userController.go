package controllers

import (
	"context"

	"github.com/TimDebug/FitByte/src/services"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do/v2"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func NewUserControllerInject(i do.Injector) (UserController, error) {
	return NewUserController(
		do.MustInvoke[services.UserService](i),
	), nil
}

func (u *UserController) GetProfile(c *fiber.Ctx) error {
	id, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
	}

	response, err := u.userService.GetProfile(context.Background(), id)
	if err != nil {
		return err
	}

	c.Set("X-Author", "TIM-DEBUG")
	return c.Status(fiber.StatusOK).JSON(response)
}
