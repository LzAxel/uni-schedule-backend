package schedule

import (
	"errors"
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
)

type ScheduleService struct {
	scheduleRepo repository.ScheduleRepository
	entryRepo    repository.EntryRepository
}

func NewScheduleService(
	scheduleRepo repository.ScheduleRepository,
	entryRepo repository.EntryRepository,
) *ScheduleService {
	return &ScheduleService{
		scheduleRepo: scheduleRepo,
		entryRepo:    entryRepo,
	}
}

func (s *ScheduleService) Create(schedule domain.CreateScheduleDTO) (uint64, error) {
	if schedule.Slug == "" {
		return 0, apperror.ErrInvalidSlug
	}
	if schedule.Title == "" {
		return 0, apperror.ErrInvalidTitle
	}

	_, err := s.scheduleRepo.GetBySlug(schedule.Slug)
	if err == nil {
		return 0, apperror.ErrSlugAlreadyInUse
	}
	if !errors.Is(err, apperror.ErrNotFound) {
		return 0, err
	}

	return s.scheduleRepo.Create(schedule)
}

func (s *ScheduleService) GetByID(id uint64) (domain.ScheduleView, error) {
	schedule, err := s.scheduleRepo.GetByID(id)
	if err != nil {
		return domain.ScheduleView{}, err
	}
	entries, err := s.entryRepo.GetEntriesView(schedule.ID)
	if err != nil {
		return domain.ScheduleView{}, err
	}

	return schedule.ToView(entries), nil
}

func (s *ScheduleService) GetBySlug(slug string) (domain.ScheduleView, error) {
	schedule, err := s.scheduleRepo.GetBySlug(slug)
	if err != nil {
		return domain.ScheduleView{}, err
	}
	entries, err := s.entryRepo.GetEntriesView(schedule.ID)
	if err != nil {
		return domain.ScheduleView{}, err
	}

	return schedule.ToView(entries), nil
}

func (s *ScheduleService) GetMy(userID uint64, limit uint64, offset uint64) ([]domain.Schedule, domain.Pagination, error) {
	schedules, total, err := s.scheduleRepo.GetAll(limit, offset, domain.ScheduleGetAllFilters{
		UserID: &userID,
	})
	if err != nil {
		return schedules, domain.Pagination{}, err
	}

	return schedules, domain.NewPagination(limit, offset, total), nil
}

func (s *ScheduleService) Delete(userID uint64, id uint64) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}
	return s.scheduleRepo.Delete(id)
}

func (s *ScheduleService) Update(userID uint64, id uint64, update domain.UpdateScheduleDTO) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}
	return s.scheduleRepo.Update(id, update)
}

func (s *ScheduleService) isScheduleOwner(userID uint64, scheduleID uint64) error {
	schedule, err := s.scheduleRepo.GetByID(scheduleID)
	if err != nil {
		return err
	}
	if schedule.UserID != userID {
		return apperror.ErrDontHavePermission
	}

	return nil
}
