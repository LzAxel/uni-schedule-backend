package app

import (
	"fmt"
	"uni-schedule-backend/internal/config"
	"uni-schedule-backend/internal/repository/psql"
	"uni-schedule-backend/internal/service"
)

type App struct {
}

func New() App {
	return App{}
}

func (a *App) Init() {
	cfg := config.GetConfig()

	db, err := psql.NewDBConnection(cfg.Postgres.GetDSN())
	if err != nil {
		panic(err)
	}
	repo := psql.NewPostgresRepository(db)
	serviceInstance := service.NewService(repo)
	fmt.Println(serviceInstance)
}
