package psql

import (
	"database/sql"
	lessonrepo "uni-schedule-backend/internal/lesson/repository"
	"uni-schedule-backend/internal/repository"
	schedulerepo "uni-schedule-backend/internal/schedule/repository"
	teacherrepo "uni-schedule-backend/internal/teacher/repository"
	userrepo "uni-schedule-backend/internal/user/repository"
)

func NewPostgresRepository(db *sql.DB) *repository.Repository {
	return &repository.Repository{
		User:         userrepo.NewUserRepo(db),
		Teacher:      teacherrepo.NewTeacherRepo(db),
		Lesson:       lessonrepo.NewLessonRepo(db),
		Schedule:     schedulerepo.NewScheduleRepo(db),
		ScheduleSlot: schedulerepo.NewScheduleSlotRepo(db),
	}
}
