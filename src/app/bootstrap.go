package app

import (
	"context"
	log "github.com/sirupsen/logrus"
	"mongo-fiber-api/config"
	"mongo-fiber-api/repository"
)

func (app *App) Bootstrap() {
	mongo := repository.Mongo{
		ConnURI: config.Params.MongoURI,
		DBName:  config.Params.MongoDbName,
	}
	app.Mongo = &mongo

	if err := mongo.Connect(); err != nil {
		log.Fatal(err.Error())
		return
	}

	ctx := context.Background()
	bannerRepo := repository.Banner{Ctx: ctx}
	if err := bannerRepo.CreateUserIndex(); err != nil {
		log.Error(err.Error())
	}

	//app.RedisStorage()
}
