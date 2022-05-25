package app

import "mongo-fiber-api/repository"

type App struct {
	Mongo *repository.Mongo
}

func New() *App {
	return &App{}
}
