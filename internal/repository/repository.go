package repository

import (
	"uni-schedule-backend/internal/domain"
)

type UserRepository interface {
	Create(user domain.UserCreate) (domain.ID, error)
	GetByID(id domain.ID) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Update(id domain.ID, update domain.UserUpdateDTO) error
	Delete(id domain.ID) error
}

type TeacherRepository interface {
	Create(teacher domain.Teacher) (domain.ID, error)
	GetByID(id domain.ID) (domain.Teacher, error)
	GetAll() ([]domain.Teacher, error)
	Update(id domain.ID, update domain.TeacherUpdate) error
	Delete(id domain.ID) error
}

type LessonRepository interface {
	Create(lesson domain.Lesson) (domain.ID, error)
	GetByID(id domain.ID) (domain.Lesson, error)
	Update(id domain.ID, update domain.LessonUpdate) error
	Delete(id domain.ID) error
}

type ScheduleRepository interface {
	Create(schedule domain.Schedule) (domain.ID, error)
	GetByID(id domain.ID) (domain.Schedule, error)
	Update(id domain.ID, update domain.ScheduleUpdate) error
	Delete(id domain.ID) error
}

type ScheduleSlotRepository interface {
	Create(slot domain.ScheduleSlot) (domain.ID, error)
	GetByID(id domain.ID) (domain.ScheduleSlot, error)
	Update(id domain.ID, update domain.ScheduleSlotUpdate) error
	Delete(id domain.ID) error
}

type TokenRepository interface {
	CreateOrUpdate(token domain.RefreshToken) error
	GetByUserID(userID domain.ID) (domain.RefreshToken, error)
	Delete(userID domain.ID) error
}

type Repository struct {
	Token        TokenRepository
	User         UserRepository
	Teacher      TeacherRepository
	Lesson       LessonRepository
	Schedule     ScheduleRepository
	ScheduleSlot ScheduleSlotRepository
}
