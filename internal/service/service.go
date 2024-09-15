package service

import (
	"uni-schedule-backend/internal/domain"
	lessonmodel "uni-schedule-backend/internal/lesson/model"
	"uni-schedule-backend/internal/repository"
	schedulemodel "uni-schedule-backend/internal/schedule/model"
	teachermodel "uni-schedule-backend/internal/teacher/model"
	usermodel "uni-schedule-backend/internal/user/model"

	lessonservice "uni-schedule-backend/internal/lesson/service"
	scheduleservice "uni-schedule-backend/internal/schedule/service"
	teacherservice "uni-schedule-backend/internal/teacher/service"
	userservice "uni-schedule-backend/internal/user/service"
)

type UserService interface {
	Create(user usermodel.User) (domain.ID, error)
	GetByID(id domain.ID) (usermodel.User, error)
	GetByUsername(username string) (usermodel.User, error)
	Update(id domain.ID, update usermodel.UserUpdateDTO) error
	Delete(id domain.ID) error
}

type TeacherService interface {
	Create(teacher teachermodel.Teacher) (domain.ID, error)
	GetByID(id domain.ID) (teachermodel.Teacher, error)
	GetAll() ([]teachermodel.Teacher, error)
	Update(id domain.ID, update teachermodel.TeacherUpdate) error
	Delete(id domain.ID) error
}

type LessonService interface {
	Create(lesson lessonmodel.Lesson) (domain.ID, error)
	GetByID(id domain.ID) (lessonmodel.Lesson, error)
	Update(id domain.ID, update lessonmodel.LessonUpdate) error
	Delete(id domain.ID) error
}

type ScheduleService interface {
	CreateSchedule(schedule schedulemodel.Schedule) (domain.ID, error)
	GetScheduleByID(id domain.ID) (schedulemodel.Schedule, error)
	UpdateSchedule(id domain.ID, update schedulemodel.ScheduleUpdate) error
	DeleteSchedule(id domain.ID) error

	CreateSlot(slot schedulemodel.ScheduleSlot) (domain.ID, error)
	GetSlotByID(id domain.ID) (schedulemodel.ScheduleSlot, error)
	UpdateSlot(id domain.ID, update schedulemodel.ScheduleSlotUpdate) error
	DeleteSlot(id domain.ID) error
}

type Service struct {
	User     UserService
	Teacher  TeacherService
	Lesson   LessonService
	Schedule ScheduleService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:     userservice.NewUserService(repository.User),
		Teacher:  teacherservice.NewTeacherService(repository.Teacher),
		Lesson:   lessonservice.NewLessonService(repository.Lesson),
		Schedule: scheduleservice.NewScheduleService(repository.Schedule, repository.ScheduleSlot),
	}
}
