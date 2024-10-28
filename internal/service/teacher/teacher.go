package teacher

import (
	"uni-schedule-backend/internal/apperror"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/repository"
)

type TeacherService struct {
	repo         repository.TeacherRepository
	scheduleRepo repository.ScheduleRepository
}

func NewTeacherService(repo repository.TeacherRepository, scheduleRepo repository.ScheduleRepository) *TeacherService {
	return &TeacherService{repo: repo, scheduleRepo: scheduleRepo}
}

func (s *TeacherService) Create(teacher domain.TeacherCreateDTO) (uint64, error) {
	return s.repo.Create(teacher)
}

func (s *TeacherService) GetByID(id uint64) (domain.Teacher, error) {
	return s.repo.GetByID(id)
}
func (s *TeacherService) GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Teacher, domain.Pagination, error) {
	teachers, total, err := s.repo.GetAll(scheduleID, limit, offset)
	if err != nil {
		return teachers, domain.Pagination{}, err
	}

	return teachers, domain.NewPagination(limit, offset, total), nil
}

func (s *TeacherService) Update(userID uint64, id uint64, update domain.TeacherUpdateDTO) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}

	return s.repo.Update(id, update)
}

func (s *TeacherService) Delete(userID uint64, id uint64) error {
	if err := s.isScheduleOwner(userID, id); err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *TeacherService) isScheduleOwner(userID uint64, teacherID uint64) error {
	teacher, err := s.repo.GetByID(teacherID)
	if err != nil {
		return err
	}
	schedule, err := s.scheduleRepo.GetByID(teacher.ScheduleID)
	if err != nil {
		return err
	}
	if schedule.UserID != userID {
		return apperror.ErrDontHavePermission
	}

	return nil
}
