package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo-fiber-api/dto"
	"mongo-fiber-api/repository"
	"time"
)

func (b *Banner) EditBanner(c *fiber.Ctx, objId primitive.ObjectID, input dto.EditBannerInput, email string) (err error) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*30)
	defer cancel()
	bannerRepo := repository.Banner{Ctx: ctx}
	_, err = bannerRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("error in  editing banner")
	}
	err = bannerRepo.EditBanner(objId, input)
	if err != nil {
		log.Error(err)
		return err
	}
	return

}
