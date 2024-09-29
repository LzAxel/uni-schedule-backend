package domain

import (
	"time"
)

type Schedule struct {
	ID        ID
	CreatorID ID
	Name      string
	Slug      string
}

type ScheduleUpdate struct {
	Name      *string
	Slug      *string
	CreatorID *ID
}

type ScheduleSlot struct {
	ID               ID
	ScheduleID       ID
	Weekday          time.Weekday
	Number           uint
	IsAlternating    bool
	EvenWeekLessonID *ID
	OddWeekLessonID  *ID
}

type ScheduleSlotUpdate struct {
	ScheduleID       *ID
	Weekday          *time.Weekday
	Number           *uint
	IsAlternating    *bool
	EvenWeekLessonID *ID
	OddWeekLessonID  *ID
}

type ScheduleView struct {
	ID       ID                        `json:"id"`
	Name     string                    `json:"name"`
	Slug     string                    `json:"slug"`
	Weekdays []ScheduleGroupedSlotView `json:"weekdays"`
}

type ScheduleGroupedSlotView struct {
	Day   time.Weekday       `json:"day"`
	Slots []ScheduleSlotView `json:"slots"`
}

type ScheduleSlotView struct {
	ID             ID          `json:"id"`
	Number         uint        `json:"number"`
	IsAlternating  bool        `json:"is_alternating"`
	EvenWeekLesson *LessonView `json:"even_week_lesson"`
	OddWeekLesson  *LessonView `json:"odd_week_lesson"`
}
