package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"mongo-fiber-api/dto"
	response "mongo-fiber-api/responses"
	"mongo-fiber-api/service"
)

func (h *Handler) Signup(c *fiber.Ctx) error {
	input := new(dto.LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: response.BodyParseFailedErrorMsg,
			Errors:  errors.New(response.BodyParseFailedErrorMsg),
		})
	}

	err := input.Validate()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: response.ValidationFailedMsg,
			Errors:  err,
		})
	}

	svc := service.Banner{}
	err = svc.Signup(*input, c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: err.Error(),
			Errors:  err,
		})
	}

	return c.JSON(response.Payload{
		Message: "user created successfully",
	})
}
