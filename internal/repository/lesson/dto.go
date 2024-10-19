package lesson

import "uni-schedule-backend/internal/domain"

type lessonWithRelations struct {
	LessonID   uint64            `db:"lesson_id"`
	Name       string            `db:"lesson_name"`
	Location   string            `db:"lesson_location"`
	LessonType domain.LessonType `db:"lesson_type"`

	TeacherID uint64 `db:"teacher_id"`
	ShortName string `db:"teacher_short_name"`
	FullName  string `db:"teacher_full_name"`
}

func (l lessonWithRelations) ToDomain() domain.LessonView {
	return domain.LessonView{
		ID:         l.LessonID,
		Name:       l.Name,
		Location:   l.Location,
		LessonType: l.LessonType,
		Teacher: domain.TeacherView{
			ID:        l.TeacherID,
			ShortName: l.ShortName,
			FullName:  l.FullName,
		},
	}
}
