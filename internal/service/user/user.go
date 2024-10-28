package user

import (
	"time"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
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

func (s *UserService) Create(user domain.User) (uint64, error) {
	return s.userRepo.Create(domain.UserCreate{
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Role:         user.Role,
		CreatedAt:    time.Now(),
	})
}
func (s *UserService) GetByID(id uint64) (domain.User, error) {
	return s.userRepo.GetByID(id)
}
func (s *UserService) GetByUsername(username string) (domain.User, error) {
	return s.userRepo.GetByUsername(username)
}
func (s *UserService) Update(userID uint64, id uint64, update domain.UserUpdateDTO) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	if user.ID != userID {
		return apperror.ErrDontHavePermission
	}

	return s.userRepo.Update(id, update)
}
func (s *UserService) Delete(id uint64) error {
	return s.userRepo.Delete(id)
}
