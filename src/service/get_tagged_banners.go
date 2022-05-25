package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"mongo-fiber-api/repository"
	"time"
)

func (b *Banner) GetTaggedBanners(c *fiber.Ctx, tag string, email string) (taggedBanners []bson.M, err error) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*2)
	defer cancel()
	bannerRepo := repository.Banner{Ctx: ctx}
	_, err = bannerRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("error in  loading data")
	}

	cursor, err := bannerRepo.GetTaggedBanner(tag)

	if err = cursor.All(ctx, &taggedBanners); err != nil {
		return
	}
	return
}
