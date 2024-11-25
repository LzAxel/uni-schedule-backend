package service

import (
	"uni-schedule-backend/internal/config"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/jwt"
	"uni-schedule-backend/internal/repository"
	authservice "uni-schedule-backend/internal/service/auth"
	classservice "uni-schedule-backend/internal/service/class"
	scheduleservice "uni-schedule-backend/internal/service/schedule"
	subjectservice "uni-schedule-backend/internal/service/subject"
	teacherservice "uni-schedule-backend/internal/service/teacher"
	userservice "uni-schedule-backend/internal/service/user"
)

type UserService interface {
	Create(user domain.User) (uint64, error)
	GetByID(id uint64) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Update(userID uint64, id uint64, update domain.UserUpdateDTO) error
	Delete(id uint64) error
}

type TeacherService interface {
	Create(teacher domain.TeacherCreateDTO) (uint64, error)
	GetByID(id uint64) (domain.Teacher, error)
	GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Teacher, domain.Pagination, error)
	Update(userID uint64, id uint64, update domain.TeacherUpdateDTO) error
	Delete(userID uint64, id uint64) error
}

type ScheduleService interface {
	Create(schedule domain.CreateScheduleDTO) (uint64, error)
	GetByID(id uint64) (domain.ScheduleView, error)
	GetBySlug(slug string) (domain.ScheduleView, error)
	GetMy(userID uint64, limit uint64, offset uint64) ([]domain.Schedule, domain.Pagination, error)
	Update(userID uint64, id uint64, update domain.UpdateScheduleDTO) error
	Delete(userID uint64, id uint64) error
}

type ClassService interface {
	Create(class domain.CreateClassDTO) (uint64, error)
	GetByID(id uint64) (domain.Class, error)
	GetAll(scheduleID uint64) ([]domain.ClassView, error)
	Update(userID uint64, id uint64, update domain.UpdateClassDTO) error
	Delete(userID uint64, id uint64) error
}

type AuthService interface {
	Login(username, password string) (domain.TokenPair, error)
	Register(username, password string) (domain.TokenPair, error)
	RefreshToken(refreshToken string) (domain.TokenPair, error)
	GetUserFromAccessToken(accessToken string) (domain.User, error)
}

type SubjectService interface {
	Create(subject domain.CreateSubjectDTO) (uint64, error)
	GetByID(id uint64) (domain.Subject, error)
	GetAll(scheduleID uint64, limit uint64, offset uint64) ([]domain.Subject, domain.Pagination, error)
	Update(userID uint64, id uint64, update domain.UpdateSubjectDTO) error
	Delete(userID uint64, id uint64) error
}

type Service struct {
	Auth     AuthService
	User     UserService
	Teacher  TeacherService
	Schedule ScheduleService
	Subject  SubjectService
	Class    ClassService
}

func NewService(repository *repository.Repository) *Service {
	cfg := config.GetConfig()
	jwtManager := jwt.NewJWTManager(cfg.JWT)

	return &Service{
		Auth:     authservice.NewAuthService(repository.User, repository.Token, jwtManager, cfg.AppConfig.PasswordSalt),
		User:     userservice.NewUserService(repository.User),
		Teacher:  teacherservice.NewTeacherService(repository.Teacher, repository.Schedule),
		Schedule: scheduleservice.NewScheduleService(repository.Schedule, repository.Class),
		Subject:  subjectservice.NewSubjectService(repository.Subject, repository.Schedule),
		Class:    classservice.NewClassService(repository.Class, repository.Schedule),
	}
}
