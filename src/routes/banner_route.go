package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"mongo-fiber-api/config"
	"mongo-fiber-api/handler"
)

func Banner(server *fiber.App, handler *handler.Handler) {
	banners := server.Group("/banners")
	banners.Get("/", handler.Home)
	banners.Post("/signup", handler.Signup)
	banners.Post("/login", handler.Login)

	// JWT Middleware
	banners.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Params.JWTSignatureSecret),
	}))
	//All routes related to users comes here
	banners.Get("/show-banners", handler.GetAllBanners)
	banners.Post("/add-banner", handler.CreateBanner)
	banners.Delete("/delete/:bannerId", handler.DeleteABanner)
	banners.Get("/show-banner/:bannerId", handler.GetABanner)
	banners.Put("/edit/:bannerId", handler.EditABanner)
	banners.Get("/tagged-banners", handler.GetTaggedBanners)
	banners.Get("/status-banners", handler.GetStatusBanners)
	banners.Get("/bind-banners", handler.BindBanners)
	banners.Get("/tagged-status-banners", handler.GetTaggedStatusBanners)

}
