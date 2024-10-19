package schedule

import (
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
)

func (s *ScheduleService) CreateSlot(slot domain.ScheduleSlotCreate) (uint64, error) {
	if slot.Weekday >= 7 {
		return 0, apperror.ErrInvalidWeekdayNumber
	}
	return s.scheduleSlotRepo.Create(slot)
}
func (s *ScheduleService) GetSlotByID(id uint64) (domain.ScheduleSlot, error) {
	return s.scheduleSlotRepo.GetByID(id)
}
func (s *ScheduleService) GetAllSlotsByScheduleID(scheduleID uint64) ([]domain.ScheduleSlot, error) {
	return s.scheduleSlotRepo.GetAllSlotsByScheduleID(scheduleID)
}
func (s *ScheduleService) UpdateSlot(id uint64, update domain.ScheduleSlotUpdate) error {
	if update.Weekday != nil && *update.Weekday >= 7 {
		return apperror.ErrInvalidWeekdayNumber
	}
	return s.scheduleSlotRepo.Update(id, update)
}
func (s *ScheduleService) DeleteSlot(id uint64) error {
	return s.scheduleSlotRepo.Delete(id)
}
