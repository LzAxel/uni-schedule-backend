package config

import (
	"os"
	"sync"
	"uni-schedule-backend/internal/jwt"
	"uni-schedule-backend/internal/repository/psql"

	"github.com/ilyakaznacheev/cleanenv"
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
		_, err := os.Stat(configPath)
		if err != nil {
			if os.IsNotExist(err) {
				err = cleanenv.ReadEnv(&cfg)
			}
		} else {
			err = cleanenv.ReadConfig(configPath, &cfg)
		}
		if err != nil {
			panic(err)
		}
	})
	return cfg
}
