package service

import (
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
	"uni-schedule-backend/internal/user/model"
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
	return s.userRepo.Create(user)
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
