package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	response "mongo-fiber-api/responses"
	"mongo-fiber-api/service"
)

func (h *Handler) GetTaggedStatusBanners(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	tag := c.Query("tag")
	status := "active"
	svc := service.Banner{}
	banners, err := svc.GetTaggedStatusBanners(c, tag, status, email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: err.Error(),
			Errors:  err,
		})
	}
	return c.JSON(response.Payload{
		Message: "Banners Found",
		Data:    fiber.Map{"banners": banners},
	})
}
