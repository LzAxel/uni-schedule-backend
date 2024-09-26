package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
	"uni-schedule-backend/internal/domain"
)

func (j *JWTManager) GenerateAccessToken(userID domain.ID) (string, error) {
	return j.generateToken(j.accessTokenLifetime, j.accessTokenSecret, userID)
}

func (j *JWTManager) GenerateRefreshToken(userID domain.ID) (string, error) {
	return j.generateToken(j.refreshTokenLifetime, j.refreshTokenSecret, userID)
}

func (j *JWTManager) generateToken(tokenLifetime time.Duration, tokenSecret string, userID domain.ID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Add(tokenLifetime).Unix(),
		IssuedAt:  time.Now().UTC().Unix(),
		Issuer:    j.issuer,
		NotBefore: time.Now().UTC().Unix(),
		Subject:   strconv.FormatUint(uint64(userID), 10),
	})

	signedToken, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}
