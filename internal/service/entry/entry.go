package entry

import (
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
)

type ScheduleEntryService struct {
	repo         repository.EntryRepository
	scheduleRepo repository.ScheduleRepository
}

func NewScheduleEntryService(repo repository.EntryRepository, scheduleRepo repository.ScheduleRepository) *ScheduleEntryService {
	return &ScheduleEntryService{repo: repo, scheduleRepo: scheduleRepo}
}

func (s *ScheduleEntryService) Create(entry domain.CreateScheduleEntryDTO) (uint64, error) {
	return s.repo.Create(entry)
}

func (s *ScheduleEntryService) GetByID(id uint64) (domain.ScheduleEntry, error) {
	return s.repo.GetByID(id)
}

func (s *ScheduleEntryService) Update(userID uint64, id uint64, update domain.UpdateScheduleEntryDTO) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}
	return s.repo.Update(id, update)
}

func (s *ScheduleEntryService) Delete(userID uint64, id uint64) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *ScheduleEntryService) isScheduleOwner(userID uint64, entryID uint64) error {
	entry, err := s.repo.GetByID(entryID)
	if err != nil {
		return err
	}
	schedule, err := s.scheduleRepo.GetByID(entry.ScheduleID)
	if err != nil {
		return err
	}
	if schedule.UserID != userID {
		return apperror.ErrDontHavePermission
	}

	return nil
}
