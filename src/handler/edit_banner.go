package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo-fiber-api/dto"
	response "mongo-fiber-api/responses"
	"mongo-fiber-api/service"
)

func (h *Handler) EditABanner(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	input := new(dto.EditBannerInput)
	BannerId := c.Params("bannerId")
	objId, _ := primitive.ObjectIDFromHex(BannerId)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: response.BodyParseFailedErrorMsg,
			Errors:  errors.New(response.BodyParseFailedErrorMsg),
		})
	}

	if validationErr := input.Validate(); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: response.ValidationFailedMsg,
			Errors:  validationErr,
		})
	}

	svc := service.Banner{}
	if err := svc.EditBanner(c, objId, *input, email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Payload{
			Message: err.Error(),
			Errors:  err,
		})
	}
	return c.JSON(response.Payload{
		Message: "Banner Edited Successfully",
	})
}
