package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"mongo-fiber-api/dto"
	"mongo-fiber-api/repository"
	"time"
)

func (b *Banner) CreateBanner(input *dto.CreateBannerInput, c *fiber.Ctx, email string) (err error) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*30)
	defer cancel()
	bannerRepo := repository.Banner{Ctx: ctx}
	_, err = bannerRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("error in creating banner")
	}
	_, err = bannerRepo.CreateBanner(input)
	if err != nil {
		log.Error(err)
		return err
	}
	return

}
