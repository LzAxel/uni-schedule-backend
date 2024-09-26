package view

import (
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/domain/lesson/model"
	view2 "uni-schedule-backend/internal/domain/teacher/view"
)

type LessonView struct {
	ID         domain.ID         `json:"id"`
	Name       string            `json:"name"`
	Location   string            `json:"location"`
	Teacher    view2.TeacherView `json:"teacher"`
	LessonType model.LessonType  `json:"lesson_type"`
}
