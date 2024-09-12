package view

import (
	"time"
	"uni-schedule-backend/internal/domain"
	"uni-schedule-backend/internal/lesson/view"
)

type ScheduleView struct {
	ID       domain.ID                 `json:"id"`
	Name     string                    `json:"name"`
	Slug     string                    `json:"slug"`
	Weekdays []ScheduleGroupedSlotView `json:"weekdays"`
}

type ScheduleGroupedSlotView struct {
	Day   time.Weekday       `json:"day"`
	Slots []ScheduleSlotView `json:"slots"`
}

type ScheduleSlotView struct {
	ID             domain.ID        `json:"id"`
	Number         uint             `json:"number"`
	IsAlternating  bool             `json:"is_alternating"`
	EvenWeekLesson *view.LessonView `json:"even_week_lesson"`
	OddWeekLesson  *view.LessonView `json:"odd_week_lesson"`
}
