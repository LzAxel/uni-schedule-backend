package auth

import (
	"fmt"
	"time"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
)

func (s *AuthService) generateTokenPair(userID domain.ID) (domain.TokenPair, error) {
	accessToken, err := s.jwtManager.GenerateAccessToken(userID)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("AuthService.generateTokenPair.GenerateAccessToken: %w", err)
	}
	refreshToken, err := s.jwtManager.GenerateRefreshToken(userID)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("AuthService.generateTokenPair.GenerateRefreshToken: %w", err)
	}

	return domain.NewTokenPair(accessToken, refreshToken), nil
}

func (s *AuthService) generateAndStoreTokenPair(userID domain.ID) (domain.TokenPair, error) {
	tokenPair, err := s.generateTokenPair(userID)
	if err != nil {
		return domain.TokenPair{}, err
	}

	err = s.tokenRepo.CreateOrUpdate(domain.RefreshToken{
		UserID:       userID,
		RefreshToken: tokenPair.RefreshToken,
		UpdatedAt:    time.Now().UTC(),
	})
	if err != nil {
		return domain.TokenPair{}, apperror.NewServiceError("AuthService.generateAndStoreTokenPair: create or update token", err)
	}

	return tokenPair, nil
}
