package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	response "mongo-fiber-api/responses"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Home(c *fiber.Ctx) error {
	return c.JSON(response.Payload{
		Message: "BANNERS!",
	})
}

func (h *Handler) NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(response.Payload{
		Message: "404!",
		Errors:  errors.New("path not found"),
	})
}
