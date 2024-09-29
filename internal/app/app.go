package app

import (
	"uni-schedule-backend/internal/config"
	"uni-schedule-backend/internal/handler"
	"uni-schedule-backend/internal/repository/psql"
	"uni-schedule-backend/internal/service"
)

type App struct {
	controller *handler.Controller
}

func New() App {
	return App{}
}

func (a *App) init() {
	cfg := config.GetConfig()

	db, err := psql.NewDBConnection(cfg.Postgres.GetDSN())
	if err != nil {
		panic(err)
	}
	err = psql.UpMigrations(db)
	if err != nil {
		panic(err)
	}
	repo := psql.NewPostgresRepository(db)
	serviceInstance := service.NewService(repo)
	a.controller = handler.NewController(serviceInstance)
}

func (a *App) Run() {
	a.init()
	a.controller.Start()
}

func (a *App) Stop() {

}
