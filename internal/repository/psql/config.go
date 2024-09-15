package psql

import "fmt"

type Config struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASS"`
	DBName   string `yaml:"dbName" env:"DB_NAME"`
	SSLMode  string `yaml:"sslMode" env:"DB_SSL_MODE"`
}

func (c Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.Host, c.Port, c.User, c.DBName, c.Password, c.SSLMode)
}
