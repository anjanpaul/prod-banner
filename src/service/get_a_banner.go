package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo-fiber-api/repository"
	"time"
)

func (b *Banner) GetABanner(c *fiber.Ctx, id primitive.ObjectID, email string) (banner interface{}, err error) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*10)
	defer cancel()
	bannerRepo := repository.Banner{Ctx: ctx}
	_, err = bannerRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("error in  creating banner")
	}
	banner, err = bannerRepo.GetABanner(id)
	if err != nil {
		log.Error(err)
		return
	}
	return
}
