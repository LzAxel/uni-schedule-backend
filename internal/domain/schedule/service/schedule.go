package service

import (
	"uni-schedule-backend/internal/domain"
	model2 "uni-schedule-backend/internal/domain/schedule/model"
	"uni-schedule-backend/internal/repository"
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

func (s *ScheduleService) CreateSchedule(schedule model2.Schedule) (domain.ID, error) {
	return s.scheduleRepo.Create(schedule)
}
func (s *ScheduleService) GetScheduleByID(id domain.ID) (model2.Schedule, error) {
	return s.scheduleRepo.GetByID(id)
}
func (s *ScheduleService) UpdateSchedule(id domain.ID, update model2.ScheduleUpdate) error {
	return s.scheduleRepo.Update(id, update)
}
func (s *ScheduleService) DeleteSchedule(id domain.ID) error {
	return s.scheduleRepo.Delete(id)
}

func (s *ScheduleService) CreateSlot(slot model2.ScheduleSlot) (domain.ID, error) {
	return s.scheduleSlotRepo.Create(slot)
}
func (s *ScheduleService) GetSlotByID(id domain.ID) (model2.ScheduleSlot, error) {
	return s.scheduleSlotRepo.GetByID(id)
}
func (s *ScheduleService) UpdateSlot(id domain.ID, update model2.ScheduleSlotUpdate) error {
	return s.scheduleSlotRepo.Update(id, update)
}
func (s *ScheduleService) DeleteSlot(id domain.ID) error {
	return s.scheduleSlotRepo.Delete(id)
}
