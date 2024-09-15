package service

import (
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
	"uni-schedule-backend/internal/schedule/model"
)

type ScheduleService struct {
	scheduleRepo     repository.ScheduleRepository
	scheduleSlotRepo repository.ScheduleSlotRepository
}

func NewScheduleService(
	scheduleRepo repository.ScheduleRepository,
	scheduleSlotRepo repository.ScheduleSlotRepository,
) *ScheduleService {
	return &ScheduleService{
		scheduleRepo:     scheduleRepo,
		scheduleSlotRepo: scheduleSlotRepo,
	}
}

func (s *ScheduleService) CreateSchedule(schedule model.Schedule) (domain.ID, error) {
	return s.scheduleRepo.Create(schedule)
}
func (s *ScheduleService) GetScheduleByID(id domain.ID) (model.Schedule, error) {
	return s.scheduleRepo.GetByID(id)
}
func (s *ScheduleService) UpdateSchedule(id domain.ID, update model.ScheduleUpdate) error {
	return s.scheduleRepo.Update(id, update)
}
func (s *ScheduleService) DeleteSchedule(id domain.ID) error {
	return s.scheduleRepo.Delete(id)
}

func (s *ScheduleService) CreateSlot(slot model.ScheduleSlot) (domain.ID, error) {
	return s.scheduleSlotRepo.Create(slot)
}
func (s *ScheduleService) GetSlotByID(id domain.ID) (model.ScheduleSlot, error) {
	return s.scheduleSlotRepo.GetByID(id)
}
func (s *ScheduleService) UpdateSlot(id domain.ID, update model.ScheduleSlotUpdate) error {
	return s.scheduleSlotRepo.Update(id, update)
}
func (s *ScheduleService) DeleteSlot(id domain.ID) error {
	return s.scheduleSlotRepo.Delete(id)
}
