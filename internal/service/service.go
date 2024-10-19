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
	Create(user domain.User) (uint64, error)
	GetByID(id uint64) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Update(id uint64, update domain.UserUpdateDTO) error
	Delete(id uint64) error
}

type TeacherService interface {
	Create(teacher domain.TeacherCreate) (uint64, error)
	GetByID(id uint64) (domain.Teacher, error)
	GetAll() ([]domain.Teacher, error)
	Update(id uint64, update domain.TeacherUpdate) error
	Delete(id uint64) error
}

type LessonService interface {
	Create(lesson domain.LessonCreate) (uint64, error)
	GetByID(id uint64) (domain.Lesson, error)
	Update(id uint64, update domain.LessonUpdate) error
	Delete(id uint64) error
}

type ScheduleService interface {
	CreateSchedule(schedule domain.ScheduleCreate) (uint64, error)
	GetScheduleBySlug(slug string) (domain.ScheduleView, error)
	UpdateSchedule(id uint64, update domain.ScheduleUpdate) error
	DeleteSchedule(id uint64) error

	CreateSlot(slot domain.ScheduleSlotCreate) (uint64, error)
	GetSlotByID(id uint64) (domain.ScheduleSlot, error)
	UpdateSlot(id uint64, update domain.ScheduleSlotUpdate) error
	DeleteSlot(id uint64) error
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
		Schedule: scheduleservice.NewScheduleService(repository.Schedule, repository.ScheduleSlot, repository.Lesson, repository.Teacher),
	}
}
