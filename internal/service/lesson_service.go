package service

import (
	"uni-schedule-backend/internal/lesson/model"
	"uni-schedule-backend/internal/repository"
)

type LessonService struct {
	repo repository.LessonRepository
}

func NewLessonService(repo repository.LessonRepository) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) GetLessonByID(id int) (*model.Lesson, error) {
	return s.repo.GetLessonByID(id)
}

func (s *LessonService) AddLesson(lesson model.Lesson) (int, error) {
	return s.repo.AddLesson(lesson)
}

func (s *LessonService) GetAllLessons() ([]model.Lesson, error) {
	return s.repo.GetAllLessons()
}
