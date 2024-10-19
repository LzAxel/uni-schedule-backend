package domain

import "time"

type ScheduleSlot struct {
	ID               uint64       `db:"id"`
	ScheduleID       uint64       `db:"schedule_id"`
	Weekday          time.Weekday `db:"weekday"`
	Number           uint         `db:"number"`
	IsAlternating    bool         `db:"is_alternating"`
	EvenWeekLessonID *uint64      `db:"even_week_lesson_id"`
	OddWeekLessonID  *uint64      `db:"odd_week_lesson_id"`
}

type ScheduleSlotUpdate struct {
	Weekday          *time.Weekday
	Number           *uint
	IsAlternating    *bool
	EvenWeekLessonID *uint64
	OddWeekLessonID  *uint64
}

type ScheduleSlotCreate struct {
	ScheduleID       uint64
	Weekday          time.Weekday
	Number           uint
	IsAlternating    bool
	EvenWeekLessonID uint64
	OddWeekLessonID  uint64
}
