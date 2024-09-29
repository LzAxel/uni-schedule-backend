package domain

import "time"

type RefreshToken struct {
	UserID       ID        `db:"user_id"`
	RefreshToken string    `db:"token"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func NewTokenPair(accessToken, refreshToken string) TokenPair {
	return TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
