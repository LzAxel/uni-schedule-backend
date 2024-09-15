package repository

import (
	"uni-schedule-backend/internal/domain"
	lessonmodel "uni-schedule-backend/internal/lesson/model"
	schedulemodel "uni-schedule-backend/internal/schedule/model"
	teachermodel "uni-schedule-backend/internal/teacher/model"
	usermodel "uni-schedule-backend/internal/user/model"
)

type UserRepository interface {
	Create(user usermodel.User) (domain.ID, error)
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
	Create(schedule schedulemodel.Schedule) (domain.ID, error)
	GetByID(id domain.ID) (schedulemodel.Schedule, error)
	Update(id domain.ID, update schedulemodel.ScheduleUpdate) error
	Delete(id domain.ID) error
}

type ScheduleSlotRepository interface {
	Create(slot schedulemodel.ScheduleSlot) (domain.ID, error)
	GetByID(id domain.ID) (schedulemodel.ScheduleSlot, error)
	Update(id domain.ID, update schedulemodel.ScheduleSlotUpdate) error
	Delete(id domain.ID) error
}

type Repository struct {
	User         UserRepository
	Teacher      TeacherRepository
	Lesson       LessonRepository
	Schedule     ScheduleRepository
	ScheduleSlot ScheduleSlotRepository
}
