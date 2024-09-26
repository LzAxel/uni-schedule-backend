package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	tokenrepo "uni-schedule-backend/internal/domain/auth/repository"
	lessonrepo "uni-schedule-backend/internal/domain/lesson/repository"
	repository2 "uni-schedule-backend/internal/domain/schedule/repository"
	teacherrepo "uni-schedule-backend/internal/domain/teacher/repository"
	userrepo "uni-schedule-backend/internal/domain/user/repository"
	"uni-schedule-backend/internal/repository"
)

func NewDBConnection(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresRepository(db *sqlx.DB) *repository.Repository {
	return &repository.Repository{
		Token:        tokenrepo.NewTokenRepo(db),
		User:         userrepo.NewUserRepo(db),
		Teacher:      teacherrepo.NewTeacherRepo(db),
		Lesson:       lessonrepo.NewLessonRepo(db),
		Schedule:     repository2.NewScheduleRepo(db),
		ScheduleSlot: repository2.NewScheduleSlotRepo(db),
	}
}
