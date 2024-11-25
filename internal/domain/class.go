package domain

import "fmt"

type ClassPosition string

const (
	ClassPositionEven ClassPosition = "even"
	ClassPositionOdd  ClassPosition = "odd"
)

type ClassType string

const (
	ClassTypeLecture  ClassType = "lecture"
	ClassTypePractice ClassType = "practice"
	ClassTypeLab      ClassType = "lab"
	ClassTypeCombined ClassType = "combined"
)

type Class struct {
	ID         uint64    `json:"id" db:"id"`
	ScheduleID uint64    `json:"schedule_id" db:"schedule_id"`
	SubjectID  uint64    `json:"subject_id" db:"subject_id"`
	TeacherID  uint64    `json:"teacher_id" db:"teacher_id"`
	ClassType  ClassType `json:"class_type" db:"class_type"`
	DayOfWeek  Day       `json:"day_of_week" db:"day_of_week"`
	Number     uint64    `json:"number" db:"class_number"`
	EvenWeek   *bool     `json:"even_week" db:"even_week"`
}

type CreateClassDTO struct {
	ScheduleID uint64    `json:"schedule_id" binding:"required"`
	SubjectID  uint64    `json:"subject_id" binding:"required"`
	TeacherID  uint64    `json:"teacher_id" binding:"required"`
	ClassType  ClassType `json:"class_type" binding:"required"`
	DayOfWeek  Day       `json:"day_of_week" binding:"required"`
	Number     uint64    `json:"number" binding:"required"`
	EvenWeek   *bool     `json:"even_week,omitempty"`
}

type UpdateClassDTO struct {
	TeacherID uint64    `json:"teacher_id" binding:"required"`
	SubjectID uint64    `json:"subject_id" binding:"required"`
	ClassType ClassType `json:"class_type" binding:"required"`
	DayOfWeek Day       `json:"day_of_week" binding:"required"`
	Number    uint64    `json:"number" binding:"required"`
	EvenWeek  *bool     `json:"even_week" binding:"required"`
}

type ClassView struct {
	ID        uint64      `json:"id"`
	Subject   SubjectView `json:"subject"`
	Teacher   TeacherView `json:"teacher"`
	ClassType ClassType   `json:"class_type"`
	DayOfWeek Day         `json:"day_of_week"`
	Number    uint64      `json:"number"`
	EvenWeek  *bool       `json:"even_week"`
}

type NumberGroupedClassesView map[uint64]ClassViews
type DayGroupedClassesView map[Day]NumberGroupedClassesView

func (c Class) ToView(subject SubjectView, teacher TeacherView) ClassView {
	return ClassView{
		ID:        c.ID,
		Subject:   subject,
		Teacher:   teacher,
		ClassType: c.ClassType,
	}
}

type ClassViews []ClassView

func (c ClassViews) ToDayGroupedClassesView() DayGroupedClassesView {
	dayGroupedClassesView := make(DayGroupedClassesView, 0)
	dayGroupedClassesView[Monday] = make(NumberGroupedClassesView, 0)
	dayGroupedClassesView[Tuesday] = make(NumberGroupedClassesView, 0)
	dayGroupedClassesView[Wednesday] = make(NumberGroupedClassesView, 0)
	dayGroupedClassesView[Thursday] = make(NumberGroupedClassesView, 0)
	dayGroupedClassesView[Friday] = make(NumberGroupedClassesView, 0)
	dayGroupedClassesView[Saturday] = make(NumberGroupedClassesView, 0)

	for _, class := range c {
		numberGroupedClasses, ok := dayGroupedClassesView[class.DayOfWeek][class.Number]
		if !ok {
			fmt.Printf("\n\ndayGroupedClassesView[%s]:%+v\n\n", class.DayOfWeek, dayGroupedClassesView[class.DayOfWeek])
			dayGroupedClassesView[class.DayOfWeek][class.Number] = []ClassView{class}
			continue
		}

		numberGroupedClasses = append(numberGroupedClasses, class)
		dayGroupedClassesView[class.DayOfWeek][class.Number] = numberGroupedClasses
	}
	return dayGroupedClassesView
}
