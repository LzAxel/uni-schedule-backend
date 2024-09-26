package jwt

import "time"

type JWTConfig struct {
	Issuer               string        `yaml:"issuer" env:"JWT_ISSUER"`
	AccessTokenLifetime  time.Duration `yaml:"accessTokenLifetime" env:"JWT_ACCESS_TOKEN_LIFETIME"`
	RefreshTokenLifetime time.Duration `yaml:"refreshTokenLifetime" env:"JWT_REFRESH_TOKEN_LIFETIME"`
	AccessTokenSecret    string        `yaml:"accessTokenSecret" env:"JWT_ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret   string        `yaml:"refreshTokenSecret" env:"JWT_REFRESH_TOKEN_SECRET"`
}
