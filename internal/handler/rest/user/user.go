package user

import (
	"main/internal/entity"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) DisplayAllUsers(c *fiber.Ctx) error {
	users, err := h.uc.DisplayAllUsers(c.Context())
	if err != nil {
		c.JSON(err.Error())
		return err
	}

	c.JSON(users)
	return nil
}

func (h Handler) DisplayUser(c *fiber.Ctx) error {
	input := new(entity.DisplayUser)

	err := c.BodyParser(input)
	if err != nil {
		c.JSON(err.Error())
		return err
	}

	user, err := h.uc.DisplayUser(c.Context(), input.Userid)
	if err != nil {
		c.JSON(err.Error())
		return err
	}

	c.JSON(user)
	return nil
}
