package teacher

import (
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
)

type TeacherService struct {
	teacherRepo repository.TeacherRepository
}

func NewTeacherService(teacherRepo repository.TeacherRepository) *TeacherService {
	return &TeacherService{
		teacherRepo: teacherRepo,
	}
}

func (s *TeacherService) Create(teacher domain.TeacherCreate) (uint64, error) {
	if teacher.ShortName == "" {
		return 0, apperror.ErrEmptyShortName
	}
	return s.teacherRepo.Create(teacher)
}
func (s *TeacherService) GetByID(id uint64) (domain.Teacher, error) {
	return s.teacherRepo.GetByID(id)
}
func (s *TeacherService) GetAll() ([]domain.Teacher, error) {
	return s.teacherRepo.GetAll()
}
func (s *TeacherService) Update(id uint64, update domain.TeacherUpdate) error {
	if update.ShortName != nil && *update.ShortName == "" {
		return apperror.ErrEmptyShortName
	}
	return s.teacherRepo.Update(id, update)
}
func (s *TeacherService) Delete(id uint64) error {
	return s.teacherRepo.Delete(id)
}
