package jwt

import (
	"time"
)

type JWTManager struct {
	issuer               string
	accessTokenLifetime  time.Duration
	refreshTokenLifetime time.Duration
	accessTokenSecret    string
	refreshTokenSecret   string
}

func NewJWTManager(config JWTConfig) *JWTManager {
	return &JWTManager{
		issuer:               config.Issuer,
		accessTokenLifetime:  config.AccessTokenLifetime,
		refreshTokenLifetime: config.RefreshTokenLifetime,
		accessTokenSecret:    config.AccessTokenSecret,
		refreshTokenSecret:   config.RefreshTokenSecret,
	}
}
