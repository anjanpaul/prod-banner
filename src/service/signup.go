package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo-fiber-api/dto"
	"mongo-fiber-api/repository"
)

func (b *Banner) Signup(input dto.LoginInput, c *fiber.Ctx) (err error) {
	genericSignupFailureMsg := errors.New("Signup failed for some technical reason.")
	ctx := c.Context()
	bannerRepo := repository.Banner{Ctx: ctx}
	passwordService := PasswordService{}

	// try to get existing user
	existingUser, err := bannerRepo.GetUserByEmail(input.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error(err.Error())
		return genericSignupFailureMsg
	}
	if existingUser != nil {
		return errors.New("An user with this email already exists")
	}
	createVerificationLink := func(sessCtx mongo.SessionContext) (i interface{}, err error) {
		AuthRpo := repository.Banner{Ctx: sessCtx}
		hashedPassword := passwordService.HashPassword(input.Password)
		if _, err = AuthRpo.CreateUser(input.Email, hashedPassword); err != nil {
			return
		}
		return
	}
	var sess mongo.Session
	if sess, err = repository.MongoClient.StartSession(); err != nil {
		log.Error(err.Error())
		return genericSignupFailureMsg
	}
	defer sess.EndSession(ctx)

	if _, err = sess.WithTransaction(ctx, createVerificationLink); err != nil {
		log.Error(err.Error())
		return genericSignupFailureMsg
	}

	return
}
