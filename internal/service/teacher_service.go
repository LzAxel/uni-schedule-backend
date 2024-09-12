package service

import (
	"uni-schedule-backend/internal/repository"
	"uni-schedule-backend/internal/teacher/model"
)

type TeacherService struct {
	repo repository.TeacherRepository
}

func NewTeacherService(repo repository.TeacherRepository) *TeacherService {
	return &TeacherService{repo: repo}
}

func (s *TeacherService) GetAllTeachers() ([]model.Teacher, error) {
	return s.repo.GetAllTeachers()
}

func (s *TeacherService) AddTeacher(teacher model.Teacher) error {
	return s.repo.AddTeacher(teacher)
}
