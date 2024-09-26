package service

import (
	"time"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/user/model"
	"uni-schedule-backend/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(user model.User) (domain.ID, error) {
	return s.userRepo.Create(model.UserCreate{
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Role:         user.Role,
		CreatedAt:    time.Now(),
	})
}
func (s *UserService) GetByID(id domain.ID) (model.User, error) {
	return s.userRepo.GetByID(id)
}
func (s *UserService) GetByUsername(username string) (model.User, error) {
	return s.userRepo.GetByUsername(username)
}
func (s *UserService) Update(id domain.ID, update model.UserUpdateDTO) error {
	return s.userRepo.Update(id, update)
}
func (s *UserService) Delete(id domain.ID) error {
	return s.userRepo.Delete(id)
}
