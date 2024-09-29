package domain

type LessonType string

const (
	LessonTypeLecture  LessonType = "lecture"
	LessonTypePractice LessonType = "practice"
	LessonTypeLab      LessonType = "lab"
)

type Lesson struct {
	ID         ID
	Name       string
	Location   string
	TeacherID  int
	LessonType LessonType
}

type LessonUpdate struct {
	Name       *string
	Location   *string
	TeacherID  *ID
	LessonType *LessonType
}

type LessonView struct {
	ID         ID          `json:"id"`
	Name       string      `json:"name"`
	Location   string      `json:"location"`
	Teacher    TeacherView `json:"teacher"`
	LessonType LessonType  `json:"lesson_type"`
}
