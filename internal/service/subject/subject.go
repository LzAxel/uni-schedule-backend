package subject

import (
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
)

type SubjectService struct {
	repo         repository.SubjectRepository
	scheduleRepo repository.ScheduleRepository
}

func NewSubjectService(repo repository.SubjectRepository, scheduleRepo repository.ScheduleRepository) *SubjectService {
	return &SubjectService{repo: repo, scheduleRepo: scheduleRepo}
}

func (s *SubjectService) Create(subject domain.CreateSubjectDTO) (uint64, error) {
	return s.repo.Create(subject)
}

func (s *SubjectService) GetByID(id uint64) (domain.Subject, error) {
	return s.repo.GetByID(id)
}
func (s *SubjectService) GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Subject, domain.Pagination, error) {
	subjects, total, err := s.repo.GetAll(scheduleID, limit, offset)
	if err != nil {
		return subjects, domain.Pagination{}, err
	}

	return subjects, domain.NewPagination(limit, offset, total), nil
}

func (s *SubjectService) Update(userID uint64, id uint64, update domain.UpdateSubjectDTO) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}
	return s.repo.Update(id, update)
}

func (s *SubjectService) Delete(userID uint64, id uint64) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *SubjectService) isScheduleOwner(userID uint64, subjectID uint64) error {
	subject, err := s.repo.GetByID(subjectID)
	if err != nil {
		return err
	}
	schedule, err := s.scheduleRepo.GetByID(subject.ScheduleID)
	if err != nil {
		return err
	}
	if schedule.UserID != userID {
		return apperror.ErrDontHavePermission
	}

	return nil
}
