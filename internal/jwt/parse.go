package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"strconv"
	"time"
	"uni-schedule-backend/internal/domain"
)

func (j *JWTManager) ParseAccessToken(token string) (domain.ID, error) {
	claims, err := j.parseToken(token, j.accessTokenSecret)
	if err != nil {
		return 0, err
	}

	err = j.validateTokenExpiration(claims)
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(ErrInvalidToken, "failed to parse token claims: %v", err)
	}

	return domain.ID(id), nil
}

func (j *JWTManager) ParseRefreshToken(token string) (domain.ID, error) {
	claims, err := j.parseToken(token, j.refreshTokenSecret)
	if err != nil {
		return 0, err
	}

	err = j.validateTokenExpiration(claims)
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(ErrInvalidToken, "failed to parse token claims: %v", err)
	}

	return domain.ID(id), nil
}

func (j *JWTManager) parseToken(token string, secret string) (*jwt.StandardClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Wrapf(ErrInvalidToken, "unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.Wrapf(ErrInvalidToken, "failed to parse token: %v", err)
	}

	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
	if !ok || !parsedToken.Valid {
		return nil, errors.Wrapf(ErrInvalidToken, "failed to parse token claims: %v", err)
	}

	return claims, nil
}

func (j *JWTManager) validateTokenExpiration(claims *jwt.StandardClaims) error {
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return ErrTokenExpired
	}

	return nil
}
