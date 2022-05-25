package service

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"mongo-fiber-api/repository"
	"time"
)

type singleBanner struct {
	Title          string `json:"title"`
	Summary        string `json:"summary"`
	MediaReference string `json:"media_reference"`
	ExternalLink   string `json:"external_link"`
	Priority       string `json:"priority"`
}
type bindBanner struct {
	UsaLarge []singleBanner `json:"usaLarge"`
	UkLarge  []singleBanner `json:"ukLarge"`
	UsaSmall []singleBanner `json:"usaSmall"`
	UkSmall  []singleBanner `json:"ukSmall"`
}

func (b *Banner) GetBindBanners(c *fiber.Ctx, email string) (BindBanner bindBanner, err error) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second*2)
	defer cancel()
	bannerRepo := repository.Banner{Ctx: ctx}
	_, err = bannerRepo.GetUserByEmail(email)
	if err != nil {
		return bindBanner{}, errors.New("error in  loading data")
	}

	cursor, err := bannerRepo.GetTaggedStatusBanner("large-us", "active")
	if err = cursor.All(ctx, &BindBanner.UsaLarge); err != nil {
		return
	}
	cursor, err = bannerRepo.GetTaggedStatusBanner("large-uk", "active")
	if err = cursor.All(ctx, &BindBanner.UkLarge); err != nil {
		return
	}
	cursor, err = bannerRepo.GetTaggedStatusBanner("mini-us", "active")
	if err = cursor.All(ctx, &BindBanner.UsaSmall); err != nil {
		return
	}
	cursor, err = bannerRepo.GetTaggedStatusBanner("mini-uk", "active")
	if err = cursor.All(ctx, &BindBanner.UkSmall); err != nil {
		return
	}
	return
}
