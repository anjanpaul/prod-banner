package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	response "mongo-fiber-api/responses"
	"mongo-fiber-api/service"
)

//import (
//	"context"
//	"fmt"
//	"github.com/gofiber/fiber/v2"
//	"go.mongodb.org/mongo-driver/bson"
//	response "mongo-fiber-api/responses"
//	"mongo-fiber-api/service"
//	"net/http"
//	"time"
//)
//
//func (h *Handler) GetTaggedBanners(c *fiber.Ctx) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	var banners []models.Banner
//	defer cancel()
//	tags := c.Query("tags")
//	status := c.Query("status")
//	results, err := bannerCollection.Find(ctx, bson.M{"tags": tags, "status": status})
//	fmt.Print(results)
//	if err != nil {
//		return c.Status(http.StatusInternalServerError).JSON(responses.BannerResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
//	}
//	//reading from the db in an optimal way
//	defer results.Close(ctx)
//	for results.Next(ctx) {
//		var singleBanner models.Banner
//		if err = results.Decode(&singleBanner); err != nil {
//			return c.Status(http.StatusInternalServerError).JSON(responses.BannerResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
//		}
//
//		banners = append(banners, singleBanner)
//	}
//
//	return c.Status(http.StatusOK).JSON(
//		responses.BannerResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": banners}},
//	)
//
//}

//type singleBanner struct {
//	ID             string `json:"id"`
//	Title          string `json:"title"`
//	Summary        string `json:"summary"`
//	Status         string `json:"status"`
//	Tags           string `json:"tags"`
//	MediaType      string `json:"media_Type"`
//	MediaReference string `json:"media_reference"`
//	ExternalLink   string `json:"external_link"`
//}

func (h *Handler) GetTaggedBanners(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	tag := c.Query("tag")
	svc := service.Banner{}
	banners, err := svc.GetTaggedBanners(c, tag, email)
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
