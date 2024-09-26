package repository

import (
	"uni-schedule-backend/internal/domain"
	tokenmodel "uni-schedule-backend/internal/domain/auth/model"
	lessonmodel "uni-schedule-backend/internal/domain/lesson/model"
	"uni-schedule-backend/internal/domain/schedule/model"
	teachermodel "uni-schedule-backend/internal/domain/teacher/model"
	usermodel "uni-schedule-backend/internal/domain/user/model"
)

type UserRepository interface {
	Create(user usermodel.UserCreate) (domain.ID, error)
	GetByID(id domain.ID) (usermodel.User, error)
	GetByUsername(username string) (usermodel.User, error)
	Update(id domain.ID, update usermodel.UserUpdateDTO) error
	Delete(id domain.ID) error
}

type TeacherRepository interface {
	Create(teacher teachermodel.Teacher) (domain.ID, error)
	GetByID(id domain.ID) (teachermodel.Teacher, error)
	GetAll() ([]teachermodel.Teacher, error)
	Update(id domain.ID, update teachermodel.TeacherUpdate) error
	Delete(id domain.ID) error
}

type LessonRepository interface {
	Create(lesson lessonmodel.Lesson) (domain.ID, error)
	GetByID(id domain.ID) (lessonmodel.Lesson, error)
	Update(id domain.ID, update lessonmodel.LessonUpdate) error
	Delete(id domain.ID) error
}

type ScheduleRepository interface {
	Create(schedule model.Schedule) (domain.ID, error)
	GetByID(id domain.ID) (model.Schedule, error)
	Update(id domain.ID, update model.ScheduleUpdate) error
	Delete(id domain.ID) error
}

type ScheduleSlotRepository interface {
	Create(slot model.ScheduleSlot) (domain.ID, error)
	GetByID(id domain.ID) (model.ScheduleSlot, error)
	Update(id domain.ID, update model.ScheduleSlotUpdate) error
	Delete(id domain.ID) error
}

type TokenRepository interface {
	CreateOrUpdate(token tokenmodel.RefreshToken) error
	GetByUserID(userID domain.ID) (tokenmodel.RefreshToken, error)
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
