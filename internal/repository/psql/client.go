package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"uni-schedule-backend/internal/repository"
	tokenrepo "uni-schedule-backend/internal/repository/auth"
	lessonrepo "uni-schedule-backend/internal/repository/lesson"
	"uni-schedule-backend/internal/repository/schedule"
	teacherrepo "uni-schedule-backend/internal/repository/teacher"
	userrepo "uni-schedule-backend/internal/repository/user"
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
		Schedule:     schedule.NewScheduleRepo(db),
		ScheduleSlot: schedule.NewScheduleSlotRepo(db),
	}
}
