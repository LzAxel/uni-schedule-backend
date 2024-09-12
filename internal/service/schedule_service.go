package service

import (
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
)

type ScheduleService struct {
	repo *repository.ScheduleRepository
}

func NewScheduleService(repo *repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (s *ScheduleService) GetScheduleByDay(dayOfWeek string) (*domain.FinalSchedule, error) {
	return s.repo.GetScheduleByDay(dayOfWeek)
}

func (s *ScheduleService) AddScheduleSlot(day string, slot int, lesson *domain.LessonInfo, weekType string) error {
	return s.repo.AddScheduleSlot(day, slot, lesson, weekType)
}

func (s *ScheduleService) GetFullSchedule() (map[string]*domain.FinalSchedule, error) {
	return s.repo.GetFullSchedule()
}
