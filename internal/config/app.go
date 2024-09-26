package config

type AppConfig struct {
	Port         string `yaml:"port" env:"APP_PORT"`
	LogLevel     string `yaml:"logLevel" env:"APP_LOG_LEVEL"`
	IsDebug      bool   `yaml:"isDebug" env:"APP_IS_DEBUG"`
	PasswordSalt string `yaml:"passwordSalt" env:"APP_PASSWORD_SALT"`
}
