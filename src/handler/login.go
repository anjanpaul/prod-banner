package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"mongo-fiber-api/dto"
	response "mongo-fiber-api/responses"
	"mongo-fiber-api/service"
)

func (h *Handler) Login(c *fiber.Ctx) error {
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
	res := svc.Login(*input)
	if res.Redirect {
		return c.JSON(response.Payload{Message: "DONE", Data: fiber.Map{"token": res.Token}})
	}
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: res.Error.Error(),
			Errors:  err,
		})
	}
	return nil
}
