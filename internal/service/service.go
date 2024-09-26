package service

import (
	"uni-schedule-backend/internal/config"
	"uni-schedule-backend/internal/domain"
	tokenModel "uni-schedule-backend/internal/domain/auth/model"
	authservice "uni-schedule-backend/internal/domain/auth/service"
	lessonmodel "uni-schedule-backend/internal/domain/lesson/model"
	lessonservice "uni-schedule-backend/internal/domain/lesson/service"
	"uni-schedule-backend/internal/domain/schedule/model"
	scheduleservice "uni-schedule-backend/internal/domain/schedule/service"
	teachermodel "uni-schedule-backend/internal/domain/teacher/model"
	teacherservice "uni-schedule-backend/internal/domain/teacher/service"
	usermodel "uni-schedule-backend/internal/domain/user/model"
	userservice "uni-schedule-backend/internal/domain/user/service"
	"uni-schedule-backend/internal/jwt"
	"uni-schedule-backend/internal/repository"
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
	CreateSchedule(schedule model.Schedule) (domain.ID, error)
	GetScheduleByID(id domain.ID) (model.Schedule, error)
	UpdateSchedule(id domain.ID, update model.ScheduleUpdate) error
	DeleteSchedule(id domain.ID) error

	CreateSlot(slot model.ScheduleSlot) (domain.ID, error)
	GetSlotByID(id domain.ID) (model.ScheduleSlot, error)
	UpdateSlot(id domain.ID, update model.ScheduleSlotUpdate) error
	DeleteSlot(id domain.ID) error
}

type AuthService interface {
	Login(username, password string) (tokenModel.TokenPair, error)
	Register(username, password string) (tokenModel.TokenPair, error)
	RefreshToken(refreshToken string) (tokenModel.TokenPair, error)
	GetUserFromAccessToken(accessToken string) (usermodel.User, error)
}

type Service struct {
	Auth     AuthService
	User     UserService
	Teacher  TeacherService
	Lesson   LessonService
	Schedule ScheduleService
}

func NewService(repository *repository.Repository) *Service {
	cfg := config.GetConfig()
	jwtManager := jwt.NewJWTManager(cfg.JWT)

	return &Service{
		Auth:     authservice.NewAuthService(repository.User, repository.Token, jwtManager, cfg.AppConfig.PasswordSalt),
		User:     userservice.NewUserService(repository.User),
		Teacher:  teacherservice.NewTeacherService(repository.Teacher),
		Lesson:   lessonservice.NewLessonService(repository.Lesson),
		Schedule: scheduleservice.NewScheduleService(repository.Schedule, repository.ScheduleSlot),
	}
}
