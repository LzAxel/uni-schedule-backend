package service

import (
	"errors"
	"fmt"
	"time"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	tokenModel "uni-schedule-backend/internal/domain/auth/model"
	"uni-schedule-backend/internal/domain/user/model"
	"uni-schedule-backend/internal/repository"
	"uni-schedule-backend/pkg/hash"
)

type JWTManager interface {
	ParseAccessToken(token string) (domain.ID, error)
	ParseRefreshToken(token string) (domain.ID, error)
	GenerateAccessToken(userID domain.ID) (string, error)
	GenerateRefreshToken(userID domain.ID) (string, error)
}

type AuthService struct {
	passwordSalt string
	userRepo     repository.UserRepository
	tokenRepo    repository.TokenRepository
	jwtManager   JWTManager
}

func NewAuthService(userRepo repository.UserRepository, tokenRepo repository.TokenRepository, jwtManager JWTManager, passwordSalt string) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		tokenRepo:    tokenRepo,
		jwtManager:   jwtManager,
		passwordSalt: passwordSalt,
	}
}

func (s *AuthService) Login(username, password string) (tokenModel.TokenPair, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return tokenModel.TokenPair{}, apperror.ErrInvalidLoginOrPassword
		}
		return tokenModel.TokenPair{}, err
	}

	if !hash.VerifyPassword(password, user.PasswordHash, s.passwordSalt) {
		return tokenModel.TokenPair{}, apperror.ErrInvalidLoginOrPassword
	}

	tokenPair, err := s.generateTokenPair(user.ID)
	if err != nil {
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.Login:", err)
	}

	err = s.tokenRepo.CreateOrUpdate(tokenModel.RefreshToken{
		UserID:       user.ID,
		RefreshToken: tokenPair.RefreshToken,
		UpdatedAt:    time.Now().UTC(),
	})
	if err != nil {
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.Login: create or update token", err)
	}

	return tokenPair, nil
}

func (s *AuthService) Register(username, password string) (tokenModel.TokenPair, error) {
	passwordHash, err := hash.HashPassword(password, s.passwordSalt)
	if err != nil {
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.Register: hashing password", err)
	}

	createdID, err := s.userRepo.Create(model.UserCreate{
		Username:     username,
		PasswordHash: passwordHash,
		Role:         model.RoleStudent,
		CreatedAt:    time.Now(),
	})
	if err != nil {
		if errors.Is(err, apperror.ErrAlreadyExists) {
			return tokenModel.TokenPair{}, apperror.ErrUsernameAlreadyTaken
		}
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.Register: create user", err)
	}

	tokenPair, err := s.generateTokenPair(createdID)
	if err != nil {
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.Register:", err)
	}

	err = s.tokenRepo.CreateOrUpdate(tokenModel.RefreshToken{
		UserID:       createdID,
		RefreshToken: tokenPair.RefreshToken,
		UpdatedAt:    time.Now().UTC(),
	})
	if err != nil {
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.Register: create or update token", err)
	}

	return tokenPair, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (tokenModel.TokenPair, error) {
	userID, err := s.jwtManager.ParseRefreshToken(refreshToken)
	if err != nil {
		return tokenModel.TokenPair{}, apperror.ErrInvalidRefreshToken
	}

	storedToken, err := s.tokenRepo.GetByUserID(userID)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return tokenModel.TokenPair{}, apperror.ErrUserNotFound
		}
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.RefreshToken: getting user by id", err)
	}

	if storedToken.RefreshToken != refreshToken {
		return tokenModel.TokenPair{}, apperror.ErrInvalidRefreshToken
	}

	tokenPair, err := s.generateTokenPair(userID)
	if err != nil {
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.RefreshToken:", err)
	}

	err = s.tokenRepo.CreateOrUpdate(tokenModel.RefreshToken{
		UserID:       userID,
		RefreshToken: tokenPair.RefreshToken,
		UpdatedAt:    time.Now().UTC(),
	})
	if err != nil {
		return tokenModel.TokenPair{}, apperror.NewServiceError("AuthService.RefreshToken: create or update token", err)
	}

	return tokenPair, nil
}

func (s *AuthService) GetUserFromAccessToken(accessToken string) (model.User, error) {
	userID, err := s.jwtManager.ParseAccessToken(accessToken)
	if err != nil {
		return model.User{}, apperror.ErrInvalidAccessToken
	}
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return model.User{}, apperror.ErrUserNotFound
	}

	return user, nil
}

func (s *AuthService) generateTokenPair(userID domain.ID) (tokenModel.TokenPair, error) {
	accessToken, err := s.jwtManager.GenerateAccessToken(userID)
	if err != nil {
		return tokenModel.TokenPair{}, fmt.Errorf("generateTokenPair.GenerateAccessToken: %w", err)
	}
	refreshToken, err := s.jwtManager.GenerateRefreshToken(userID)
	if err != nil {
		return tokenModel.TokenPair{}, fmt.Errorf("generateTokenPair.GenerateRefreshToken: %w", err)
	}

	return tokenModel.NewTokenPair(accessToken, refreshToken), nil
}
