package config

type AppConfig struct {
	Port     string `yaml:"port" env:"APP_PORT"`
	LogLevel string `yaml:"logLevel" env:"APP_LOG_LEVEL"`
}
