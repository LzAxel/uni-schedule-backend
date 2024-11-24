package class

import (
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
)

type ClassService struct {
	repo         repository.ClassRepository
	entryRepo    repository.EntryRepository
	scheduleRepo repository.ScheduleRepository
}

func NewClassService(repo repository.ClassRepository, scheduleRepo repository.ScheduleRepository, entryRepo repository.EntryRepository) *ClassService {
	return &ClassService{repo: repo, scheduleRepo: scheduleRepo, entryRepo: entryRepo}
}

func (s *ClassService) AddClassWithEntry(dto domain.CreateClassWithEntryDTO) (uint64, error) {
	createdClassID, err := s.repo.Create(domain.CreateClassDTO{
		ScheduleID: dto.ScheduleID,
		SubjectID:  dto.SubjectID,
		TeacherID:  dto.TeacherID,
		ClassType:  dto.ClassType,
	})
	if err != nil {
		return 0, err
	}

	entryDTO := domain.CreateScheduleEntryDTO{
		Day:         dto.Day,
		ScheduleID:  dto.ScheduleID,
		ClassNumber: dto.ClassNumber,
		IsStatic:    dto.IsStatic,
	}
	if dto.Position == domain.ClassPositionEven {
		entryDTO.EvenClassID = &createdClassID
	} else {
		entryDTO.OddClassID = &createdClassID
	}

	return s.entryRepo.Create(entryDTO)
}

func (s *ClassService) Create(class domain.CreateClassDTO) (uint64, error) {
	entry, err := s.entryRepo.GetByID(class.EntryID)
	if err != nil {
		return 0, err
	}

	if entry.ScheduleID != class.ScheduleID {
		return 0, apperror.ErrDontHavePermission
	}

	if class.Position == domain.ClassPositionEven && entry.EvenClassID != nil {
		return 0, apperror.ErrClassInEntryAlreadySet
	}
	if class.Position == domain.ClassPositionOdd && entry.OddClassID != nil {
		return 0, apperror.ErrClassInEntryAlreadySet
	}

	createdClassID, err := s.repo.Create(class)
	if err != nil {
		return 0, err
	}

	var updateEntry = domain.UpdateScheduleEntryDTO{}
	if class.Position == domain.ClassPositionEven {
		updateEntry.EvenClassID = &createdClassID
	} else {
		updateEntry.OddClassID = &createdClassID
	}

	if err := s.entryRepo.Update(entry.ID, updateEntry); err != nil {
		return 0, err
	}

	return createdClassID, nil
}

func (s *ClassService) GetByID(id uint64) (domain.Class, error) {
	return s.repo.GetByID(id)
}

func (s *ClassService) GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.ClassView, domain.Pagination, error) {
	classes, total, err := s.repo.GetAllViews(scheduleID, limit, offset)
	if err != nil {
		return classes, domain.Pagination{}, err
	}

	return classes, domain.NewPagination(limit, offset, total), nil
}

func (s *ClassService) Update(userID uint64, id uint64, update domain.UpdateClassDTO) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}
	return s.repo.Update(id, update)
}

func (s *ClassService) Delete(userID uint64, id uint64) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *ClassService) isScheduleOwner(userID uint64, classID uint64) error {
	entry, err := s.repo.GetByID(classID)
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
