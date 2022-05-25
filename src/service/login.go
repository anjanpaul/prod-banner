package service

import (
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"mongo-fiber-api/config"
	"mongo-fiber-api/dto"
	"mongo-fiber-api/repository"
	"time"
)

type LoginResponse struct {
	Redirect bool
	Token    string
	Error    error
}

func (b *Banner) Login(input dto.LoginInput) (res LoginResponse) {
	//genericLoginFailureMsg := errors.New("Login failed for some technical reason.")
	var ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	bannerRepo := repository.Banner{Ctx: ctx}
	passwordService := PasswordService{}
	_, inputMarshalError := json.Marshal(input)

	// try to get existing user in db
	existingUser, err := bannerRepo.GetUserByEmail(input.Email)
	if err == nil {
		passwordMatched := passwordService.ComparePasswords(existingUser.Password, []byte(input.Password))
		//JWT TOKEN  GENERATE
		claims := jwt.MapClaims{
			"email": existingUser.Email,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, _ := token.SignedString([]byte(config.Params.JWTSignatureSecret))
		if passwordMatched {
			return LoginResponse{
				Redirect: true,
				Token:    signedToken,
				Error:    inputMarshalError,
			}
		}
		return LoginResponse{Error: errors.New("Wrong password")}
	}
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Error(errors.New("User not found"))
		}
		log.Error(err.Error())
		return
	}
	return
}
