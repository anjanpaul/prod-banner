package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	response "mongo-fiber-api/responses"
	"mongo-fiber-api/service"
)

func (h *Handler) GetABanner(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	userId := c.Params("bannerId", "")
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: err.Error(),
			Errors:  err,
		})
	}
	svc := service.Banner{}
	banner, err := svc.GetABanner(c, objId, email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: err.Error(),
			Errors:  err,
		})
	}
	return c.JSON(response.Payload{
		Message: "banner found",
		Data:    fiber.Map{"banner": banner},
		Errors:  nil,
	})

}
