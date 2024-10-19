package domain

import "time"

type Schedule struct {
	ID        uint64 `db:"id"`
	CreatorID uint64 `db:"creator_id"`
	Name      string `db:"name"`
	Slug      string `db:"slug"`
}

type ScheduleCreate struct {
	Name      string
	Slug      string
	CreatorID uint64
}

type ScheduleUpdate struct {
	Name      *string
	Slug      *string
	CreatorID *uint64
}

type ScheduleView struct {
	ID       uint64                    `json:"id"`
	Name     string                    `json:"name"`
	Slug     string                    `json:"slug"`
	Weekdays []ScheduleGroupedSlotView `json:"weekdays"`
}

type ScheduleGroupedSlotView struct {
	Day   time.Weekday       `json:"day"`
	Slots []ScheduleSlotView `json:"slots"`
}

type ScheduleSlotView struct {
	ID             uint64      `json:"id"`
	Number         uint        `json:"number"`
	IsAlternating  bool        `json:"is_alternating"`
	EvenWeekLesson *LessonView `json:"even_week_lesson"`
	OddWeekLesson  *LessonView `json:"odd_week_lesson"`
}
