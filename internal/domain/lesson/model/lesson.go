package model

import "uni-schedule-backend/internal/domain"

type LessonType string

const (
	LessonTypeLecture  LessonType = "lecture"
	LessonTypePractice LessonType = "practice"
	LessonTypeLab      LessonType = "lab"
)

type Lesson struct {
	ID         domain.ID
	Name       string
	Location   string
	TeacherID  int
	LessonType LessonType
}

type LessonUpdate struct {
	Name       *string
	Location   *string
	TeacherID  *domain.ID
	LessonType *LessonType
}
