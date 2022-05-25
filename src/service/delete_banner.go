package service

import (
	"context"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo-fiber-api/repository"
	"time"
)

func (b *Banner) DeleteBanner(c *fiber.Ctx, id primitive.ObjectID, email string) (err error) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*30)
	defer cancel()
	bannerRepo := repository.Banner{Ctx: ctx}

	_, err = bannerRepo.DeleteBanner(id)
	if err != nil {
		log.Error(err)
		return err
	}
	return
}
