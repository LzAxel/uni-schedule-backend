package lesson

import (
	"uni-schedule-backend/internal/domain"
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

func (s *LessonService) Create(lesson domain.LessonCreate) (uint64, error) {
	if lesson.Name == "" {

	}
	return s.repo.Create(lesson)
}
func (s *LessonService) GetByID(id uint64) (domain.Lesson, error) {
	return s.repo.GetByID(id)
}
func (s *LessonService) Update(id uint64, update domain.LessonUpdate) error {
	return s.repo.Update(id, update)
}
func (s *LessonService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
