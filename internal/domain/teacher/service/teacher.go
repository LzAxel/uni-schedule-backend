package service

import (
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/teacher/model"
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

func (s *TeacherService) Create(teacher model.Teacher) (domain.ID, error) {
	return s.Create(teacher)
}
func (s *TeacherService) GetByID(id domain.ID) (model.Teacher, error) {
	return s.GetByID(id)
}
func (s *TeacherService) GetAll() ([]model.Teacher, error) {
	return s.GetAll()
}
func (s *TeacherService) Update(id domain.ID, update model.TeacherUpdate) error {
	return s.Update(id, update)
}
func (s *TeacherService) Delete(id domain.ID) error {
	return s.Delete(id)
}
