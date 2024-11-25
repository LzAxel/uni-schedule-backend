package psql

import (
	"uni-schedule-backend/internal/repository"
	tokenrepo "uni-schedule-backend/internal/repository/auth"
	classrepo "uni-schedule-backend/internal/repository/class"
	"uni-schedule-backend/internal/repository/schedule"
	subjectrepo "uni-schedule-backend/internal/repository/subject"
	teacherrepo "uni-schedule-backend/internal/repository/teacher"
	userrepo "uni-schedule-backend/internal/repository/user"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
		Token:    tokenrepo.NewTokenRepo(db),
		User:     userrepo.NewUserRepo(db),
		Teacher:  teacherrepo.NewTeacherRepo(db),
		Subject:  subjectrepo.NewSubjectRepo(db),
		Schedule: schedule.NewScheduleRepo(db),
		Class:    classrepo.NewClassRepo(db),
	}
}
