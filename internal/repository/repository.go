package repository

import (
	"uni-schedule-backend/internal/domain"
)

type UserRepository interface {
	Create(user domain.UserCreate) (uint64, error)
	GetByID(id uint64) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Update(id uint64, update domain.UserUpdateDTO) error
	Delete(id uint64) error
}

type TeacherRepository interface {
	Create(teacher domain.TeacherCreateDTO) (uint64, error)
	GetByID(id uint64) (domain.Teacher, error)
	GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Teacher, uint64, error)
	Update(id uint64, update domain.TeacherUpdateDTO) error
	Delete(id uint64) error
}

type ScheduleRepository interface {
	Create(schedule domain.CreateScheduleDTO) (uint64, error)
	GetByID(id uint64) (domain.Schedule, error)
	GetBySlug(slug string) (domain.Schedule, error)
	GetAll(limit uint64, offset uint64, filters domain.ScheduleGetAllFilters) ([]domain.Schedule, uint64, error)
	Update(id uint64, update domain.UpdateScheduleDTO) error
	Delete(id uint64) error
}

type TokenRepository interface {
	CreateOrUpdate(token domain.RefreshToken) error
	GetByUserID(userID uint64) (domain.RefreshToken, error)
	Delete(userID uint64) error
}

type SubjectRepository interface {
	Create(subject domain.CreateSubjectDTO) (uint64, error)
	GetByID(id uint64) (domain.Subject, error)
	GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Subject, uint64, error)
	Update(id uint64, update domain.UpdateSubjectDTO) error
	Delete(id uint64) error
}

type ClassRepository interface {
	Create(class domain.CreateClassDTO) (uint64, error)
	CreateOrSplit(class domain.CreateClassDTO) (uint64, error)
	GetByID(id uint64) (domain.Class, error)
	GetAllByDayAndNumber(scheduleID uint64, day domain.Day, number uint64) ([]domain.Class, error)
	GetAllViews(scheduleID uint64) (domain.ClassViews, uint64, error)
	GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Class, uint64, error)
	Update(id uint64, update domain.UpdateClassDTO) error
	UpdateOrSwitch(id uint64, scheduleID uint64, update domain.UpdateClassDTO) error
	Delete(id uint64) error
}

type Repository struct {
	Token    TokenRepository
	User     UserRepository
	Subject  SubjectRepository
	Teacher  TeacherRepository
	Class    ClassRepository
	Schedule ScheduleRepository
}
