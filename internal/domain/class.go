package domain

type ClassPosition string

const (
	ClassPositionEven ClassPosition = "even"
	ClassPositionOdd  ClassPosition = "odd"
)

type Class struct {
	ID         uint64    `json:"id" db:"id"`
	ScheduleID uint64    `json:"schedule_id" db:"schedule_id"`
	SubjectID  uint64    `json:"subject_id" db:"subject_id"`
	TeacherID  uint64    `json:"teacher_id" db:"teacher_id"`
	ClassType  ClassType `json:"class_type" db:"class_type"`
}

type CreateClassWithEntryDTO struct {
	ScheduleID  uint64        `json:"schedule_id" binding:"required"`
	SubjectID   uint64        `json:"subject_id" binding:"required"`
	TeacherID   uint64        `json:"teacher_id" binding:"required"`
	ClassType   ClassType     `json:"class_type" binding:"required"`
	ClassNumber int           `json:"class_number" binding:"required"`
	IsStatic    bool          `json:"is_static" binding:"required"`
	Position    ClassPosition `json:"position" binding:"required"`
	Day         Day           `json:"day" binding:"required"`
}

type CreateClassDTO struct {
	ScheduleID uint64        `json:"schedule_id" binding:"required"`
	EntryID    uint64        `json:"entry_id" binding:"required"`
	SubjectID  uint64        `json:"subject_id" binding:"required"`
	TeacherID  uint64        `json:"teacher_id" binding:"required"`
	Position   ClassPosition `json:"position" binding:"required"`
	ClassType  ClassType     `json:"class_type" binding:"required"`
}

type UpdateClassDTO struct {
	SubjectID *uint64    `json:"subject_id,omitempty"`
	TeacherID *uint64    `json:"teacher_id,omitempty"`
	ClassType *ClassType `json:"class_type,omitempty"`
}

type ClassView struct {
	ID        uint64      `json:"id"`
	Subject   SubjectView `json:"subject"`
	Teacher   TeacherView `json:"teacher"`
	ClassType ClassType   `json:"class_type"`
}

func (c Class) ToView(subject SubjectView, teacher TeacherView) ClassView {
	return ClassView{
		ID:        c.ID,
		Subject:   subject,
		Teacher:   teacher,
		ClassType: c.ClassType,
	}
}
