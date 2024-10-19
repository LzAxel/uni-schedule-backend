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
	Create(teacher domain.TeacherCreate) (uint64, error)
	GetByID(id uint64) (domain.Teacher, error)
	GetAll() ([]domain.Teacher, error)
	Update(id uint64, update domain.TeacherUpdate) error
	Delete(id uint64) error
}

type LessonRepository interface {
	Create(lesson domain.LessonCreate) (uint64, error)
	GetByID(id uint64) (domain.Lesson, error)
	GetWithRelationsByID(id uint64) (domain.LessonView, error)
	Update(id uint64, update domain.LessonUpdate) error
	Delete(id uint64) error
}

type ScheduleRepository interface {
	Create(schedule domain.ScheduleCreate) (uint64, error)
	GetByID(id uint64) (domain.Schedule, error)
	GetBySlug(slug string) (domain.Schedule, error)
	Update(id uint64, update domain.ScheduleUpdate) error
	Delete(id uint64) error
}

type ScheduleSlotRepository interface {
	Create(slot domain.ScheduleSlotCreate) (uint64, error)
	GetByID(id uint64) (domain.ScheduleSlot, error)
	GetAllSlotsByScheduleID(scheduleID uint64) ([]domain.ScheduleSlot, error)
	Update(id uint64, update domain.ScheduleSlotUpdate) error
	Delete(id uint64) error
}

type TokenRepository interface {
	CreateOrUpdate(token domain.RefreshToken) error
	GetByUserID(userID uint64) (domain.RefreshToken, error)
	Delete(userID uint64) error
}

type Repository struct {
	Token        TokenRepository
	User         UserRepository
	Teacher      TeacherRepository
	Lesson       LessonRepository
	Schedule     ScheduleRepository
	ScheduleSlot ScheduleSlotRepository
}
