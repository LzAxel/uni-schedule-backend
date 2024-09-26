package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"uni-schedule-backend/internal/jwt"
	"uni-schedule-backend/internal/repository/psql"
)

const configPath = "./configs/dev.yml"

var (
	cfg  = Config{}
	once = new(sync.Once)
)

type Config struct {
	AppConfig AppConfig     `yaml:"app"`
	Postgres  psql.Config   `yaml:"postgres"`
	JWT       jwt.JWTConfig `yaml:"jwt"`
}

func GetConfig() Config {
	once.Do(func() {
		err := cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			panic(err)
		}
	})
	return cfg
}
