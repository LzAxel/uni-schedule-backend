package service

import (
	"uni-schedule-backend/internal/config"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/jwt"
	"uni-schedule-backend/internal/repository"
	authservice "uni-schedule-backend/internal/service/auth"
	lessonservice "uni-schedule-backend/internal/service/lesson"
	scheduleservice "uni-schedule-backend/internal/service/schedule"
	teacherservice "uni-schedule-backend/internal/service/teacher"
	userservice "uni-schedule-backend/internal/service/user"
)

type UserService interface {
	Create(user domain.User) (domain.ID, error)
	GetByID(id domain.ID) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Update(id domain.ID, update domain.UserUpdateDTO) error
	Delete(id domain.ID) error
}

type TeacherService interface {
	Create(teacher domain.Teacher) (domain.ID, error)
	GetByID(id domain.ID) (domain.Teacher, error)
	GetAll() ([]domain.Teacher, error)
	Update(id domain.ID, update domain.TeacherUpdate) error
	Delete(id domain.ID) error
}

type LessonService interface {
	Create(lesson domain.Lesson) (domain.ID, error)
	GetByID(id domain.ID) (domain.Lesson, error)
	Update(id domain.ID, update domain.LessonUpdate) error
	Delete(id domain.ID) error
}

type ScheduleService interface {
	CreateSchedule(schedule domain.Schedule) (domain.ID, error)
	GetScheduleByID(id domain.ID) (domain.Schedule, error)
	UpdateSchedule(id domain.ID, update domain.ScheduleUpdate) error
	DeleteSchedule(id domain.ID) error

	CreateSlot(slot domain.ScheduleSlot) (domain.ID, error)
	GetSlotByID(id domain.ID) (domain.ScheduleSlot, error)
	UpdateSlot(id domain.ID, update domain.ScheduleSlotUpdate) error
	DeleteSlot(id domain.ID) error
}

type AuthService interface {
	Login(username, password string) (domain.TokenPair, error)
	Register(username, password string) (domain.TokenPair, error)
	RefreshToken(refreshToken string) (domain.TokenPair, error)
	GetUserFromAccessToken(accessToken string) (domain.User, error)
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
