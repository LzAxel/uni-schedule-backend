package service

import (
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/lesson/model"
	"uni-schedule-backend/internal/repository"
)

type LessonService struct {
	repo repository.LessonRepository
}

func NewLessonService(repo repository.LessonRepository) *LessonService {
	return &LessonService{
		repo: repo,
	}
}

func (s *LessonService) Create(lesson model.Lesson) (domain.ID, error) {
	return s.repo.Create(lesson)
}
func (s *LessonService) GetByID(id domain.ID) (model.Lesson, error) {
	return s.repo.GetByID(id)
}
func (s *LessonService) Update(id domain.ID, update model.LessonUpdate) error {
	return s.repo.Update(id, update)
}
func (s *LessonService) Delete(id domain.ID) error {
	return s.repo.Delete(id)
}
