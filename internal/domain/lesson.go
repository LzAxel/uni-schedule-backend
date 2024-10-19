package domain

type LessonType uint64

const (
	LessonTypeLecture  LessonType = iota
	LessonTypePractice LessonType = iota
	LessonTypeLab      LessonType = iota
)

type Lesson struct {
	ID         uint64     `db:"id"`
	Name       string     `db:"name"`
	Location   string     `db:"location"`
	TeacherID  uint64     `db:"teacher_id"`
	LessonType LessonType `db:"lesson_type"`
}

type LessonWithRelations struct {
	ID         uint64
	Name       string
	Location   string
	Teacher    Teacher
	LessonType LessonType
}

type LessonUpdate struct {
	Name       *string
	Location   *string
	TeacherID  *uint64
	LessonType *LessonType
}

type LessonCreate struct {
	Name       string
	Location   string
	TeacherID  uint64
	LessonType LessonType
}

type LessonView struct {
	ID         uint64      `json:"id"`
	Name       string      `json:"name"`
	Location   string      `json:"location"`
	Teacher    TeacherView `json:"teacher"`
	LessonType LessonType  `json:"lesson_type"`
}
